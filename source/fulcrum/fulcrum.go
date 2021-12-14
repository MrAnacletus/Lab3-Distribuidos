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

// declarar tipo de dato vector
type Vector struct {
	servidor1 int
	servidor2 int
	servidor3 int
}
// Variable global de arreglo de vectores
var listaVector []Vector
// variable global de arreglo de nombres de planetas
var nombres []string
// Variable que guarda que servidor fulcrum es
var servidor int


func (s *serverFulcrum) EnviarComando(ctx context.Context, in *pb.ComandoSend) (*pb.ComandoReply, error) {
	fmt.Println("Comando recibido")
	fmt.Println("Comando: " + in.Comando)
	vectorNuevo := interpretarMensaje(in.Comando,in.Vector)
	return &pb.ComandoReply{Comando: in.Comando,Vector: vectorNuevo}, nil
}

func (s *serverFulcrum)EnviarComandoLeia(ctx context.Context, in *pb.ComandoSend) (*pb.Rebeldes, error) {
	fmt.Println("Comando de Leia recibido")
	fmt.Println("Comando: " + in.Comando)
	numeroRebeldes, vectorNuevo := interpretarMensajeLeia(in.Comando,in.Vector)
	return &pb.Rebeldes{Numero: fmt.Sprintf("%d",numeroRebeldes),Vector: vectorNuevo}, nil
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func interpretarMensajeLeia(mensaje string, vector string) (int, string) {
	fmt.Println("Interpretando mensaje")
	fmt.Println("Mensaje: ", mensaje)
	fmt.Println("Vector: ", vector)
	// Separar el mensaje
	// Separar el vector
	// Obtener el numero de rebeldes de la ciudad
	// Enviar el numero de rebeldes al cliente
	mensajeSplit := strings.Split(mensaje, " ")
	// vectorSplit := strings.Split(vector, " ")
	numeroRebeldes, vectorNuevo := GetNumberRebelds(mensajeSplit[1], mensajeSplit[2])
	fmt.Println("Numero de rebeldes: ", numeroRebeldes)
	// Enviar el numero de rebeldes al cliente
	return numeroRebeldes, vectorNuevo
}
func interpretarMensaje(mensaje string, vector string) string {
	//interpretar el mensaje
	fmt.Println("Interpretando mensaje")
	//separar el mensaje en espacios
	fmt.Println("Separando mensaje")
	palabras := regexp.MustCompile(" ").Split(mensaje, -1)
	// Revisar consistencia con el vector <--------------------------------------------------Importante
	vectorRecibido := StringAVector(vector)
	fmt.Println("Vector recibido: ", vectorRecibido)
	// Corroborar que el vector recibido es igual al vector de la ciudad
	// Buscar indice del planeta
	if !stringInSlice(palabras[1], nombres) {
		// Si no esta en la lista de planetas, crear el vector
		NuevoVector(palabras[1])
	}
	for idx, val := range nombres {
		fmt.Println("Nombre planeta: ", val)
		if val == palabras[1] {
			// Si el planeta es el mismo que el que se esta enviando
			if servidor == 1{
				if vectorRecibido.servidor1 > listaVector[idx].servidor1 {
					// Inconsistencia encontrada
					fmt.Println("Inconsistencia encontrada")
				}
			}else if servidor == 2{
				if vectorRecibido.servidor2 > listaVector[idx].servidor2 {
					// Inconsistencia encontrada
					fmt.Println("Inconsistencia encontrada")
				}
			}else if servidor == 3{
				if vectorRecibido.servidor3 > listaVector[idx].servidor3 {
					// Inconsistencia encontrada
					fmt.Println("Inconsistencia encontrada")
				}
			}
		}
	}
	fmt.Println("Interpretando comando")
	var vectorNuevo string
	if palabras[0] == "AddCity" {
		if len(palabras) == 4 {
			vectorNuevo = AddCity(palabras[1], palabras[2], palabras[3])
		} else {
			vectorNuevo = AddCity(palabras[1], palabras[2], "0")
		}
	}else if palabras[0] == "UpdateName" {
		vectorNuevo = UpdateName(palabras[1], palabras[2], palabras[3])
	}else if palabras[0] == "UpdateNumber" {
		vectorNuevo = UpdateValor(palabras[1], palabras[2], palabras[3])
	}else if palabras[0] == "DeleteCity" {
		vectorNuevo = DeleteCity(palabras[1], palabras[2])
	}
	return vectorNuevo
}

func StringAVector(S string) Vector{
	// Separar el vector
	vectorSpliteado := regexp.MustCompile("[,]{1}").Split(S, -1)
	// Convertir a int
	servidor1, _ := strconv.Atoi(vectorSpliteado[0])
	servidor2, _ := strconv.Atoi(vectorSpliteado[1])
	servidor3, _ := strconv.Atoi(vectorSpliteado[2])
	// Crear el vector
	vector := Vector{servidor1, servidor2, servidor3}
	return vector
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

func NuevoVector(pais string) {
	// Crear el vector
	listaVector = append(listaVector, Vector{0, 0, 0})
	// Agregar planeta a lista de nombres
	nombres = append(nombres, pais)
}

func ActualizarVector(pais string, servidor int) string {
	// Buscar indice del planeta
	for idx, val := range nombres {
		if val == pais {
			// Si el planeta es el mismo que el que se esta enviando
			if servidor == 1{
				listaVector[idx].servidor1++
			}else if servidor == 2{
				listaVector[idx].servidor2++
			}else if servidor == 3{
				listaVector[idx].servidor3++
			}
			// Transformar el vector a string
			vectorString := strconv.Itoa(listaVector[idx].servidor1) + "," + strconv.Itoa(listaVector[idx].servidor2) + "," + strconv.Itoa(listaVector[idx].servidor3)
			return vectorString
		}
	}
	return "0,0,0"
}


func AddCity(pais string, ciudad string, valor string) string {
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
				return "0,0,0"
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
				return "0,0,0"
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
	return ActualizarVector(pais, servidor)
}
func UpdateName(pais string, ciudad string, nombre string) string {
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
				return "0,0,0"
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
				return "0"
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
	return ActualizarVector(pais, servidor)
}

func UpdateValor(pais string, ciudad string, valor string) string {
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
				return "0,0,0"
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
				return "0,0,0"
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
	return ActualizarVector(pais, servidor)
}
func DeleteCity(pais string, ciudad string) string {
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
				return "0,0,0"
			}
			file.Close()
			fmt.Println("Archivo creado")
			return "0,0,0"
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
				return "0,0,0"
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
	return ActualizarVector(pais, servidor)
}

func GetNumberRebelds(pais string, ciudad string) (int, string) {
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
				return -1, "0,0,0"
			}
			file.Close()
			fmt.Println("Archivo creado")
			return 0, "0,0,0"
		}
	}
	// Buscar vector de la ciudad consultada
	var vectorCiudad string
	for idx, n := range nombres{
		if n == ciudad {
			vectorEncontrado := listaVector[idx]
			vectorCiudad = strconv.Itoa(vectorEncontrado.servidor1) + "," + strconv.Itoa(vectorEncontrado.servidor2) + "," + strconv.Itoa(vectorEncontrado.servidor3)
		}
	}
	// Leer el archivo linea a linea guardandolas en una variable
	// Buscar la ciudad en el archivo
	// Separar el valor de la ciudad
	// Retornar el valor de la ciudad
	lineas := leerArchivo(pais)
	for _, linea := range lineas {
		if strings.Contains(linea, ciudad) {
			fmt.Println("Encontrada")
			fmt.Println("Linea: ", linea)
			valorCiudad := strings.Split(linea, " ")
			intVar, err := strconv.Atoi(valorCiudad[2])
			if err != nil {
				fmt.Println("Error al convertir el valor a entero")
				return -1, "0,0,0"
			}
			return intVar, vectorCiudad
		}
	}
	return -1, "0,0,0"
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

func (s *serverFulcrum)EnviarComandoMerge(ctx context.Context, m *pb.ComandoMerge) (*pb.Respuesta, error) {
	// Recibe los comandos de otros fulcrum y los añade a una lista global de comandos

}

func merge(planeta string, listaComandos1 []string, listaComandos2 []string) ([]string, string) {
	// Funcion que tomara dos listas de comandos y un planeta y las unira en una sola lista de comandos
	// Retorna la lista de comandos
	// leer lista de comandos  desde el archivo log{planeta}
	comandosPropios := leerArchivo("log"+planeta)
	// unir las dos listas de comandos
	comandosPropios = append(comandosPropios, listaComandos1...)
	comandosPropios = append(comandosPropios, listaComandos2...)
	// Eliminar los comandos repetidos
	comandosPropios = eliminarRepetidos(comandosPropios)
	// eliminar los archivos log{planeta} y planeta
	os.Remove("log"+planeta+".txt")
	os.Remove(planeta+".txt")
	// Ejecutar los comandos
	var vectorProvisional string
	for _, comando := range comandosPropios {
		fmt.Println("Ejecutando comando: ", comando)
		// Obtener el vector del planeta
		vectorProvisional = interpretarMensaje(comando, "0,0,0")
	}
	// Merger el vector del planeta con el vector provisional
	vectorProvisional = mergerVector(planeta, vectorProvisional)

	return comandosPropios, vectorProvisional
}

func mergerVector(planeta string, vector1 string) string {
	// Toma el indice correspondiente al servidor y lo copia a los otros dos indices
	// Retorna el vector actualizado
	// aplicar split al vector
	vectorSeparado := strings.Split(vector1, ",")
	if servidor == 1 {
		vectorSeparado[1] = vectorSeparado[0]
		vectorSeparado[2] = vectorSeparado[0]
	}else if servidor == 2 {
		vectorSeparado[0] = vectorSeparado[1]
		vectorSeparado[2] = vectorSeparado[1]
	}else if servidor == 3 {
		vectorSeparado[0] = vectorSeparado[2]
		vectorSeparado[1] = vectorSeparado[2]
	}
	// Convertir el vector a string
	vectorString := strings.Join(vectorSeparado, ",")
	return vectorString
}

func eliminarRepetidos(lista []string) []string {
	// Funcion que elimina los elementos repetidos de una lista
	// Retorna la lista sin los elementos repetidos
	var listaSinRepetidos []string
	for _, elemento := range lista {
		if !stringInSlice(elemento, listaSinRepetidos) {
			listaSinRepetidos = append(listaSinRepetidos, elemento)
		}
	}
	return listaSinRepetidos
}

func main() {
	fmt.Println("Servidor fulcrum")
	//pedir cual servidor fulcrum sera este
	numeroserver := 0
	fmt.Scanln(&numeroserver)
	servidor = numeroserver
	puertoserver := 50051 + numeroserver
	fmt.Println("Puerto: ", puertoserver)
	// Iniciar la escucha del servidor
	ServidorFulcrum(puertoserver)
}