package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
	"google.golang.org/grpc"
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
	stream, err := serviceClient.EnviarComando(context.Background(), &pb.HelloRequest{Name: mensaje})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	//Recibir mensajes
	fmt.Println("Recibiendo mensajes")
	fmt.Println(stream.Message)
}

func ConstruirMensaje(){
	//Pregunta cual de los cuatro comandos utilizar
	fmt.Println("Ingrese el comando que desea utilizar")
	fmt.Println("1. AddCity")
	fmt.Println("2. UpdateName")
	fmt.Println("3. UpdateNumber")
	fmt.Println("4. DeleteCity")
	fmt.Println("5. Salir")
	var comando int
	fmt.Scan(&comando)
	switch comando {
		case 1:
			fmt.Println("Ingrese el nombre del planeta")
			var nombrePlaneta string
			fmt.Scan(&nombrePlaneta)
			fmt.Println("Ingrese el nombre de la ciudad")
			var nombreCiudad string
			fmt.Scan(&nombreCiudad)
			fmt.Println("Ingrese el valor, si no quiere ingresar un valor, ingrese 0")
			var numero int
			fmt.Scan(&numero)
			if numero == 0{
				enviarMensaje("AddCity " + nombrePlaneta + " " + nombreCiudad)
			}else{
				enviarMensaje("AddCity " + nombrePlaneta + " " + nombreCiudad + " " + strconv.Itoa(numero))
			}
			break
		case 2:
			fmt.Println("Ingrese el nombre del planeta")
			var nombrePlaneta string
			fmt.Scan(&nombrePlaneta)
			fmt.Println("Ingrese el nombre de la ciudad")
			var nombreCiudad string
			fmt.Scan(&nombreCiudad)
			fmt.Println("Ingrese el nuevo nombre de la ciudad")
			var nuevoNombre string
			fmt.Scan(&nuevoNombre)
			enviarMensaje("UpdateName " + nombrePlaneta + " " + nombreCiudad + " " + nuevoNombre)
			break
		case 3:
			fmt.Println("Ingrese el nombre del planeta")
			var nombrePlaneta string
			fmt.Scan(&nombrePlaneta)
			fmt.Println("Ingrese el nombre de la ciudad")
			var nombreCiudad string
			fmt.Scan(&nombreCiudad)
			fmt.Println("Ingrese el nuevo numero de habitantes")
			var nuevoNumero int
			fmt.Scan(&nuevoNumero)
			enviarMensaje("UpdateNumber " + nombrePlaneta + " " + nombreCiudad + " " + strconv.Itoa(nuevoNumero))
			break
		case 4:
			fmt.Println("Ingrese el nombre del planeta")
			var nombrePlaneta string
			fmt.Scan(&nombrePlaneta)
			fmt.Println("Ingrese el nombre de la ciudad")
			var nombreCiudad string
			fmt.Scan(&nombreCiudad)
			enviarMensaje("DeleteCity " + nombrePlaneta + " " + nombreCiudad)
			break
		case 5:
			fmt.Println("Saliendo...")
			break
		default:
			fmt.Println("Comando no reconocido")
			break
	}
}

func main(){
	//Establecemos conexion con el servidor broker
	mensajeInicial()
	//Enviar mensaje
	mensaje:= ConstruirMensaje()
	enviarMensaje(mensaje)
}