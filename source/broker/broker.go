package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"math/rand"
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
	numero := 1 + rand.Intn(3)
	return &pb.HelloReply{Message: strconv.Itoa(numero)}, nil
}

func (s *server) EnviarComandoLeia(ctx context.Context, in *pb.ComandoSend) (*pb.Rebeldes, error) {
	fmt.Println("Recibiendo respuesta desde Fulcrum hacia Leia")
	//recibir vector y numero de rebeldes desde fulcrum
	n := 1 + rand.Intn(3)
	puerto := ""
	if n == 1{
		// Se eligio el fulcrum 1
		puerto = "10.6.40.218:8080"
	}else if n == 2{
		// Se eligio el fulcrum 2
		puerto = "10.6.40.219:8080"
	}else if n == 3{
		// Se eligio el fulcrum 3
		puerto = "10.6.40.220:8080"
	}
	conn, err := grpc.Dial(puerto, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewFulcrumServiceClient(conn)
	stream, err := serviceClient.EnviarComandoLeia(context.Background(), &pb.ComandoSend{Comando: in.Comando, Vector: in.Vector})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	rebeldes := stream.Numero
	vector := stream.Vector

	return &pb.Rebeldes{Numero: rebeldes, Vector: vector}, nil
}


func ServidorBroker(){
	// Crear el servidor
	fmt.Println("Servido Broker")
	lis, err := net.Listen("tcp", ":8080")
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