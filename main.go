package main

import (
	"log"
	"os"

	"github.com/dwahyudi/golang-smtp/emailing"

	"github.com/dwahyudi/golang-smtp/util"
	"github.com/streadway/amqp"
)

func main() {
	// demo.SimpleMailDemo()

	emailSendWaiter()
}

func emailSendWaiter() {
	channelName := "registration-email-welcome"

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	util.CheckErr(err)
	defer conn.Close()

	ch, err := conn.Channel()
	util.CheckErr(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		channelName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	util.CheckErr(err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.CheckErr(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			emailAddress := string(d.Body)
			emailing.RegistrationWelcomeSend(emailAddress)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
