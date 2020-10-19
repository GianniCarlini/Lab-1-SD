package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"
	"fmt"
	"reflect"
	"os"
	"encoding/csv"
	"encoding/json"

	"github.com/streadway/amqp"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
	port2 = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}
var seguimiento int64 = 0 // variable para el codigo de seguimieto
var seguimiento2 int64 = 0 // variable para el codigo de seguimieto
var seguimiento3 int64 = 0// variable para el codigo de seguimieto
var registros2 [][]string // registro en memoria de paquetes


type PaqueteCola struct{
	IdPaquete string
	Seguimiento int64
	Tipo string
	Valor int64
	Intentos int64
	Estado int64 //0: En bodega / 1: En Camino / 2: Recibido / 3: No Recibido
}

var finanzas[] PaqueteCola
var colaretail[] PaqueteCola
var colaprioritario[] PaqueteCola
var colanormal[] PaqueteCola
//--------------funcion de internet para borrar --------------------------------------------
func Remove(s []PaqueteCola, index int) []PaqueteCola {
	return append(s[:index], s[index+1:]...)
}
//-------------------funcion buscar origen destino para enviar a camiones-------------------
func Buscar(id string)(string, string){
	for _,i:= range registros2{
		if id==i[1]{
			return i[5],i[6]
		}
	}
	return "",""
}
//------------------------------ generar paquetes para camiones-------------------------------
func GenerarEnvio() [6]PaqueteCola {
	var enviocamion [6]PaqueteCola //6 paquetes que se envian al camion
	aux := PaqueteCola{} //paquete vacio para hacer comparacion
	
	for _,p:= range colaretail{
		if (!(reflect.DeepEqual(enviocamion[3],aux))){ //compara si la cola esta vacia hasta los primeros 4 paquetes para camiones 1 y 2
			break
		}
		for i := 0; i < 4 ; i++{ //sino busco en los primeros 4 paquetes
			if reflect.DeepEqual(enviocamion[i],p){
				break
			}
			if reflect.DeepEqual(enviocamion[i],aux){
				enviocamion[i]=p
				break
			}
		}
	}

	for _,p:= range colaprioritario{
		if (!(reflect.DeepEqual(enviocamion[5],aux))){ //compara si la cola esta vacia hasta los ultimos 2 paquetes 
				break
		}
		for j := 0 ; j < 6 ; j++{
			if reflect.DeepEqual(enviocamion[j],p){
				break
			}
			if reflect.DeepEqual(enviocamion[j],aux){
				enviocamion[j]=p
				break
			}
		}
		
	}

	for _,p:= range colanormal{
		if (!(reflect.DeepEqual(enviocamion[5],aux))){ //compara si la cola esta vacia hasta los ultimos 2 paquetes 
				break
		}
		for k := 4 ; k < 6 ; k++{
			if reflect.DeepEqual(enviocamion[k],p){
				break
			}
			if reflect.DeepEqual(enviocamion[k],aux){
				enviocamion[k]=p
				break
			}
		}
	}
	return enviocamion
}
//-------------------------------------------server logistica---------------------------------

func (s *server) SendPacket(ctx context.Context, in *pb.PacketRequest) (*pb.PacketReply, error) {
	log.Printf("Peticion: %v, Producto: %v", in.GetId(), in.GetProducto())
	normal := "normal"
	prioritario := "prioritario"
	if in.GetTipo() == normal {
		seguimiento += 1
		seguimiento3 = seguimiento
	}else if in.GetTipo() == prioritario{
		seguimiento += 1
		seguimiento3 = seguimiento
	}else{
		seguimiento3 = seguimiento2
	}
	seg := strconv.Itoa(int(seguimiento3))
	Value := strconv.FormatInt(in.GetValor(), 10)
	t := time.Now().Format("02-01-2006 15:04:05")
	var registros []string
	registros = append(registros, t,in.GetId(),in.GetTipo(),in.GetProducto(),Value,in.GetTienda(),in.GetDestino(),seg) //registros de logistica en memoria
	registros2 = append(registros2, registros) //arreglo 2D para guardar todos los registros
 //-----------------------------------escribiendo registro-----------------------------
	file, err := os.OpenFile("retail.csv", os.O_CREATE|os.O_WRONLY, 0777)
    defer file.Close()
 
    if err != nil {
        os.Exit(1)
    }
    csvWriter := csv.NewWriter(file)
    strWrite := registros2
    csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
  //---------------------------------------------colas----------------------------------
	reg := PaqueteCola{
		IdPaquete : in.GetId(),
		Seguimiento : seguimiento3,
		Tipo : in.GetTipo(),
		Valor : in.GetValor(),
		Intentos : 0,
		Estado : 0,
	}

	if in.GetTipo() == normal {
		colanormal= append(colanormal,reg)
	}else if in.GetTipo() == prioritario{
		colaprioritario= append(colaprioritario,reg)
	}else{
		colaretail= append(colaretail,reg)
	}
  //-----------------------------------------------------------------------------------

	return &pb.PacketReply{Message: in.GetId(), Nseg: seguimiento3,}, nil //envio respuesta al cliente
}
 //-----------------------------------servidor camiones---------------------------------
 func (s *server) Camion(stream pb.Packet_CamionServer) error {
	contador := 0
	entrega := [6]PaqueteCola{} // array de entregas para recepcion de la entrega del paqute
	for {
		r, err := stream.Recv()
		if err != nil {
			log.Fatalf("RPC failed: %v", err)
		}
		// aca va logica si el paquete es nulo continuo sino borro las colas
		if r.Estado != int64(-9999){ //si el paquete que envia camiones no es "fantasma" o estado -9999 realizo el borrado
			fmt.Println("Paquetes recividos!")
			//maximo 6 paquetes
			entrega[contador] = PaqueteCola{
				IdPaquete: r.IdPaquete,   
				Seguimiento: r.Seguimiento,
				Tipo: r.Tipo,        
				Valor: r.Valor,       
				Intentos: r.Intentos,    
				Estado: r.Estado,	
			}
			contador += 1
			if contador == 6{//cuando recepcione los 6 paquetes
				normal := "normal"
				prioritario := "prioritario"
				retail := "retail"
				// borrado en los registros de colas de entregas completadas
				for _,paquete := range entrega{
					if reflect.DeepEqual(paquete,PaqueteCola{Estado:1}){ //reviso si estado es 1 "en camino" para no tener errores
						continue
					}else{
					if paquete.Tipo == normal{
						for i,paquetecolanormal := range colanormal{
							if paquete.IdPaquete == paquetecolanormal.IdPaquete{
								//borro
								colanormal = Remove(colanormal,i)
								break
							}
						}
					}
					if paquete.Tipo == prioritario{
						for i,paquetecolaprioritario := range colaprioritario{
							if paquete.IdPaquete == paquetecolaprioritario.IdPaquete{
								colaprioritario = Remove(colaprioritario,i)
								break
							}
						}
					}
					if paquete.Tipo == retail{
						for i,paquetecolaretail := range colaretail{
							if paquete.IdPaquete == paquetecolaretail.IdPaquete{
								colaretail = Remove(colaretail,i)
								break
							}
						}
					}
						Rabbit(paquete)
						finanzas = append(finanzas,paquete)
					}
				}
				contador = 0
			}
		}else{ //si el paquete es fantasma genero los 6 paquetes para los camiones
			envio := GenerarEnvio() //genero los paquetes a enviar
			fmt.Println("Enviando paquetes a reparto!")
			for _, paquetecamion:= range envio{
				resp := pb.CamionReply{
					IdPaquete: paquetecamion.IdPaquete,
					Seguimiento: paquetecamion.Seguimiento,
					Tipo: paquetecamion.Tipo,
					Valor: paquetecamion.Valor,
					Intentos: paquetecamion.Intentos,
					Estado: 1,
				}
				if err := stream.Send(&resp); err != nil { //envio los paquetes a camiones
					log.Printf("send error %v", err)
				}
				
			}


		}

	}
	
	
}
func (s *server) Od(ctx context.Context, msg *pb.OdRequest) (*pb.OdReply, error){

	id := msg.GetId()

	var ori string
	var dest string

	if id != "" {
		ori, dest = Buscar(id)
	}
	return &pb.OdReply {
		Origen: ori,
		Destino: dest,
	}, nil

}
func (s *server) Seguimiento(ctx context.Context, in *pb.SeguimientoRequest) (*pb.SeguimientoReply, error) {
	log.Printf("Received: %v", in.GetCodigo())
	esta2 := "xd"
	for _,codigo := range finanzas{
		if in.GetCodigo() == codigo.Seguimiento{
			//0: En bodega / 1: En Camino / 2: Recibido / 3: No Recibido
			if codigo.Estado == 0{
				esta2 = "En bodega"
			}else if codigo.Estado == 1{
				esta2 = "En Camino"
			}else if codigo.Estado == 2{
				esta2 = "Recibido "
			}else{
				esta2 = "No Recibido"
			}
			
		}
	}
	for _,codigo := range colanormal{
		if in.GetCodigo() == codigo.Seguimiento{
			//0: En bodega / 1: En Camino / 2: Recibido / 3: No Recibido
			if codigo.Estado == 0{
				esta2 = "En bodega"
			}else if codigo.Estado == 1{
				esta2 = "En Camino"
			}else if codigo.Estado == 2{
				esta2 = "Recibido "
			}else{
				esta2 = "No Recibido"
			}
		}
	}
	for _,codigo := range colaprioritario{
		if in.GetCodigo() == codigo.Seguimiento{
			//0: En bodega / 1: En Camino / 2: Recibido / 3: No Recibido
			if codigo.Estado == 0{
				esta2 = "En bodega"
			}else if codigo.Estado == 1{
				esta2 = "En Camino"
			}else if codigo.Estado == 2{
				esta2 = "Recibido "
			}else{
				esta2 = "No Recibido"
			}
		}
	}
	if esta2 == "xd"{
		esta2 = "No esta en sistema"
	}

	return &pb.SeguimientoReply{Estadoseguimiento: esta2,}, nil
}
func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func Rabbit(){
	conn, err := amqp.Dial("amqp://mqadmin:mqadminpassword@10.10.28.7:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()		

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

		body, _ := json.Marshal(p)

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing {
			  ContentType: "application/json",
			  Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")

		fmt.Println("Successfully Published Message to Queue")
}

func main() {
	fmt.Println("Bienvenido al servidor de logistica de PrestigioExpress")
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterPacketServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	//------------------------------------------------------


}
