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
	fmt.Println("Leia iniciada")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewBrokerServiceClient(conn)
	//Crear un canal para recibir mensajes
	stream, err := serviceClient.SayHello(context.Background(), &pb.HelloRequest{Name: "Leia"})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	//Recibir mensajes
	fmt.Println("Recibiendo mensajes")
	fmt.Println(stream.Message)
}

func enviarComando(mensaje string){
	//Establecer conexion con el servidor broker
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewBrokerServiceClient(conn)
	//Crear un canal para recibir mensajes
	stream, err := serviceClient.EnviarComando(context.Background(), &pb.HelloRequest{Name: mensaje})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	//Recibir mensajes
	fmt.Println("Recibiendo mensajes")
	fmt.Println(stream.Message)
}

func construirComando(){
	fmt.Println("wena llegaste a la boti")
	fmt.Println("sirvase")
	fmt.Println("1. chela")
	fmt.Println("2. chela")
	fmt.Println("3. chela")
	fmt.Println("4. chela")
	fmt.Println("5. palida")
	var opcion int
	fmt.Scan(&opcion)
	switch opcion{
		case 1:
			fmt.Println("elija su marca do chela")
			fmt.Println("1. andes")
			fmt.Println("2. baltica")
			fmt.Println("3. cristal")
			fmt.Println("4. dorada")
			fmt.Println("5. era")
	}
}

func main(){
	mensajeInicial()
	enviarComando("compren chelas")
}