package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedBrokerServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Peticion recibida, aceptando juego")
	fmt.Println("Mensaje: ", in.Name)
	return &pb.HelloReply{Message: "Juego aceptado"}, nil
}

func (s *server) EnviarComando(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Peticion recibida, aceptando juego")
	fmt.Println("Mensaje: ", in.Name)
	// Generar un numero entre 1 y 3
	numero := 1
	// numero := rand.Intn(3)
	return &pb.HelloReply{Message: strconv.Itoa(numero)}, nil
}


func ServidorBroker(){
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

func main(){
	// Servidor broker
	ServidorBroker()
}