package amqp

// Start start all consumers and publisher
func Start(consumers []ConsumerProducer) {
	for _, consumer := range consumers {
		go consumer.Run()
	}
}

// StartProducers start all consumer producers only as producers, don't declare or announce a queue
func StartProducers(consumers []ConsumerProducer) {
	for _, consumer := range consumers {
		go consumer.RunProducer()
	}
}
