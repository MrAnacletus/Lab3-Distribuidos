package main

import (
	"fmt"
	"log"
	"net"
	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)

func ServidorBroker(){
	// Crear el servidor
	fmt.Println("Servido Broker")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBrokerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}

func main(){
	// Servidor broker
	ServidorBroker()
	
}