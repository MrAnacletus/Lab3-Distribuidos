package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"bufio"
	"regexp"
	"google.golang.org/grpc"
	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
)

//Fulcrumvectores
type Vector struct {
	servidor1 int
	servidor2 int
	servidor3 int
}

type weas struct {
	planeta string
	ciudad string
	rebeldes int
	vector Vector
	servidor string
}

// Variable global de arreglo de vectores
var listaVector []Vector
var listaweas []weas
// variable global de arreglo de nombres de planetas
var nombres []string

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

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

func enviarComando(mensaje string) string{
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
	return stream.Message
}

func ConstruirMensaje(){
	//Pregunta cual de los cuatro comandos utilizar
	fmt.Println("Ingrese el comando que desea utilizar")
	fmt.Println("El comando disponible es:")
	fmt.Println("GetNumberRebelds: GetNumberRebelds <nombre_planeta> <nombre_ciudad>")
	var comando string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	comando = scanner.Text()
	respuesta := enviarComando(comando)
	if respuesta == "1"{
		// Se eligio el fulcrum 1
		fmt.Println("Se eligio el fulcrum 1")
		fmt.Println("reenviando mensaje")
		// Enviar Comando a fulcrum 1
		enviarABroker(comando)
	}
	if respuesta == "2"{
		// Se eligio el fulcrum 2
		fmt.Println("Se eligio el fulcrum 2")
		fmt.Println("reenviando mensaje")
		// Enviar Comando a fulcrum 2
		enviarABroker(comando)
	}
	if respuesta == "3"{
		// Se eligio el fulcrum 3
		fmt.Println("Se eligio el fulcrum 3")
		fmt.Println("reenviando mensaje")
		// Enviar Comando a fulcrum 3
		enviarABroker(comando)
	}
}

func enviarABroker(S string) string{
	// Establecer conexion con el broker
	fmt.Println("Leia iniciada")
	puerto := "localhost:" + fmt.Sprintf("%d", 50051)
	conn, err := grpc.Dial(puerto, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewBrokerServiceClient(conn)
	stream, err := serviceClient.EnviarComandoLeia(context.Background(), &pb.ComandoSend{Comando: S})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	return stream.Numero, stream.Vector
}

func main(){
	mensajeInicial()
	for{
		ConstruirMensaje()
	}
}