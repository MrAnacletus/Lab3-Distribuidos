package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"

	"google.golang.org/grpc"
)

type Jugada struct {
	ID int32
	jugada int32
}
type Jugada2 struct{
	ID1 []int32
	ID2 []int32
	Jugada1 []int32
	Jugada2 []int32
}

type Equipo struct {
	ID1 int32
	ID2 int32
	Jugada1 int32
	Jugada2 int32
}

func EnviarPeticionJugar(){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la petición de jugar al servidor
	res, err := serviceCLient.SayHello(context.Background(), &pb.HelloRequest{Name: "Preparandoce para iniciar el juego!"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}

func EnviarJugada(J Jugada)(bool){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la Jugada al servidor
	res, err := serviceCLient.SendJugada(context.Background(), &pb.Jugada{ID: J.ID, Jugada: J.jugada})
	if err != nil {
		log.Fatalf("Error al enviar la Jugada: %v", err)
	}
	return res.GetEstado() == 1
}

func EnviarJugada2(J Jugada2)(Jugada){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la Jugada al servidor
	res, err := serviceCLient.SendJugada2(context.Background(), &pb.Jugada2{ID1: J.ID1,ID2: J.ID2, Jugada1: J.Jugada1, Jugada2: J.Jugada2})
	if err != nil {
		log.Fatalf("Error al enviar la Jugada: %v", err)
	}

	return Jugada{ID: res.GetID(), jugada: res.GetEstado()}
}

func EnviarJugada3(J Equipo)(Jugada){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la Jugada al servidor
	res, err := serviceCLient.SendJugada3(context.Background(), &pb.Jugada3{ID1: J.ID1,ID2: J.ID2, Jugada1: J.Jugada1, Jugada2: J.Jugada2})
	if err != nil {
		log.Fatalf("Error al enviar la Jugada: %v", err)
	}
	return Jugada{ID: res.GetID(), jugada: res.GetEstado()}

}
func PedirPozo(){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la petición de pozo al servidor
	res, err := serviceCLient.RequestPozo(context.Background(), &pb.RequestPozoActual{Pozo: 0})
	if err != nil {
		panic(err)
	}
	fmt.Println("El pozo actual es: " + fmt.Sprint(res.GetPozo()))

}

type Bots struct{
	ID int32
	Estado int32
	Team int32
}

var ListaJugadores [16]Bots

var Vivos int32



func juego1()(bool){
	fmt.Println("--------------------------------------------------------O--------------------------------------------------------")
	fmt.Println("Primer Juego!!!!!!!!!!")
	var suma int32 = 0
	for i := 0; i < 4; i++ {
		if ListaJugadores[0].Estado == 1 {
			fmt.Println("Es tu turno")
			fmt.Println("Selecciona una jugada")
			fmt.Println("Elija un numero entre 1 y 10, recuerde que sus numeros deben sumar 21. Actualmente tienes" + fmt.Sprint(suma))
			var numero1 int
			fmt.Scanln(&numero1)
			suma += int32(numero1)
			var j = Jugada{ID: 1, jugada: int32(numero1)}
			if EnviarJugada(j){
				fmt.Println("El jugador 1 sigue vivo")
			}else{
				fmt.Println("El jugador 1 ha muerto")
				ListaJugadores[0].Estado = 0
				Vivos -= 1
			}
			if Vivos == 1 {
				fmt.Println("Tenemos un ganador la conchadesumadre")
				return false
			}
		}

		//Juego de los bots
		for i:= 1; i < 16; i++ {
			if ListaJugadores[i].Estado == 1 {
				numero2 := rand.Intn(10) + 1
				var j = Jugada{ID: ListaJugadores[i].ID, jugada: int32(numero2)}
				if EnviarJugada(j){
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " sigue vivo")
				}else{
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " ha muerto")
					ListaJugadores[i].Estado = 0
					Vivos -= 1
				}
				if Vivos == 1 {
					fmt.Println("Tenemos un ganador")
					return false
				}
			}
		}
	}

	return true
}

func juego2()(bool){
	fmt.Println("--------------------------------------------------------O--------------------------------------------------------")
	fmt.Println("Segundo Juego!!!!!!!!!!")
	var listaT1 []int32
	var listaT2 []int32
	var listaJugadas1 []int32
	var listaJugadas2 []int32
	var suma1 int32 = 0
	var suma2 int32 = 0
	//Dividir los equipos
	//Contar los vivos
	if Vivos % 2 != 0 {
		Elegido := rand.Intn(int(Vivos))
		for i := 0; i < 16; i++ {
			if ListaJugadores[i].Estado == 1 {
				if Elegido == 0 {
					ListaJugadores[i].Estado = 0
				}
				Elegido -= 1
			}
		}
		Vivos -= 1
	}
	for i := 0 ;i < 16; i++ {
		if ListaJugadores[i].Estado == 1 {
			team := rand.Intn(1) + 1
			if team == 1 {
				if suma1 == Vivos/2 {
					team = 2
				}else{
					suma1 += 1
				}
			}else{
				if suma2 == Vivos/2 {
					team = 1
				}else{
					suma2 += 1
				}
			}
			ListaJugadores[i].Team = int32(team)
		}
	}
	suma1 = 0
	suma2 = 0
	//Contar las sumas
	for i := 0; i < 16; i++ {
		if ListaJugadores[i].Estado == 1 {
			if ListaJugadores[i].Team == 1 {
				if ListaJugadores[i].ID == 1 {
					fmt.Println("Es tu turno, usted pertenece al equipo 1")
					fmt.Println("Selecciona una jugada")
					fmt.Println("Elija un numero entre 1 y 4, la suma de su equipo debera tener la misma paridad que el numero elegido por el lider")
					var numero1 int
					fmt.Scanln(&numero1)
					listaJugadas1 = append(listaJugadas1, int32(numero1))
					suma1 += int32(numero1)

				}else{
					numero2 := rand.Intn(4) + 1
					suma1 += int32(numero2)
					listaJugadas1 = append(listaJugadas1, int32(numero2))
				}
			}else{
				if ListaJugadores[i].ID == 1 {
					fmt.Println("Es tu turno, usted pertenece al equipo 2")
					fmt.Println("Selecciona una jugada")
					fmt.Println("Elija un numero entre 1 y 4, la suma de su equipo debera tener la misma paridad que el numero elegido por el lider")
					var numero1 int
					fmt.Scanln(&numero1)
					listaJugadas2 = append(listaJugadas2, int32(numero1))
					suma2 += int32(numero1)

				}else{
					numero2 := rand.Intn(4) + 1
					listaJugadas2 = append(listaJugadas2, int32(numero2))
					suma2 += int32(numero2)
				}
			}
		}
	}

	//Enviar sumas al Lider
	for i := 0; i < 16; i++ {
		if ListaJugadores[i].Estado == 1 {
			if ListaJugadores[i].Team == 1 {
				listaT1 = append(listaT1, ListaJugadores[i].ID)
			}else{
				listaT2 = append(listaT2, ListaJugadores[i].ID)
			}
		}
	}
	var resultado = EnviarJugada2(Jugada2{ID1: listaT1,ID2: listaT2, Jugada1: listaJugadas1, Jugada2: listaJugadas2})
	//Recorrer los jugadores
	for i := 0; i < 16; i++{
		if ListaJugadores[i].Estado == 1{
			if ListaJugadores[i].Team == 1 {
				if resultado.ID == 1 {
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " sigue vivo")
				}else{
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " ha muerto")
					ListaJugadores[i].Estado = 0
					Vivos -= 1
				}
			}else{
				if resultado.jugada == 1 {
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " sigue vivo")
				}else{
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " ha muerto")
					ListaJugadores[i].Estado = 0
					Vivos -= 1
				}
			}
		}
	}
	if Vivos == 1 {
		fmt.Println("Tenemos un ganador")
		return false
	}
	return true
}

func juego3()(bool){
	fmt.Println("--------------------------------------------------------O--------------------------------------------------------")
	fmt.Println("Tercer Juego!!!!!!!!!!")
	equipos := Vivos/2
	//Dividir los equipos
	var TIMS []Equipo

	flag := true

	for i:=0; i < 16; i++ {
		if ListaJugadores[i].Estado == 1 {
			if flag{
				TIMS = append(TIMS, Equipo{ID1: ListaJugadores[i].ID})
				equipos -= 1
				flag = false
			}else{
				TIMS[Vivos/2 - 1 - equipos].ID2 = ListaJugadores[i].ID
				flag = true
			}
		}
	}
	for i:=0 ; i<int(Vivos/2); i++ {
		fmt.Println("Equipo " + fmt.Sprint(i+1))
		fmt.Println("Jugador 1: " + fmt.Sprint(TIMS[i].ID1))
		fmt.Println("Jugador 2: " + fmt.Sprint(TIMS[i].ID2))
		if TIMS[i].ID1 == 1 || TIMS[i].ID2 == 1 {
			fmt.Println("Es tu turno, usted pertenece al equipo " + fmt.Sprint(i+1))
			fmt.Println("Selecciona una jugada")
			fmt.Println("Elija un numero entre 1 y 10, el numero que elijas debe estar cercano al numero que eligira el lider")
			var numero1 int
			fmt.Scanln(&numero1)
			if TIMS[i].ID1 == 1{
				TIMS[i].Jugada1 = int32(numero1)
			}else{
				TIMS[i].Jugada2 = int32(numero1)
			}
		}else{
			numero2 := rand.Intn(10) + 1
			TIMS[i].Jugada1 = int32(numero2)
			numero3 := rand.Intn(10) + 1
			TIMS[i].Jugada2 = int32(numero3)
		}
		//Enviar la jugada del juego 3
		J := EnviarJugada3(TIMS[i])

		if J.ID == 1{
			fmt.Println("El jugador " + fmt.Sprint(TIMS[i].ID1) + " sigue vivo")
			fmt.Println("El jugador " + fmt.Sprint(TIMS[i].ID2) + " ha muerto")
			ListaJugadores[TIMS[i].ID2].Estado = 0
		}else{
			fmt.Println("El jugador " + fmt.Sprint(TIMS[i].ID1) + " ha muerto")
			fmt.Println("El jugador " + fmt.Sprint(TIMS[i].ID2) + " sigue vivo")
			ListaJugadores[TIMS[i].ID1].Estado = 0
		}

	}
	if Vivos == 1 {
		fmt.Println("Tenemos un ganador")
		return false
	}
	return true
}


func main(){
	for i := 0; i < 16; i++ {
		ListaJugadores[i].ID = int32(i + 1)
		ListaJugadores[i].Estado = 1
	}

	Vivos = 16

	fmt.Println("Iniciando el cliente")
	fmt.Println("Bienvenido al Juego del Calamar")
	fmt.Println("Para comenzar a jugar, presiona enter")
	fmt.Scanln()
	EnviarPeticionJugar()
	fmt.Println("El juego ha comenzado")
	fmt.Println("Selecciona que vas a hacer:")
	fmt.Println("1. Jugar")
	fmt.Println("2. Ver pozo")
	var opcion int
	fmt.Scanln(&opcion)
	for {
		if opcion == 1 {
			if !juego1(){
				fmt.Println("El juego ha terminado, y tenemos un ganador!!!!!!!!!")
				for i := 0; i < 16; i++ {
					if ListaJugadores[i].Estado == 1 {
						fmt.Println("El ganador es el jugador " + fmt.Sprint(ListaJugadores[i].ID))
					}
				}
				fmt.Println("Selecciona que vas a hacer:")
				fmt.Println("1. Jugar")
				fmt.Println("2. Ver pozo")
				fmt.Scanln(&opcion)
				if opcion == 1 {
					break
				}else{
					PedirPozo()
				}
			}
			fmt.Println("Selecciona que vas a hacer:")
				fmt.Println("1. Jugar")
				fmt.Println("2. Ver pozo")
				fmt.Scanln(&opcion)
				if opcion == 1 {
					continue
				}else{
					PedirPozo()
				}
			if !juego2(){
				fmt.Println("El juego ha terminado, y tenemos un ganador!!!!!!!!!")
				for i := 0; i < 16; i++ {
					if ListaJugadores[i].Estado == 1 {
						fmt.Println("El ganador es el jugador " + fmt.Sprint(ListaJugadores[i].ID))
					}
				}
				fmt.Println("Selecciona que vas a hacer:")
				fmt.Println("1. Jugar")
				fmt.Println("2. Ver pozo")
				fmt.Scanln(&opcion)
				if opcion == 1 {
					break
				}else{
					PedirPozo()
				}
			}
			fmt.Println("Selecciona que vas a hacer:")
			fmt.Println("1. Jugar")
			fmt.Println("2. Ver pozo")
			fmt.Scanln(&opcion)
			if opcion == 1 {
				continue

			}else{
				PedirPozo()
			}
			if !juego3(){
				fmt.Println("El juego ha terminado, y tenemos un ganador!!!!!!!!!")
				for i := 0; i < 16; i++ {
					if ListaJugadores[i].Estado == 1 {
						fmt.Println("El ganador es el jugador " + fmt.Sprint(ListaJugadores[i].ID))
					}
				}
			}else{
				fmt.Println("El juego ha terminado, y tenemos varios ganadores, filicitaciones a :")
				for i:=0; i<16; i++{
					if ListaJugadores[i].Estado == 1{
						fmt.Println("Jugador " + fmt.Sprint(ListaJugadores[i].ID))
					}
				}
				fmt.Println("Selecciona que vas a hacer:")
				fmt.Println("1. Jugar")
				fmt.Println("2. Ver pozo")
				fmt.Scanln(&opcion)
				if opcion == 1 {
					break
				}else{
					PedirPozo()
				}
			}
		}
		if opcion == 2 {
			PedirPozo()
		}
		fmt.Println("Selecciona que vas a hacer:")
		fmt.Println("1. Jugar")
		fmt.Println("2. Ver pozo")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			break
		}else{
			PedirPozo()
		}
	}
	
}