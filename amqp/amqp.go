package amqp

// Start start all consumers and publisher
func Start(consumers []ConsumerProducer) {
	for _, consumer := range consumers {
		go consumer.Run()
	}
}
