package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"strconv"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)
type server struct{
	pb.UnimplementedLiderServiceServer
}
type Jugada struct {
	ID int32
	Jugada int32
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Peticion recibida, aceptando juego")
	return &pb.HelloReply{Message: "Juego aceptado"}, nil
}

func (s *server) SendJugada(ctx context.Context, in *pb.Jugada) (*pb.Resultado, error) {
	// Enviarla a NameNode
	fmt.Println("Jugadas recibidas, jugada:" + fmt.Sprint(in.GetJugada()) + " del jugador: " + fmt.Sprint(in.GetID()))
	conn, err := grpc.Dial("10.6.40.219:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el server NameNode: %v",err)
	}
	defer conn.Close()

	//Se crea un cliente para la conexion
	serviceClient := pb.NewNameNodeServiceClient(conn)
	//Se envia la jugada para que sea escrita en el archivo
	_, err = serviceClient.SendJugada(context.Background(), &pb.Jugada{Jugada: in.GetJugada(), ID: in.GetID()})
	if err != nil {
		log.Fatalf("No se pudo enviar la jugada: %v",err)
	}
	//Lider debe elegir un numero entre 6 y 10
	JugadaLider := rand.Intn(4) + 6
	if JugadaLider <= int(in.GetJugada()){
		//Notificar a pozo que alguien murio por rabbitmq
		sender(strconv.Itoa(int(in.GetID())))
		return &pb.Resultado{ID: in.GetID(),Estado: 0}, nil
	}
	return &pb.Resultado{ID: in.GetID(),Estado: 1}, nil
}

func sum(array []int32) int32 { 
	var sum int32
	sum = 0
	for _, v := range array {  
		sum += v  
	}  
	return sum
   }  

func (s *server) SendJugada2(ctx context.Context, in *pb.Jugada2) (*pb.Resultado, error) {
	//Recibir las jugadas
	ID1 := in.GetID1()
	ID2 := in.GetID2()
	T1 := sum(in.GetJugada1())
	T2 := sum(in.GetJugada2())
	//Enviar jugadas a NameNode


	JugadaLider := (rand.Intn(4) + 1) % 2
	var resultado [2]int32
	if JugadaLider == 0{
		if  T1%2 == 0 && T2%2 == 0{
			resultado = [2]int32{1,1} 
		}else if T1%2 == 1 && T2%2 == 1{
			Decididor := rand.Intn(1)
			if Decididor == 0{
				resultado = [2]int32{1,0}
			}else{
				resultado = [2]int32{0,1}
			}
		}else if T1%2 == 0 && T2%2 == 1{
			resultado = [2]int32{0,1}
		}else{
			resultado = [2]int32{1,0}
		}
	}else{
		if  T1%2 == 1 && T2%2 == 1{
			resultado = [2]int32{1,1}
		}else if T1%2 == 0 && T2%2 == 0{
			Decididor := rand.Intn(1)
			if Decididor == 0{
				resultado = [2]int32{1,0}
			}else{
				resultado = [2]int32{0,1}
			}
		}else if T1%2 == 0 && T2%2 == 1{
			resultado = [2]int32{0,1}
		}else{
			resultado = [2]int32{1,0}
		}
	}
	//Enviar las jugadas a NameNode
	conn, err := grpc.Dial("10.6.40.219:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el server NameNode: %v",err)
	}
	defer conn.Close()

	//Se crea un cliente para la conexion
	serviceClient := pb.NewNameNodeServiceClient(conn)
	//Se envia la jugada para que sea escrita en el archivo
	_, err = serviceClient.SendJugada2(context.Background(), &pb.Jugada2{ID1: ID1, ID2: ID2, Jugada1: in.GetJugada1(), Jugada2: in.GetJugada2()})
	if err != nil {
		log.Fatalf("No se pudo enviar la jugada: %v",err)
	}
	if resultado[0] == 0{
		for i:=0; i<len(ID1); i++{
			sender(strconv.Itoa(int(ID1[i])))
		}
	if resultado[1] == 0{
		for i:=0; i<len(ID2); i++{
			sender(strconv.Itoa(int(ID2[i])))
		}
	}
	}
	return &pb.Resultado{ID: resultado[0],Estado: resultado[1]}, nil
}

func (s *server) SendJugada3 (ctx context.Context, in *pb.Jugada3) (*pb.Resultado, error) {
	//Recibir las jugadas
	ID1 := in.ID1
	ID2 := in.ID2
	jugada1 := in.Jugada1
	jugada2 := in.Jugada2
	//Enviar jugadas a NameNode
	conn, err := grpc.Dial("10.6.40.219:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el server NameNode: %v",err)
	}
	defer conn.Close()

	//Se crea un cliente para la conexion
	serviceClient := pb.NewNameNodeServiceClient(conn)
	//Se envia la jugada para que sea escrita en el archivo
	_, err = serviceClient.SendJugada3(context.Background(), &pb.Jugada3{ID1: ID1, ID2: ID2, Jugada1: jugada1, Jugada2: jugada2})
	if err != nil {
		log.Fatalf("No se pudo enviar la jugada: %v",err)
	}
	//Lider debe elegir un numero entre 1 y 10
	JugadaLider := rand.Intn(10) + 1
	//Calcular distancias con el lider
	Distancia1 := math.Abs(float64(JugadaLider - int(jugada1)))
	Distancia2 := math.Abs(float64(JugadaLider - int(jugada2)))
	//Decidir al ganador
	if Distancia1 < Distancia2{
		sender(strconv.Itoa(int(ID2)))
		return &pb.Resultado{ID: ID1,Estado: 1}, nil
	}
	sender(strconv.Itoa(int(ID1)))
	return &pb.Resultado{ID: ID2,Estado: 1}, nil

}


func (s *server) RequestPozo(ctx context.Context, in *pb.RequestPozoActual) (*pb.ResponsePozoActual, error) {
	// Enviarla a Pozo
	fmt.Println("Peticion recibida, enviando peticion al servidor Pozo")
	conn, err := grpc.Dial("10.6.40.220:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el server Pozo: %v",err)
	}
	defer conn.Close()

	//Se crea un cliente para la conexion
	serviceClient := pb.NewPozoServiceClient(conn)
	//Se envia la peticion al servidor pozo
	res, err := serviceClient.RequestPozo(context.Background(), &pb.RequestPozoActual{Pozo: 1})
	if err != nil {
		log.Fatalf("No se pudo enviar la peticion: %v",err)
	}
	return res, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sender(ID string) {
	conn, err := amqp.Dial("amqp://test:test@10.6.40.220:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Canal1", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	err = ch.Publish(
	"",     // exchange
	q.Name, // routing key
	false,  // mandatory
	false,  // immediate
	amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(ID),
	})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", ID)
}


func ServidorCliente(){
	fmt.Println("Servidor Lider Iniciado")
	listener , err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("No se pudo iniciar el server: %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterLiderServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("No se pudo iniciar el server: %v",err)
	}
}

func main(){
	ServidorCliente()
	
}