package amqp

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/propsproject/go-utils/logger"
	lgr "github.com/propsproject/go-utils/logger/v2"
	"github.com/streadway/amqp"
)

// RecoverIntervalTime ...
const RecoverIntervalTime = 6 * 60

// ConsumerProducer ...
type ConsumerProducer interface {
	Run()
	RunProducer()
	String() string
	Connect() error
	ReConnect(int) error
	AnnounceQueue() error
	Close()
	Consume()
	Publish([]byte)
	NewWorker()
}

// RabbitConsumerProducer ...
type RabbitConsumerProducer struct {
	ExchangeName    string
	RoutingKey      string
	Conn            *amqp.Connection
	Channel         *amqp.Channel
	Done            chan error
	ConsumerTag     string // Name that consumer identifies itself to the server with
	URI             string // uri of the rabbitmq server
	ExchangeType    string // topic, direct, etc...
	LastRecoverTime int64
	CurrentStatus   atomic.Value
	ChannelReady    atomic.Value
	Handle          func(amqp.Delivery) bool
	Messages        <-chan amqp.Delivery
	PublishBuffer
	Workers uint64
	Logger  *lgr.LoggerWrapper
}

// PublishBuffer ...
type PublishBuffer struct {
	BufferChan chan []byte
	Stop       chan bool
	Buffer     []([]byte)
}

// Identity ...
func Identity() string {
	hostname, err := os.Hostname()
	h := sha1.New()
	fmt.Fprint(h, hostname)
	fmt.Fprint(h, err)
	fmt.Fprint(h, os.Getpid())
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Run ...
func (rc *RabbitConsumerProducer) Run() {
	rc.ChannelReady.Store(false)

	if err := rc.Connect(); err != nil {
		e := fmt.Errorf("RabbitMQ consumer connect error")
		rc.Logger.Error(e,
			lgr.Field{"error", err.Error()},
			lgr.Field{"consumer-tag", rc.ConsumerTag},
		)
	}
	if err := rc.AnnounceQueue(); err != nil {
		e := fmt.Errorf("RabbitMQ consumer announce error")
		rc.Logger.Error(e,
			lgr.Field{"error", err.Error()},
			lgr.Field{"consumer-tag", rc.ConsumerTag},
		)
	}
	rc.Consume()
	rc.RunPublishBuffer()
	for {
		rc.Logger.Info("Waiting for channel")
		if rc.Channel != nil {
			rc.ChannelReady.Store(true)
			rc.PublishBuffer.Stop <- true
			break
		}
	}
}

// RunProducer ...
func (rc *RabbitConsumerProducer) RunProducer() {
	if err := rc.Connect(); err != nil {
		e := fmt.Errorf("RabbitMQ Consumer run producer error")
		rc.Logger.Error(e,
			lgr.Field{"error", err.Error()},
			lgr.Field{"consumer-tag", rc.ConsumerTag},
		)
	}
}

// ReConnect ...
func (rc *RabbitConsumerProducer) ReConnect(retryTime int) error {
	rc.Close()
	time.Sleep(time.Duration(15+rand.Intn(60)+2*retryTime) * time.Second)
	if err := rc.Connect(); err != nil {
		return err
	}

	err := rc.AnnounceQueue()
	if err != nil {
		return errors.New("Couldn't connect")
	}
	return nil
}

// Connect ...
func (rc *RabbitConsumerProducer) Connect() error {

	var err error
	rc.Conn, err = amqp.Dial(rc.URI)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}

	go func() {
		// Waits here for the channel to be closed
		for {
			select {
			case err := <-rc.Conn.NotifyClose(make(chan *amqp.Error)):
				rc.Logger.Warn("RabbitMQ connection closed",
					lgr.Field{"error", err.Error()},
					lgr.Field{"reason", err.Reason},
					lgr.Field{"code", strconv.Itoa(err.Code)},
					lgr.Field{"consumer-tag", rc.ConsumerTag},
				)
				// Let Handle know it's not time to reconnect
				rc.Done <- errors.New("Channel Closed")
			}
		}
	}()

	rc.Channel, err = rc.Conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}
	if err = rc.Channel.ExchangeDeclare(
		rc.ExchangeName, // name of the exchange
		rc.ExchangeType, // type
		true,            // durable
		false,           // delete when complete
		false,           // internal
		false,           // noWait
		nil,             // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	return nil
}

// AnnounceQueue ...
func (rc *RabbitConsumerProducer) AnnounceQueue() error {
	_, err := rc.Channel.QueueDeclare(
		rc.ConsumerTag, // name of the queue
		true,           // durable
		false,          // delete when usused
		false,          // exclusive
		false,          // noWait
		nil,            // arguments
	)

	if err != nil {
		return fmt.Errorf("Queue Declare: %s", err)
	}

	// Qos determines the amount of messages that the queue will pass to you before
	// it waits for you to ack them. This will slow down queue consumption but
	// give you more certainty that all messages are being processed. As load increases
	// I would reccomend upping the about of Threads and Processors the go process
	// uses before changing this although you will eventually need to reach some
	// balance between threads, procs, and Qos.
	err = rc.Channel.Qos(50, 0, false)
	if err != nil {
		return fmt.Errorf("Error setting qos: %s", err)
	}

	if err = rc.Channel.QueueBind(
		rc.ConsumerTag,  // name of the queue
		rc.RoutingKey,   // routingKey
		rc.ExchangeName, // sourceExchange
		false,           // noWait
		nil,             // arguments
	); err != nil {
		return fmt.Errorf("Queue Bind: %s", err)
	}

	rc.Messages, err = rc.Channel.Consume(
		rc.ConsumerTag, // name
		rc.ConsumerTag, // consumerTag,
		true,           // autoAck
		false,          // exclusive
		false,          // noLocal
		false,          // noWait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("Consume: %s", err)
	}

	return nil
}

// Close ...
func (rc *RabbitConsumerProducer) Close() {
	if rc.Channel != nil {
		rc.Channel.Close()
		rc.Channel = nil
	}
	if rc.Conn != nil {
		rc.Conn.Close()
		rc.Conn = nil
	}
}

// MonitorConn ...
func (rc *RabbitConsumerProducer) MonitorConn() {
	// Go into reconnect loop when
	// c.done is passed non nil values
	for {
		if <-rc.Done != nil {
			rc.CurrentStatus.Store(false)
			retryTime := 1
			for {
				err := rc.ReConnect(retryTime)
				if err != nil {
					e := fmt.Errorf("RabbitMQ Consumer reconnection error")
					rc.Logger.Error(e,
						lgr.Field{"error", err.Error()},
						lgr.Field{"consumer-tag", rc.ConsumerTag},
					)
					retryTime++
				} else {
					break
				}
			}
		}
	}
}

// NewWorker ...
func (rc *RabbitConsumerProducer) NewWorker() {
	atomic.AddUint64(&rc.Workers, 1)
	go func() {
		for msg := range rc.Messages {
			if ok := rc.Handle(msg); ok {
				currentTime := time.Now().Unix()
				if currentTime-rc.LastRecoverTime > RecoverIntervalTime && !rc.CurrentStatus.Load().(bool) {
					rc.CurrentStatus.Store(true)
					rc.LastRecoverTime = currentTime
					rc.Channel.Recover(true)
				} else {
					//msg.Nack(false, true)
					rc.CurrentStatus.Store(false)
				}
			}
		}
	}()
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

// Consume ...
func (rc *RabbitConsumerProducer) Consume() {
	threads := maxParallelism()

	for i := 0; i < threads; i++ {
		rc.NewWorker()
	}
	go rc.MonitorConn()
}

// NewConsumer ...
func NewConsumer(uri, exchangeName, routingKey, exchangeType string, handle func(amqp.Delivery) bool, logger *lgr.LoggerWrapper) *RabbitConsumerProducer {
	consumer := &RabbitConsumerProducer{
		ConsumerTag:  Identity(),
		URI:          uri,
		ExchangeName: exchangeName,
		ExchangeType: exchangeType,
		RoutingKey:   routingKey,
		Done:         make(chan error),
		PublishBuffer: PublishBuffer{
			BufferChan: make(chan []byte),
			Stop:       make(chan bool),
		},
		LastRecoverTime: time.Now().Unix(),
		Handle:          handle,
		Logger:          logger,
	}
	return consumer
}

// Publish ...
func (rc *RabbitConsumerProducer) Publish(payload []byte) {
	if rc.ChannelReady.Load() != nil {
		if rc.ChannelReady.Load().(bool) {
			rc.Channel.Publish(
				rc.ExchangeName, // exchange
				rc.RoutingKey,   // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        payload,
				})

			return
		}
	}

	rc.Logger.Warn("Not Ready Stashing This")
	rc.PublishBuffer.BufferChan <- payload
}

// RunPublishBuffer ...
func (rc *RabbitConsumerProducer) RunPublishBuffer() {
	go func() {
		for {
			select {
			case payload := <-rc.PublishBuffer.BufferChan:
				rc.PublishBuffer.Buffer = append(rc.PublishBuffer.Buffer, payload)
			case <-rc.PublishBuffer.Stop:
				rc.Logger.Info("Stopping")
				return
			}
		}
	}()
}

func (rc *RabbitConsumerProducer) String() string {
	return fmt.Sprintf("Consumer Name: %v, Routing Key: %v, ExchangeName %v ", rc.ConsumerTag, rc.RoutingKey, rc.ExchangeName)
}

func (rc *RabbitConsumerProducer) failOnError(errs ...interface{}) {
	var b bytes.Buffer

	for _, err := range errs {
		b.WriteString(fmt.Sprintf("%v\n", err.(error).Error()))
	}

	rc.Logger.Error(errors.New(b.String()))
}

func (rc *RabbitConsumerProducer) handleSigTerm() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, syscall.SIGTERM)
	for {
		select {
		case <-interrupt:
			logger.Info(fmt.Sprintf("Received interrupt signal, closing. Consumer Name: %v, Routing Key: %v, ExchangeName %v ", rc.ConsumerTag, rc.RoutingKey, rc.ExchangeName))
			rc.Close()
			os.Exit(0)
		}
	}
}
