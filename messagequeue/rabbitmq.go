package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

/*
Common RabbitMQ Exchange Types
Direct Exchange: Same exchange and router key
Topic Exchange: Same exchange and matching wilcard router key
Fanout Exchange: Same exchange => Broadcast to all consumer
Headers Exchange: Combine with header of publishing message
More info:
https://hevodata.com/learn/rabbitmq-exchange-type/,
https://gpcoder.com/6931-su-dung-headers-exchange-trong-rabbitmq/
*/

func TestAMPQ() {
	connectRabbitMQ, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Print(err)
	}
	defer connectRabbitMQ.Close()

	ch, err := connectRabbitMQ.Channel()

	// With the instance and declare Queues that we can
	err = ch.ExchangeDeclare(
		"service.fanout", // name
		"fanout",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	FailOnError(err, "Failed to declare an exchange")

	go InitAMQPConsumer(ch, 1)
	go InitAMQPConsumer(ch, 2)

	time.Sleep(time.Second)

	// Attempt to publish a message to the queue.
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("test message"),
	}
	if err = ch.Publish(
		"service.fanout", // exchange
		"1",              // router key name
		false,            // mandatory
		false,            // immediate
		message,          // message to publish
	); err != nil {
		fmt.Println(err)
	}

}

func InitAMQPConsumer(ch *amqp.Channel, num int) {
	q, err := ch.QueueDeclare(
		strconv.Itoa(num), //name
		false,             //durable
		false,             //delete when usused
		true,              //exclusive
		false,             //no-wait
		nil,               //arguments
	)

	err = ch.QueueBind(
		q.Name,           //queue name
		q.Name,           //routing key
		"service.fanout", //exchange
		false,
		nil,
	)

	deliveries, err := ch.Consume(
		q.Name, //queue
		strconv.Itoa(num),
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		fmt.Println(err)
	}
	for d := range deliveries {
		log.Printf("%d Received  a message: %s \n", num, d.Body)
	}

}

func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}
