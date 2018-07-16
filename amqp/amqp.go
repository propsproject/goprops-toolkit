package amqp

import "context"

// Start start all consumers and publisher
func Start(consumers []ConsumerProducer, ctx context.Context) {
	for _, consumer := range consumers {
		go consumer.Run(ctx)
	}
}

// StartProducers start all consumer producers only as producers, don't declare or announce a queue
func StartProducers(consumers []ConsumerProducer, ctx context.Context) {
	for _, consumer := range consumers {
		go consumer.RunProducer(ctx)
	}
}
