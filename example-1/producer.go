package example1

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func RunProducer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"gophers",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("queue: ", q)

	for {
		message := "sent at " + time.Now().String()
		err := ch.Publish(
			"",
			"gophers",
			false,
			false,
			amqp.Publishing{
				Headers:     nil,
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		if err != nil {
			break
		}

		time.Sleep(2 * time.Second)
	}
}
