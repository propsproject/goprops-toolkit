// +build integration

package amqp

import (
	"testing"

	lgr "github.com/propsproject/go-utils/logger/v2"
	"github.com/streadway/amqp"
)

var testConsumer *RabbitConsumerProducer
var testProducer *RabbitConsumerProducer
var log = lgr.Logger
var testCases []func(*testing.T)

const (
	url           = "amqp://dthifnky:aQCDD5Cim6ddAJ5CZr8u-iQKqpUDTlCY@porpoise.rmq.cloudamqp.com/dthifnky"
	exchange_name = "props"
	routing_key   = "props.transfers"
	exchange_type = "topic"
)

var handle = func(message amqp.Delivery) bool {
	log.Info(string(message.Body))
	return true
}

func setup(t *testing.T) func(t *testing.T) {
	log.Info("setup")
	// t.Log("Setting up test")
	testConsumer = NewConsumer(url, exchange_name, routing_key, exchange_type, handle, log)
	testProducer = NewConsumer(url, exchange_name, routing_key, exchange_type, handle, log)
	testConsumer.Run()
	testConsumer.RunProducer()
	// for {
	// 	consumerReady := testConsumer.CurrentStatus.Load()
	// 	producerReady := testProducer.CurrentStatus.Load()
	// 	if consumerReady != nil && producerReady != nil {
	// 		if consumerReady.(bool) == true && producerReady.(bool) == true {
	// 			testConsumer.NewWorker()
	// 			break
	// 		}
	// 	}
	// 	log.Info("NAH BRUH")
	// }
	testCases = append(testCases, ProduceAndConsume)
	return teardown
}

func teardown(t *testing.T) {
	log.Info("teardown")
	// t.Log("Tearing down test")
	testConsumer.Close()
}

func TestConsumerProducer(t *testing.T) {
	setup(t)
	//defer tearItDown(t)
	for _, testcase := range testCases {
		t.Run("First Case", testcase)
	}
}

func ProduceAndConsume(t *testing.T) {
	// testConsumer.Handle = func(message amqp.Delivery) bool {
	// 	log.Info("YO WE MADE IT", lgr.Field{"payload", string(message.Body)})
	// 	return true
	// }
	testProducer.Publish([]byte("test"))
}
