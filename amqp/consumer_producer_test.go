// +build integration

package amqp

import (
	"sync"
	"testing"

	"github.com/cenkalti/backoff"
	"github.com/propsproject/goprops-toolkit/testutils"
	"github.com/streadway/amqp"
	"github.com/propsproject/goprops-toolkit/logger"
)

var testConsumer *RabbitConsumerProducer
var testProducer *RabbitConsumerProducer
var testCases []testutils.TestCase
var log = logger.NewLogger()
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

	backOff := backoff.WithMaxRetries(backoff.NewConstantBackOff(6), 530)
	oProducer := func() error {
		if ok, err := testProducer.RunProducer(); !ok {
			return err
		}

		return nil
	}

	oConsumer := func() error {
		if ok, err := testConsumer.Run(); !ok {
			return err
		}

		return nil
	}
	t.Log("Starting Rabbitmq docker container")
	stopContainer = testutils.StartRabbitmqContainerV1()

	t.Logf("Trying to connect consumer with constant backoff %v seconds", 3)
	err := backoff.Retry(oConsumer, backOff)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Trying to connect producer with constant backoff %v seconds", 3)
	err = backoff.Retry(oProducer, backOff)
	if err != nil {
		t.Error(err)
	}
	if ok, err := testConsumer.Run(); !ok {
		t.Error(err)
	}

	testCases = append(testCases, ProduceAndConsume)
	return teardown
}

//release any resources, close any connections
func teardown(t *testing.T) {
	t.Log("Running test teardown *consumer_producer_test.go*")
	testConsumer.Close()
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
