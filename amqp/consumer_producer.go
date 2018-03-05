package amqp

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/propsproject/go-utils/logger"
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
	Handle          func(amqp.Delivery) bool
	Messages        <-chan amqp.Delivery
	Workers         uint64
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
func (rc *RabbitConsumerProducer) Run(threads int) {
	if err := rc.Connect(); err != nil {
		logger.Error(fmt.Errorf("[%v]Connect error: %v", rc.ConsumerTag, err))
	}
	if err := rc.AnnounceQueue(); err != nil {
		logger.Error(fmt.Errorf("[%v]AnnounceQueue error: %v", rc.ConsumerTag, err))
	}
	rc.Consume(threads)
}

// RunProducer ...
func (rc *RabbitConsumerProducer) RunProducer() {
	if err := rc.Connect(); err != nil {
		logger.Error(fmt.Errorf("[%v]Connect error: %v", rc.ConsumerTag, err))
	}
}

// ReConnect ...
func (rc *RabbitConsumerProducer) ReConnect(retryTime int) error {
	rc.Close()
	time.Sleep(time.Duration(15+rand.Intn(60)+2*retryTime) * time.Second)

	logger.Info(fmt.Sprintf("Attempting Reconnect: %v", rc.ConsumerTag))
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
		logger.Info(fmt.Sprintf("closing: %v", <-rc.Conn.NotifyClose(make(chan *amqp.Error))))
		// Let Handle know it's not time to reconnect
		rc.Done <- errors.New("Channel Closed")
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
		false,          // noAck
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
					log.Println(fmt.Sprintf("Reconnecting Error %v", err))
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
				msg.Ack(false)
				currentTime := time.Now().Unix()
				if currentTime-rc.LastRecoverTime > RecoverIntervalTime && !rc.CurrentStatus.Load().(bool) {
					rc.CurrentStatus.Store(true)
					rc.LastRecoverTime = currentTime
					rc.Channel.Recover(true)
				} else {
					// this really a litter dangerous. if the worker is panic very quickly,
					// it will ddos our sentry server......plz, add [retry-ttl] in header.
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
func (rc *RabbitConsumerProducer) Consume(threads int) {
	if threads < 1 {
		threads = 1
	}
	for {
		for i := 0; i < threads; i++ {
			rc.NewWorker()
		}
		go rc.MonitorConn()
	}
}

// NewConsumer ...
func NewConsumer(uri, exchangeName, routingKey, exchangeType string, handle func(amqp.Delivery) bool) *RabbitConsumerProducer {
	consumer := &RabbitConsumerProducer{
		ConsumerTag:     Identity(),
		URI:             uri,
		ExchangeName:    exchangeName,
		ExchangeType:    exchangeType,
		RoutingKey:      routingKey,
		Done:            make(chan error),
		LastRecoverTime: time.Now().Unix(),
		Handle:          handle,
	}
	consumer.CurrentStatus.Store(true)
	return consumer
}

// Publish ...
func (rc *RabbitConsumerProducer) Publish(payload []byte) {
	rc.Channel.Publish(
		rc.ExchangeName, // exchange
		rc.RoutingKey,   // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
		})
}

func (rc *RabbitConsumerProducer) String() string {
	return fmt.Sprintf("Consumer Name: %v, Routing Key: %v, ExchangeName %v ", rc.ConsumerTag, rc.RoutingKey, rc.ExchangeName)
}

func failOnError(errs ...interface{}) {
	var b []byte
	fmt.Sprintln(errs)
	for _, err := range errs {
		copy(b, []byte(fmt.Sprintf("%v ", err)))
	}
	logger.Info(errors.New(string(b)).Error())
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
