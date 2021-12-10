package main

import (
	"net"
	"context"
	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"log"
	"sync"

)


type server struct{
	pb.UnimplementedPozoServiceServer
}

var POZOACTUAL int32

var mutex = &sync.Mutex{}

func (s *server) RequestPozo(ctx context.Context, req *pb.RequestPozoActual) (*pb.ResponsePozoActual, error) {
	mutex.Lock()
	aux := POZOACTUAL
	mutex.Unlock()
	return &pb.ResponsePozoActual{
		Pozo: aux,
	}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	go ServidorPozo()
	conn, err := amqp.Dial("amqp://test:test@10.6.40.220:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Canal1", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		mutex.Lock()
		POZOACTUAL = POZOACTUAL + 100000000
		mutex.Unlock()
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ServidorPozo(){
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("No se pudo iniciar el servidor %v", err)
	}
	log.Println("Servidor Pozo Iniciado")
	s := grpc.NewServer()
	pb.RegisterPozoServiceServer(s, &server{})
	if err:= s.Serve(listener); err != nil {
		log.Fatalf("No se pudo iniciar el servidor %v", err)
	}
}