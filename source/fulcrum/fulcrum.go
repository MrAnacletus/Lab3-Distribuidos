package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"regexp"
	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
	"google.golang.org/grpc"
)

type serverFulcrum struct{
	pb.UnimplementedFulcrumServiceServer
}
// Variable global de vector
var vector []int = []int{0,0,0}


func (s *serverFulcrum) EnviarComando(ctx context.Context, in *pb.ComandoSend) (*pb.ComandoReply, error) {
	fmt.Println("Enviando comando")
	fmt.Println("Name: " + in.Comando)
	interpretarMensaje(in.Comando,in.Vector)
	return &pb.ComandoReply{Comando: in.Comando}, nil
}

func interpretarMensaje(mensaje string, vector string) {
	//interpretar el mensaje
	fmt.Println("Interpretando mensaje")
	//separar el mensaje en espacios
	fmt.Println("Separando mensaje")
	palabras := regexp.MustCompile(" ").Split(mensaje, -1)
	// Revisar consistencia con el vector
	vectorSpliteado := regexp.MustCompile("[(),]+").Split(vector, -1)
	fmt.Println("Vector: ", vectorSpliteado)
	// Interpretar el comando
	fmt.Println("Interpretando comando")
	if palabras[0] == "AddCity" {
		if len(palabras) == 4 {
			AddCity(palabras[1], palabras[2], palabras[3])
		} else {
			AddCity(palabras[1], palabras[2], "0")
		}
	}else if palabras[0] == "UpdateName" {
		UpdateName(palabras[1], palabras[2], palabras[3])
	}else if palabras[0] == "UpdateValor" {
		UpdateValor(palabras[1], palabras[2], palabras[3])
	}else if palabras[0] == "DeleteCity" {
		DeleteCity(palabras[1], palabras[2])
	}
}

func AddCity(pais string, ciudad string, valor string) {
	fmt.Println("Agregando ciudad")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
	fmt.Println("Valor: ", valor)

}
func UpdateName(pais string, ciudad string, nombre string) {
	fmt.Println("Actualizando nombre")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
	fmt.Println("Nombre: ", nombre)
}
func UpdateValor(pais string, ciudad string, valor string) {
	fmt.Println("Actualizando valor")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
	fmt.Println("Valor: ", valor)
}
func DeleteCity(pais string, ciudad string) {
	fmt.Println("Eliminando ciudad")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
}

func ServidorFulcrum(puertoserver int) {
	// Crear el servidor
	fmt.Println("Servidor Fulcrum")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", puertoserver))
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFulcrumServiceServer(s, &serverFulcrum{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}

func main() {
	fmt.Println("Servidor fulcrum")
	//pedir cual servidor fulcrum sera este
	numeroserver := 0
	fmt.Scanln(&numeroserver)
	puertoserver := 50051 + numeroserver
	fmt.Println("Puerto: ", puertoserver)
	// Iniciar la escucha del servidor
	ServidorFulcrum(puertoserver)
}