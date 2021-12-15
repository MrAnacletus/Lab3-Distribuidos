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
	"time"

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
// Lista de comandos del servidor siguiente a este
var listaComandosServidor1 []string
// Lista de comandos del servidor subsiguiente a este
var listaComandosServidor2 []string
// Indice de que servidor debemos recibir comandos
var indiceComandos int = 1
// Planeta al cual se le está haciendo el merge
var planetaMerge string = ""
// Lista de comandos y vector final
var listaComandosFinal []string


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
	vectorNuevoMerge := ""
	flag := false
	// Separar el mensaje
	// Separar el vector
	// Obtener el numero de rebeldes de la ciudad
	// Enviar el numero de rebeldes al cliente
	mensajeSplit := strings.Split(mensaje, " ")
	vectorRecibido := StringAVector(vector)
	// Buscar el planeta en la lista de nombres
	for idx, planeta := range nombres {
		if planeta == mensajeSplit[1] {
			// comprar el vector con el vector del planeta
			if vectorRecibido.servidor1 > listaVector[idx].servidor1 || vectorRecibido.servidor2 > listaVector[idx].servidor2 || vectorRecibido.servidor3 > listaVector[idx].servidor3 {
				// Inconsistencia encontrada
				fmt.Println("Inconsistencia encontrada")
				// Agregar comando a log{planeta}.txt
				// Checar que log{planeta}.txt exista
				// Si no existe, crearlo
				// Escribir el comando en el archivo
				// Verificar si existe el archivo log
				_, errLog := os.Stat("log" + planeta + fmt.Sprintf("%d",servidor) + ".txt")
				if errLog != nil {
					if os.IsNotExist(errLog) {
						fmt.Println("El archivo no existe")
						file, err := os.Create("log" + planeta + fmt.Sprintf("%d",servidor) + ".txt")
						if err != nil {
							fmt.Println("Error al crear el archivo")
							return 0,"0,0,0"
						}
						file.Close()
						fmt.Println("Archivo creado")
					}
				}
				// Leer archivo Log{planeta}
				lineasLog := leerArchivo("log"+planeta)
				// Agregar comando a la lista de lineas
				lineasLog = append(lineasLog, mensaje)
				// Escribir archivo Log{planeta}
				escribirArchivo("log"+planeta, lineasLog)
				// Ejecutar comando de merge
				vectorNuevoMerge = EJECUTARMERGE(planeta)
				flag = true
			}
		}
	}
	numeroRebeldes, vectorNuevo := GetNumberRebelds(mensajeSplit[1], mensajeSplit[2])
	if !flag {
		fmt.Println("Numero de rebeldes: ", numeroRebeldes)
		fmt.Println("Vector nuevo: ", vectorNuevo)
		// Enviar el numero de rebeldes al cliente
		return numeroRebeldes, vectorNuevo
	} else {
		fmt.Println("Numero de rebeldes: ", numeroRebeldes)
		fmt.Println("Vector nuevo: ", vectorNuevoMerge)
		// Enviar el numero de rebeldes al cliente
		return numeroRebeldes, vectorNuevoMerge
	}
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
			fmt.Println("El vector recibido es ", vectorRecibido)
			fmt.Println("El vector de la ciudad es ", listaVector[idx])
			if vectorRecibido.servidor1 > listaVector[idx].servidor1 || vectorRecibido.servidor2 > listaVector[idx].servidor2 || vectorRecibido.servidor3 > listaVector[idx].servidor3 {
				// Inconsistencia encontrada
				fmt.Println("Inconsistencia encontrada")
				// Agregar comando a log{planeta}.txt
				// Checar que log{planeta}.txt exista
				// Si no existe, crearlo
				// Escribir el comando en el archivo
				// Verificar si existe el archivo log
				_, errLog := os.Stat("log" + palabras[1] + fmt.Sprintf("%d",servidor) + ".txt")
				if errLog != nil {
					if os.IsNotExist(errLog) {
						fmt.Println("El archivo no existe")
						file, err := os.Create("log" + palabras[1] + fmt.Sprintf("%d",servidor) + ".txt")
						if err != nil {
							fmt.Println("Error al crear el archivo")
							return "0,0,0"
						}
						file.Close()
						fmt.Println("Archivo creado")
					}
				}
				// Leer archivo Log{planeta}
				lineasLog := leerArchivo("log"+palabras[1])
				// Agregar comando a la lista de lineas
				lineasLog = append(lineasLog, mensaje)
				// Escribir archivo Log{planeta}
				escribirArchivo("log"+palabras[1], lineasLog)
				// Ejecutar comando de merge
				return EJECUTARMERGE(palabras[1])
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
	file, err := os.Open(pais+ fmt.Sprintf("%d", servidor) +".txt")
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
	file.Close()
	return lineas
}

func escribirArchivo(pais string, lineas []string) {
	fmt.Println("Escribiendo archivo")
	fmt.Println("Pais: ", pais)
	// Eliminar el archivo
	err := os.Remove(pais + fmt.Sprintf("%d",servidor) + ".txt")
	if err != nil {
		fmt.Println("[191]Error al eliminar el archivo", err)
		return
	}
	// Crear el archivo
	file, err := os.Create(pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, err := os.Stat(pais + fmt.Sprintf("%d",servidor) + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Archivo no existe")
			// Crear el archivo
			file, err := os.Create(pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, errLog := os.Stat("log" + pais + fmt.Sprintf("%d",servidor) + ".txt")
	if errLog != nil {
		if os.IsNotExist(errLog) {
			fmt.Println("El archivo no existe")
			file, err := os.Create("log" + pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, err := os.Stat(pais + fmt.Sprintf("%d",servidor) + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
			file, err := os.Create(pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, errLog := os.Stat("log" + pais + fmt.Sprintf("%d",servidor) + ".txt")
	if errLog != nil {
		if os.IsNotExist(errLog) {
			fmt.Println("El archivo no existe")
			file, err := os.Create("log" + pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, err := os.Stat(pais + fmt.Sprintf("%d",servidor) + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
			file, err := os.Create(pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, errLog := os.Stat("log" + pais + fmt.Sprintf("%d",servidor) + ".txt")
	if errLog != nil {
		if os.IsNotExist(errLog) {
			fmt.Println("El archivo no existe")
			file, err := os.Create("log" + pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, err := os.Stat(pais + fmt.Sprintf("%d",servidor) + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
			file, err := os.Create(pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, errLog := os.Stat("log" + pais + fmt.Sprintf("%d",servidor) + ".txt")
	if errLog != nil {
		if os.IsNotExist(errLog) {
			fmt.Println("El archivo no existe")
			file, err := os.Create("log" + pais + fmt.Sprintf("%d",servidor) + ".txt")
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
	_, err := os.Stat(pais + fmt.Sprintf("%d",servidor) + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
			file, err := os.Create(pais + fmt.Sprintf("%d",servidor) + ".txt")
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
		if n == pais {
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

func ServidorFulcrum() {
	// Crear el servidor
	fmt.Println("Servidor Fulcrum")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFulcrumServiceServer(s, &serverFulcrum{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}

// Funcion driver de todo el merge
func EJECUTARMERGE(planeta string) string{
	// Este servidor sera el que haga el merge primero
	// Informar a los otros servers fulcrum para que envien sus comandos
	// Recibir los comandos de los otros servers
	// Ejecutar el merge
	// Enviar los comandos finales a los otros fulcrum
	// Informar a los otros servers fulcrum que se han actualizado
	// Actualizar los vectores de los otros servers
	// retornar el vector final
	fmt.Println("Ejecutando merge")
	fmt.Println("Planeta: ", planeta)
	// Crear conexion que utilizaremos con los servidores fulcrum
	// Distinguir a quienes debemos hacer conexion
	// Crear conexion con el servidor fulcrum 1
	// Crear conexion con el servidor fulcrum 2
	listaServidores := [3]string{"fulcrum1", "fulcrum2", "fulcrum3"}
	if servidor == 1{
		listaServidores[0] = "10.6.40.219:8080"
		listaServidores[1] = "10.6.40.220:8080"
		listaServidores[2] = "10.6.40.218:8080"
	}else if servidor == 2{
		listaServidores[0] = "10.6.40.220:8080"
		listaServidores[1] = "10.6.40.218:8080"
		listaServidores[2] = "10.6.40.219:8080"
	}else{
		listaServidores[0] = "10.6.40.218:8080"
		listaServidores[1] = "10.6.40.219:8080"
		listaServidores[2] = "10.6.40.220:8080"
	}
	// Crear conexion con el servidor fulcrum 1
	conn, err := grpc.Dial(listaServidores[0], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor: %v", err)
	}
	defer conn.Close()
	client := pb.NewFulcrumServiceClient(conn)
	// Crear conexion con el servidor fulcrum 2
	conn2, err := grpc.Dial(listaServidores[1], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor: %v", err)
	}
	defer conn2.Close()
	client2 := pb.NewFulcrumServiceClient(conn2)
	// Enviar comando a los servidores fulcrum
	ComandoEnviado := &pb.HelloRequest{
		Name: planeta + " " + listaServidores[2],
	}
	respuesta, err := client.InformarMerge(context.Background(), ComandoEnviado)
	if err != nil {
		log.Fatalf("[581]Error al enviar el comando: %v", err)
	}
	fmt.Println("Respuesta del servidor 1: ", respuesta.Message)
	// Esperar a que acabe de enviar comandos
	time.Sleep(time.Second * 2)
	respuesta2, err2 := client2.InformarMerge(context.Background(), ComandoEnviado)
	if err2 != nil {
		log.Fatalf("Error al enviar el comando: %v", err2)
	}
	fmt.Println("Respuesta del servidor 2: ", respuesta2.Message)
	// Esperar a que acabe de enviar comandos
	time.Sleep(time.Second * 2)
	// Ejecutar el merge
	ListaComandosFinal, vectorFinal := merge(planeta, listaComandosServidor1, listaComandosServidor2)
	// Enviar los comandos finales a los otros servidores
	fmt.Println("Enviando comandos finales a los servidores fulcrum")
	for idx, comando := range ListaComandosFinal {
		// checar si es el ultimo comando
		var ComandoEnviado *pb.ComandoSend
		if idx == len(ListaComandosFinal)-1 {
			ComandoEnviado = &pb.ComandoSend{
				Comando: comando,
				Vector: "fin",
			}
		}else{
			ComandoEnviado = &pb.ComandoSend{
				Comando: comando,
				Vector: "0,0,0",
			}
		}
		fmt.Println("Enviando comandos finalesa los servidores fulcrum")
		respuesta, err := client.EnviarComandoMergeFinal(context.Background(), ComandoEnviado)
		fmt.Println("Enviando comandos finalesa los servidores fulcrum")
		if err != nil {
			log.Fatalf("Error al enviar el comando: %v", err)
		}
		fmt.Println("Respuesta del servidor 1: ", respuesta.Comando)
		respuesta2, err2 := client2.EnviarComandoMergeFinal(context.Background(), ComandoEnviado)
		if err2 != nil {
			log.Fatalf("Error al enviar el comando: %v", err2)
		}
		fmt.Println("Respuesta del servidor 2: ", respuesta2.Comando)
	}
	fmt.Println("Fin del merge")
	// retornar vector final
	return vectorFinal
}


func (s *serverFulcrum)InformarMerge(ctx context.Context, m *pb.HelloRequest) (*pb.HelloReply, error){
	fmt.Println("Informar Merge")
	fmt.Println("Planeta: ", m.Name)
	// Separar planeta de numero de servidor
	mensajeSeparado := strings.Split(m.Name, " ")
	planetaMerge = mensajeSeparado[0]
	ipServidor := mensajeSeparado[1]
	// Iniciar envio de comandos
	var comandosPropios []string = ObtenerComandosPropios(planetaMerge)
	fmt.Println("Empieza envio de comandos")
	for idx, comando := range comandosPropios {
		// Enviar comando a server que corresponda}
		fmt.Println("Envio comandos a servidor fulcrum: ", ipServidor)
		// Revisar si llegamos al ultimo comando
		if idx == len(comandosPropios)-1 {
			// Enviar comando a servidor
			EnviarComandoServidorFulcrum(comando, ipServidor,true)
		} else {
			EnviarComandoServidorFulcrum(comando, ipServidor,false)
		}
	}
	return &pb.HelloReply{Message: "Finalizado"}, nil
}

func EnviarComandoServidorFulcrum(comando string, ipServidor string, flag bool)(bool){
	// Crear el cliente
	conn, err := grpc.Dial(ipServidor, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar: %v", err)
	}
	defer conn.Close()
	c := pb.NewFulcrumServiceClient(conn)
	// Enviar comando
	var comandoCreado *pb.ComandoSend
	if flag {
		comandoCreado = &pb.ComandoSend{Comando: comando,Vector: "fin"}
	} else {
		comandoCreado = &pb.ComandoSend{Comando: comando,Vector: "0,0,0"}
	}
	respuesta, err := c.EnviarComandoMerge(context.Background(), comandoCreado)
	if err != nil {
		log.Fatalf("[667]Error al enviar comando: %v", err)
	}
	if respuesta.Comando == "Finalizado"{
		return true
	}
	return false
}

func (s *serverFulcrum)EnviarComandoMerge(ctx context.Context, m *pb.ComandoSend) (*pb.ComandoReply, error) {
	// Recibe los comandos de otros fulcrum y los añade a una lista global de comandos
	fmt.Println("Recibiendo comando de merge")
	if indiceComandos == 1{
		if m.Vector == "fin"{
			// Final de la recoleccion de comandos, inicio de la segunda recoleccion
			indiceComandos = 2
			listaComandosServidor1 = append(listaComandosServidor1, m.Comando)
			return &pb.ComandoReply{Comando: "Finalizado"}, nil
		}
		listaComandosServidor1 = append(listaComandosServidor1, m.Comando)
		return &pb.ComandoReply{Comando: "Recibido"}, nil
	}else{
		if m.Vector == "fin"{
			// Final de la recoleccion de comandos, inicio de la segunda recoleccion
			listaComandosServidor2 = append(listaComandosServidor2, m.Comando)
			return &pb.ComandoReply{Comando: "Finalizado"}, nil
		}
		listaComandosServidor2 = append(listaComandosServidor2, m.Comando)
		return &pb.ComandoReply{Comando: "Recibido"}, nil
	}
}

func (s *serverFulcrum)EnviarComandoMergeFinal(ctx context.Context, m *pb.ComandoSend) (*pb.ComandoReply, error) {
	// Recibe los comandos de otros fulcrum y los añade a una lista global de comandos
	fmt.Println("Recibiendo comando de merge final")
	if m.Vector == "fin"{
		// Final de la recoleccion de comandos, inicio de la ejecucion de los comandos
		listaComandosFinal = append(listaComandosFinal, m.Comando)
		// Ejecutar los comandos
		EjecutarComandosMerge(listaComandosFinal)
	}
	listaComandosFinal = append(listaComandosFinal, m.Comando)
	return &pb.ComandoReply{Comando: "Finalizado"}, nil
}

func EjecutarComandosMerge(listaComandos []string){
	// Eliminar los archivos log{planeta} y planeta
	os.Remove("log" + planetaMerge +fmt.Sprintf("%d",servidor)+ ".txt")
	os.Remove(planetaMerge+fmt.Sprintf("%d",servidor)+ ".txt")
	// Ejectuar los comandos
	// Borrar duplicados
	listaComandos = eliminarRepetidos(listaComandos)
	var vectorProvisional string = "0,0,0"
	for _, comando := range listaComandos {
		fmt.Println("Ejecutando comando: ", comando)
		// Obtener el vector del planeta
		vectorProvisional = InterpretarComandosMerge(comando, vectorProvisional)
	}
	vectorProvisional = mergerVector(planetaMerge,vectorProvisional)
	// Actualizar el vector del planeta
	for idx, valor := range nombres {
		if valor == planetaMerge {
			// Separar el vector
			vectorSeparado := strings.Split(vectorProvisional, ",")
			// Actualizar el vector del planeta
			listaVector[idx].servidor1,_ = strconv.Atoi(vectorSeparado[0])
			listaVector[idx].servidor2,_ = strconv.Atoi(vectorSeparado[1])
			listaVector[idx].servidor3,_ = strconv.Atoi(vectorSeparado[2])
		}
	}
	os.Remove("log"+planetaMerge+fmt.Sprintf("%d",servidor)+".txt")
}

func merge(planeta string, listaComandos1 []string, listaComandos2 []string) ([]string, string) {
	// Funcion que tomara dos listas de comandos y un planeta y las unira en una sola lista de comandos
	// Retorna la lista de comandos
	// leer lista de comandos  desde el archivo log{planeta}
	fmt.Println("Merge de comandos")
	// Obtener comndos propios
	var comandosPropios []string = ObtenerComandosPropios(planeta)
	// unir las dos listas de comandos
	comandosPropios = append(comandosPropios, listaComandos1...)
	comandosPropios = append(comandosPropios, listaComandos2...)
	// Eliminar los comandos repetidos
	comandosPropios = eliminarRepetidos(comandosPropios)
	// Ejecutar los comandos
	var vectorProvisional string = "0,0,0"
	os.Remove(planeta+fmt.Sprintf("%d",servidor)+".txt")
	for _, comando := range comandosPropios {
		fmt.Println("Ejecutando comando: ", comando)
		// Obtener el vector del planeta
		// Borrar archivo planeta
		vectorProvisional = InterpretarComandosMerge(comando, vectorProvisional)
	}
	// Merger el vector del planeta con el vector provisional
	vectorProvisional = mergerVector(planeta, vectorProvisional)
	// eliminar los archivos log{planeta} y planeta
	os.Remove("log"+planeta+fmt.Sprintf("%d",servidor)+".txt")
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

func ObtenerComandosPropios(planeta string)([]string){
	// Funcion que obtiene los comandos propios del planeta
	// Retorna la lista de comandos propios
	// leer lista de comandos  desde el archivo log{planeta}
	// Revisar si existe el archivo
	if _, err := os.Stat("log"+planeta+fmt.Sprintf("%d",servidor)+".txt"); err == nil {
		// Si existe el archivo, leerlo
		comandosPropios := leerArchivo("log"+planeta)
		return comandosPropios
	}else{
		// Si no existe el archivo, crearlo
		// crear el archivo con os
		file, err := os.Create("log"+planeta+fmt.Sprintf("%d",servidor)+".txt")
		if err != nil {
			fmt.Println("Error al crear archivo")
		}
		file.Close()
		return []string{}
	}
}

func InterpretarComandosMerge(mensaje string, vector string) string{
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

func main() {
	fmt.Println("Servidor fulcrum")
	//pedir cual servidor fulcrum sera este
	numeroserver := 0
	fmt.Scanln(&numeroserver)
	servidor = numeroserver
	// Iniciar la escucha del servidor
	ServidorFulcrum()
	if servidor == 1{
		// esperar 2 minutos y hacer merge
		fmt.Println("Esperando 2 minutos para hacer merge")
		time.Sleep(time.Minute)
		// Ejecutar el merge
		for _, planeta := range nombres{
			EJECUTARMERGE(planeta)
		}
	}
}