package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/MrAnacletus/Lab3-Distribuidos/source/proto"
	"google.golang.org/grpc"
)

type serverFulcrum struct{
	pb.UnimplementedFulcrumServiceServer
}
// Variable global de vector
// var vector []int = []int{0,0,0}


func (s *serverFulcrum) EnviarComando(ctx context.Context, in *pb.ComandoSend) (*pb.ComandoReply, error) {
	fmt.Println("Comando recibido")
	fmt.Println("Comando: " + in.Comando)
	interpretarMensaje(in.Comando,in.Vector)
	return &pb.ComandoReply{Comando: in.Comando}, nil
}

func (s *serverFulcrum)EnviarComandoLeia(ctx context.Context, in *pb.ComandoSend) (*pb.ComandoReply, error) {
	fmt.Println("Comando de Leia recibido")
	fmt.Println("Comando: " + in.Comando)
	numeroRebeldes := interpretarMensajeLeia(in.Comando,in.Vector)
	return &pb.Rebeldes{Numero: fmt.Sprintf("%d",numeroRebeldes)}, nil
}

func interpretarMensajeLeia(mensaje string, vector string) int {
	fmt.Println("Interpretando mensaje")
	fmt.Println("Mensaje: ", mensaje)
	fmt.Println("Vector: ", vector)
	// Separar el mensaje
	// Separar el vector
	// Obtener el numero de rebeldes de la ciudad
	// Enviar el numero de rebeldes al cliente
	mensajeSplit := strings.Split(mensaje, " ")
	// vectorSplit := strings.Split(vector, " ")
	numeroRebeldes := GetNumberRebelds(mensajeSplit[1], mensajeSplit[2])
	fmt.Println("Numero de rebeldes: ", numeroRebeldes)
	// Enviar el numero de rebeldes al cliente
	return numeroRebeldes
}
func interpretarMensaje(mensaje string, vector string) {
	//interpretar el mensaje
	fmt.Println("Interpretando mensaje")
	//separar el mensaje en espacios
	fmt.Println("Separando mensaje")
	palabras := regexp.MustCompile(" ").Split(mensaje, -1)
	// Revisar consistencia con el vector <--------------------------------------------------Importante
	// vectorSpliteado := regexp.MustCompile("[,]{1}").Split(vector, -1)
	// for idx, valor := range vectorSpliteado {
	// 	fmt.Println("Vector: ", idx, valor)
	// }
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
	}else if palabras[0] == "UpdateNumber" {
		UpdateValor(palabras[1], palabras[2], palabras[3])
	}else if palabras[0] == "DeleteCity" {
		DeleteCity(palabras[1], palabras[2])
	}
}

func leerArchivo(pais string) []string {
	fmt.Println("Leyendo archivo")
	fmt.Println("Pais: ", pais)
	// Abrir archivo en modo lectura
	file, err := os.Open(pais+".txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		log.Fatal(err)
		return nil
	}
	// Leer el archivo
	scanner := bufio.NewScanner(file)
	var lineas []string
	// Leer linea por linea
	for scanner.Scan() {
		lineas = append(lineas, scanner.Text())
	}
	// Cerrar el archivo
	defer file.Close()
	return lineas
}

func escribirArchivo(pais string, lineas []string) {
	fmt.Println("Escribiendo archivo")
	fmt.Println("Pais: ", pais)
	// Eliminar el archivo
	err := os.Remove(pais + ".txt")
	if err != nil {
		fmt.Println("Error al eliminar el archivo")
		return
	}
	// Crear el archivo
	file, err := os.Create(pais + ".txt")
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return
	}
	defer file.Close()
	// Escribir las lineas en el archivo usando bufio
	writer := bufio.NewWriter(file)
	for _, linea := range lineas {
		fmt.Fprintln(writer, linea)
	}
	// Flush del buffer
	writer.Flush()
}
func AddCity(pais string, ciudad string, valor string) {
	fmt.Println("Agregando ciudad")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
	fmt.Println("Valor: ", valor)
	// Checar si el archivo existe
	_, err := os.Stat(pais + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Archivo no existe")
			// Crear el archivo
			file, err := os.Create(pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return
			}
			file.Close()
			fmt.Println("Archivo creado")
		}
	}
	// Leer el archivo
	lineas := leerArchivo(pais)
	// Agregar la ciudad al archivo
	lineas = append(lineas, pais + " " + ciudad + " " + valor)
	// Escribir el archivo
	escribirArchivo(pais, lineas)
	// Verificar si existe el archivo log
	_, errLog := os.Stat("log" + pais + ".txt")
	if errLog != nil {
		if os.IsNotExist(errLog) {
			fmt.Println("El archivo no existe")
			file, err := os.Create("log" + pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return
			}
			file.Close()
			fmt.Println("Archivo creado")
		}
	}
	// Leer archivo Log{planeta}
	lineasLog := leerArchivo("log"+pais)
	// Agregar comando a la lista de lineas
	lineasLog = append(lineasLog, "AddCity "+pais+" "+ciudad+" "+valor)
	// Escribir archivo Log{planeta}
	escribirArchivo("log"+pais, lineasLog)
}
func UpdateName(pais string, ciudad string, nombre string) {
	fmt.Println("Actualizando nombre")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
	fmt.Println("Nombre: ", nombre)
	// Checar si el archivo existe
	// Si no existe, crearlo
	// Si existe, cambiarle el nombre a la ciudad
	_, err := os.Stat(pais + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
			file, err := os.Create(pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return
			}
			file.Close()
			fmt.Println("Archivo creado")
		}
	}
	// Leer el archivo linea a linea guardandolas en una variable
	// Buscar la ciudad en el archivo
	// Cambiar el nombre de la ciudad
	lineas := leerArchivo(pais)
	for idx, linea := range lineas {
		fmt.Println("Linea: ", idx, linea)
		if strings.Contains(linea, ciudad) {
			lineas[idx] = strings.Replace(linea, ciudad, nombre, 1)
		}
	}
	// Nuevas lineas
	fmt.Println("Nuevas lineas: ", lineas)
	// Reescribir el archivo
	escribirArchivo(pais, lineas)
	// Verificar si existe el archivo log
	_, errLog := os.Stat("log" + pais + ".txt")
	if errLog != nil {
		if os.IsNotExist(errLog) {
			fmt.Println("El archivo no existe")
			file, err := os.Create("log" + pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return
			}
			file.Close()
			fmt.Println("Archivo creado")
		}
	}
	// Leer archivo Log{planeta}
	lineasLog := leerArchivo("log"+pais)
	// Agregar comando a la lista de lineas
	lineasLog = append(lineasLog, "UpdateName "+pais+" "+ciudad+" "+nombre)
	// Escribir archivo Log{planeta}
	escribirArchivo("log"+pais, lineasLog)
}

func UpdateValor(pais string, ciudad string, valor string) {
	fmt.Println("Actualizando valor")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
	fmt.Println("Valor: ", valor)
	// Checar si el archivo existe
	// Si no existe, crearlo
	// Si existe, cambiarle el valor a la ciudad
	_, err := os.Stat(pais + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
			file, err := os.Create(pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return
			}
			file.Close()
			fmt.Println("Archivo creado")
		}
	}
	// Leer el archivo linea a linea guardandolas en una variable
	// Buscar la ciudad en el archivo
	// Separar el valor de la ciudad
	// Cambiar el valor de la ciudad
	lineas := leerArchivo(pais)
	for idx, linea := range lineas {
		if strings.Contains(linea, ciudad) {
			// Separar el valor de la ciudad
			valorCiudad := strings.Split(linea, " ")
			// Cambiar el valor de la ciudad
			valorCiudad[2] = valor
			// Reemplazar la linea
			lineas[idx] = strings.Join(valorCiudad, " ")

		}
	}
	// Nuevas lineas
	fmt.Println("Nuevas lineas: ", lineas)
	// Reescribir el archivo
	escribirArchivo(pais, lineas)
	// Verificar si existe el archivo log
	_, errLog := os.Stat("log" + pais + ".txt")
	if errLog != nil {
		if os.IsNotExist(errLog) {
			fmt.Println("El archivo no existe")
			file, err := os.Create("log" + pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return
			}
			file.Close()
			fmt.Println("Archivo creado")
		}
	}
	// Leer archivo Log{planeta}
	lineasLog := leerArchivo("log"+pais)
	// Agregar comando a la lista de lineas
	lineasLog = append(lineasLog, "UpdateValor "+pais+" "+ciudad+" "+valor)
	// Escribir archivo Log{planeta}
	escribirArchivo("log"+pais, lineasLog)
}
func DeleteCity(pais string, ciudad string) {
	fmt.Println("Eliminando ciudad")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
	// Checar si el archivo existe
	// Si no existe, crearlo
	// Si existe, cambiarle el valor a la ciudad
	_, err := os.Stat(pais + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
			file, err := os.Create(pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return
			}
			file.Close()
			fmt.Println("Archivo creado")
			return
		}
	}
	// Leer el archivo linea a linea guardandolas en una variable
	// Buscar la ciudad en el archivo
	// Separar el valor de la ciudad
	// Cambiar el valor de la ciudad
	lineas := leerArchivo(pais)
	lineasNuevas := []string{}
	for _, linea := range lineas {
		if strings.Contains(linea, ciudad) {
			continue
		}
		lineasNuevas = append(lineasNuevas, linea)
	}
	// Nuevas lineas
	fmt.Println("Nuevas lineas: ", lineasNuevas)
	// Reescribir el archivo
	escribirArchivo(pais, lineasNuevas)
	// Verificar si existe el archivo log
	_, errLog := os.Stat("log" + pais + ".txt")
	if errLog != nil {
		if os.IsNotExist(errLog) {
			fmt.Println("El archivo no existe")
			file, err := os.Create("log" + pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return
			}
			defer file.Close()
			fmt.Println("Archivo creado")
		}
	}
	// Leer archivo Log{planeta}
	lineasLog := leerArchivo("log"+pais)
	// Agregar comando a la lista de lineas
	lineasLog = append(lineasLog, "DeleteCity "+pais+" "+ciudad)
	// Escribir archivo Log{planeta}
	escribirArchivo("log"+pais, lineasLog)
}

func GetNumberRebelds(pais string, ciudad string) int{
	fmt.Println("Obteniendo numero de rebeldes")
	fmt.Println("Pais: ", pais)
	fmt.Println("Ciudad: ", ciudad)
	// Checar si el archivo existe
	// Si no existe, crearlo
	// Si existe, retornar el valor de la ciudad
	_, err := os.Stat(pais + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
			file, err := os.Create(pais + ".txt")
			if err != nil {
				fmt.Println("Error al crear el archivo")
				return -1
			}
			file.Close()
			fmt.Println("Archivo creado")
			return 0
		}
	}
	// Leer el archivo linea a linea guardandolas en una variable
	// Buscar la ciudad en el archivo
	// Separar el valor de la ciudad
	// Retornar el valor de la ciudad
	lineas := leerArchivo(pais)
	for _, linea := range lineas {
		if strings.Contains(linea, ciudad) {
			valorCiudad := strings.Split(linea, " ")
			intVar, err := strconv.Atoi(valorCiudad[3])
			if err != nil {
				fmt.Println("Error al convertir el valor a entero")
				return -1
			}
			return intVar
		}
	}
	return -1
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