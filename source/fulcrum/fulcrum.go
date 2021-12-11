package main

import (
	"net"
	"log"
	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
	"google.golang.org/grpc"
	"fmt"
)



func ServidorFulcrum(){
	// Crear el servidor
	fmt.Println("Servido Broker")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBrokerServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}

func main() {
	fmt.Println("Servidor fulcrum")
	//pedir cual servidor fulcrum sera este
	numeroserver := 0
	fmt.Scanln(&numeroserver)
	puertoserver := 50054 + numeroserver
	fmt.Println("Puerto: ", puertoserver)
	// Iniciar la escucha del servidor

	
}