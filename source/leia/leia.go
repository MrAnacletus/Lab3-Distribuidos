package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
	"google.golang.org/grpc"
)

//Fulcrumvectores
type Vector struct {
	servidor1 int
	servidor2 int
	servidor3 int
}

// Variable global de arreglo de vectores
var listaVector []Vector
// variable global de arreglo de nombres de planetas
var nombres []string

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

func enviarComando(mensaje string) (string, string){
	//Establecer conexion con el servidor broker
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	// Crear un canal para recibir mensajes
	serviceClient := pb.NewBrokerServiceClient(conn)
	// Separar el mensaje en comando y parametros
	re := regexp.MustCompile(" ")
	parametros := re.Split(mensaje, -1)
	//Otener el planeta
	planeta := parametros[1]
	// Obtener el vector
	// verificr que el vector exista
	var vector Vector
	if len(listaVector) > 0 {
		for idx, val := range nombres {
			if val == planeta {
				vector = listaVector[idx]
				break
			}
			if idx == len(nombres) - 1 {
				fmt.Println("El planeta no existe")
				//Agregar el vector
				vector.servidor1 = 5
				vector.servidor2 = 0
				vector.servidor3 = 0
				listaVector = append(listaVector, vector)
				nombres = append(nombres, planeta)
			}
		}
	} else {
		//Agregar el vector
		vector.servidor1 = 5
		vector.servidor2 = 0
		vector.servidor3 = 0
		listaVector = append(listaVector, vector)
		nombres = append(nombres, planeta)
	}
	// Transformar el vector a string
	vectorString := fmt.Sprintf("%d,%d,%d", vector.servidor1, vector.servidor2, vector.servidor3)
	stream, err := serviceClient.EnviarComandoLeia(context.Background(), &pb.ComandoSend{Comando: mensaje,Vector: vectorString})
	if err != nil {
		log.Fatalf("Error al crear el canal: %v", err)
	}
	//Recibir mensajes
	fmt.Println("Recibiendo mensajes")
	return stream.Numero, stream.Vector
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
	respuesta, vector := enviarComando(comando)
	fmt.Println("Respuesta: ", respuesta)
	fmt.Println("Vector: ", vector)
	// Guardar el vector correspondiente al planeta
	re := regexp.MustCompile(" ")
	parametros := re.Split(comando, -1)
	planeta := parametros[1]
	for idx, val := range nombres {
		if val == planeta {
			// Transformar el string a vector
			re = regexp.MustCompile(",")
			vectorSpliteado := re.Split(vector, -1)
			listaVector[idx].servidor1, _ = strconv.Atoi(vectorSpliteado[0])
			listaVector[idx].servidor2, _ = strconv.Atoi(vectorSpliteado[1])
			listaVector[idx].servidor3, _ = strconv.Atoi(vectorSpliteado[2])
		}
	}
}

func main(){
	mensajeInicial()
	for{
		ConstruirMensaje()
	}
}