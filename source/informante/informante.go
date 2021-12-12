package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
	"google.golang.org/grpc"
)

// Vectores de Fulcrum
var vector1 = []int{0,0,0}
var vector2 = []int{0,0,0}
var vector3 = []int{0,0,0}


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

func enviarMensaje(mensaje string) string{
	//Establecer conexion con el servidor broker
	fmt.Println("Informante iniciado")
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
	fmt.Println("Respondiendo")
	return stream.Message
}

func ConstruirMensaje(){
	//Pregunta cual de los cuatro comandos utilizar
	fmt.Println("Ingrese el comando que desea utilizar")
	fmt.Println("Los comandos disponibles son:")
	fmt.Println("AddCity: AddCity <planeta> <nombre_ciudad> <poblacion> si no quiere poner poblcion escriba 0")
	fmt.Println("UpdateName: UpdateName <planeta> <nombre_ciudad> <nuevo_nombre>")
	fmt.Println("UpdateNumber: UpdateNumber <planeta> <nombre_ciudad> <nueva_poblacion>")
	fmt.Println("DeleteCity: DeleteCity <planeta> <nombre_ciudad>")
	var comando string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	comando = scanner.Text()
	log.Println("Comando: ", comando)
	respuesta := enviarMensaje(comando)
	if respuesta == "1"{
		// Se eligio el fulcrum 1
		fmt.Println("Se eligio el fulcrum 1")
		fmt.Println("reenviando mensaje")
		// Enviar Comando a fulcrum 1
		enviarAFulcrum(1, comando)
	}
}

func enviarAFulcrum(n int, S string) string{
	// Establecer conexion con el servidor fulcrum
	fmt.Println("Informante iniciado")
	puerto := "localhost:" + fmt.Sprintf("%d", 50051+n)
	conn, err := grpc.Dial(puerto, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewFulcrumServiceClient(conn)
	// Crear un canal para recibir mensajes
	// Escribir el vector de Fulcrum
	vector := ""
	switch n {
	case 1:
		for i := 0; i < 2; i++ {
			vector = vector + fmt.Sprintf("%d,", vector1[i])
		}
		vector = vector + fmt.Sprintf("%d", vector1[2])
	case 2:
		for i := 0; i < 2; i++ {
			vector = vector + fmt.Sprintf("%d,", vector2[i])
		}
		vector = vector + fmt.Sprintf("%d", vector2[2])
	case 3:
		for i := 0; i < 2; i++ {
			vector = vector + fmt.Sprintf("%d,", vector3[i])
		}
		vector = vector + fmt.Sprintf("%d", vector3[2])
	}
	stream, err := serviceClient.EnviarComando(context.Background(), &pb.ComandoSend{Comando: S, Vector: vector})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	//Recibir mensajes
	fmt.Println("Respondiendo")
	return stream.Comando
}

func main(){
	//Establecemos conexion con el servidor broker
	mensajeInicial()
	//Enviar mensaje
	ConstruirMensaje()
}