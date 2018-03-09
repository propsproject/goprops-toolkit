// +build integration

package amqp

import (
	"sync"
	"testing"

	lgr "github.com/propsproject/go-utils/logger/v2"
	"github.com/propsproject/go-utils/testutils"
	"github.com/streadway/amqp"
)

var testConsumer *RabbitConsumerProducer
var testProducer *RabbitConsumerProducer
var testCases []testutils.TestCase
var log = lgr.Logger
var stopContainer *chan bool

const (
	url           = "amqp://user:user@127.0.0.1:5672/"
	exchange_name = "goUtilsTest"
	routing_key   = "goUtilsTest.transfers"
	exchange_type = "topic"
)

var handle = func(message amqp.Delivery) bool {
	return true
}

/*
We setup our integration test by creating 2 RabbitConsumerProducer. One that will serve as a full fledged consumer, and one that will serve as just the producer.
Once we know our RabbitConsumerProducer(s) are ready we append all the test cases we want to run after the setup and return our teardown func. Our teardown func will
be invoked inside the wrapper Test that ranges through our test cases and calls t.Run on all the "Test" funcs
*/
func setup(t *testing.T) func(t *testing.T) {
	t.Log("Running test setup *consumer_producer_test.go*")
	testConsumer = NewConsumer(url, exchange_name, routing_key, exchange_type, handle, log)
	testProducer = NewConsumer(url, exchange_name, routing_key, exchange_type, handle, log)

	t.Log("Starting Rabbitmq docker container")
	stopContainer = testutils.StartRabbitmqContainerV1()

	t.Log("Starting a test consumer")
	if ok, err := testConsumer.Run(); !ok {
		t.Errorf("Unexpected error: %v", err)
	}

	t.Log("Starting a test producer")
	if ok, err := testProducer.RunProducer(); !ok {
		t.Errorf("Unexpected error: %v", err)
	}

	testCases = append(testCases, ProduceAndConsume)
	return teardown
}

//release any resources, close any connections
func teardown(t *testing.T) {
	t.Log("Running test teardown *consumer_producer_test.go*")
	testConsumer.Close()
	*stopContainer <- true
}

//Our test wrapper that gets invoke by the go runtime allowing us to range through our actual test cases and call t.Run for each test func
func TestConsumerProducer(t *testing.T) {
	tearItDown := setup(t)
	defer tearItDown(t)
	for _, testcase := range testCases {
		t.Run(testcase.Name, testcase.Test)
	}
}

var ProduceAndConsume = testutils.TestCase{
	Name: "ProduceAndConsume",
	Test: func(t *testing.T) {
		t.Logf("%v: Producer should produce payload, test consumer should consume that payload with minimal delay and no errors", t.Name())
		var consumed bool

		/*
			here we use self invoking closure so we can pass a reference to t for asserting, consumed which lets our
			waitgroup know its done, and the message that our handler receives for manipulating etc
		*/
		testConsumer.Handle = func(message amqp.Delivery) bool {
			return func(t *testing.T, c *bool, m *amqp.Delivery) bool {
				*c = true
				return true
			}(t, &consumed, &message)
		}

		var wg sync.WaitGroup
		testProducer.Publish([]byte("test"))

		// wait for our previously published message to be consumed
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				if consumed {
					break
				}
			}
		}(&wg)
		wg.Wait()
	},
}
