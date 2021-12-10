package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
)

func mensajeInicial(){
	//Establecer conexion con el servidor broker
	fmt.Println("Informante iniciado")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewBrokerServiceClient(conn)
	//Crear un canal para recibir mensajes
	stream, err := serviceClient.SayHello(context.Background(), &pb.HelloRequest{Name: "Informante"})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	//Recibir mensajes
	fmt.Println("Recibiendo mensajes")
	fmt.Println(stream.Message)
}

func enviarMensaje(mensaje string){
	//Establecer conexion con el servidor broker
	fmt.Println("Informante iniciado")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewBrokerServiceClient(conn)
	//Crear un canal para recibir mensajes
	stream, err := serviceClient.SayHello(context.Background(), &pb.HelloRequest{Name: mensaje})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	//Recibir mensajes
	fmt.Println("Recibiendo mensajes")
	fmt.Println(stream.Message)
}

func main(){
	//Establecemos conexion con el servidor broker
	mensajeInicial()
	enviarMensaje("Hola")
}