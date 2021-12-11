package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"os"
	"bufio"
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
	return stream.Message
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
	fmt.Println(stream.Message)
}

func boti(){
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
			fmt.Println("ingrese su marca do chela")
			var chela string
			fmt.Scan(&chela)
			fmt.Println("ingrese el numero do chelas que se va a llevar")
			var cantidad int
			fmt.Scan(&cantidad)
			if cantidad == 0 {
				enviarComando("era la wea")
			} else {
				enviarComando("se lleva "+strconv.Itoa(cantidad)+"de "+chela)
			}
		case 2:
			fmt.Println("ingrese su marca do chela")
			var chela string
			fmt.Scan(&chela)
			fmt.Println("ingrese el numero do chelas que se va a llevar")
			var cantidad int
			fmt.Scan(&cantidad)
			if cantidad == 0 {
				enviarComando("era la wea")
			} else{
				enviarComando("se lleva "+strconv.Itoa(cantidad)+"de "+chela)
			}
		case 3:
			fmt.Println("ingrese su marca do chela")
			var chela string
			fmt.Scan(&chela)
			fmt.Println("ingrese el numero do chelas que se va a llevar")
			var cantidad int
			fmt.Scan(&cantidad)
			if cantidad == 0{
				enviarComando("era la wea")
			} else{
				enviarComando("se lleva "+strconv.Itoa(cantidad)+"de "+chela)
			}
		case 4:
			fmt.Println("ingrese su marca do chela")
			var chela string
			fmt.Scan(&chela)
			fmt.Println("ingrese el numero do chelas que se va a llevar")
			var cantidad int
			fmt.Scan(&cantidad)
			if cantidad == 0{
				enviarComando("era la wea")
			} else{
				enviarComando("se lleva "+strconv.Itoa(cantidad)+" de "+chela)
			}
		case 5:
			fmt.Println("era")
			fmt.Println("era la wea cabros")
		default:
			fmt.Println("k chucha")
	}
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
	respuesta := enviarComando(comando)
	fmt.Println(respuesta)
}

func main(){
	mensajeInicial()
	boti()
}