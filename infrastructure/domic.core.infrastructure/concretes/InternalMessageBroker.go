package concretes

import (
	"domic.core.usecase/commons/contracts"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type InternalMessageBroker struct {
	serializer contracts.ISerializer
	connection *amqp091.Connection
}

func (broker *InternalMessageBroker) Subscribe(queue string, closure func(body []byte) error) error {

	ch, err := broker.connection.Channel()

	if err != nil {

		err := ch.Close()

		if err != nil {
			return err
		}

		failOnError(err, "Failed to open a channel")

		//long runing goroutin
		go func() {

			messageChannel, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,  // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)

			failOnError(err, "Failed to register a consumer")

			waiter := &sync.WaitGroup{}

			var counter int = 0

			for {

				//throttle policy
				if counter <= 10000 {

					counter++

					waiter.Add(counter)

					newMessage := <-messageChannel

					//concurrent processing current message ( event )
					go func() {

						defer waiter.Done()

						result := closure(newMessage.Body)

						if result == nil {
							newMessage.Acknowledger.Ack(newMessage.DeliveryTag, false)
						} else {
							newMessage.Acknowledger.Reject(newMessage.DeliveryTag, false)
						}

					}()

				} else {

					waiter.Wait()

					counter = 0

				}

			}

		}()
	}

}

func (broker *InternalMessageBroker) Publish(event interface{}, exchange string, routeKey *string) error {
	return nil
}

/*-------------------------------------------------------------------*/

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
