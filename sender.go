package main

import (
	"log"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sender() {
	conn, err := amqp.Dial("amqp://guest:guest@10.6.40.220:8080/vhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	for{
		var opcion int
		fmt.Println("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		if opcion == 2 {
			fmt.Println("erai")
			break
		}
		if opcion == 1{
			var mensaje string
			fmt.Println("Escriba el mensaje a enviar: ")
			fmt.Scanln(&mensaje)

			body := mensaje
			err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
			failOnError(err, "Failed to publish a message")
			log.Printf(" [x] Sent %s", body)

		}
	}
}

func main() {
	sender()
}
