package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"
	"fmt"
	//"io/ioutil"
	"os"
	"encoding/csv"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}
var seguimiento int64 = 0
var seguimiento2 int64 = 0 
var seguimiento3 int64 = 0// variable para el codigo de seguimieto
var registros2 [][]string


type PaqueteCola struct{
	id_paquete string
	seguimiento int64
	tipo string
	valor int64
	intentos int64
	estado int64 //0: En bodega / 1: En Camino / 2: Recibido / 3: No Recibido
}

var colaretail[] PaqueteCola
var colaprioritario[] PaqueteCola
var colanormal[] PaqueteCola

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
	t := time.Now().Format(time.ANSIC)
	var registros []string
	registros = append(registros, t,in.GetId(),in.GetTipo(),in.GetProducto(),Value,in.GetTienda(),in.GetDestino(),seg)
	registros2 = append(registros2, registros)
 //-----------------------------------escribiendo registro-----------------------------
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
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
		id_paquete : in.GetId(),
		seguimiento : seguimiento3,
		tipo : in.GetTipo(),
		valor : in.GetValor(),
		intentos : 0,
		estado : 0,
	}

	if in.GetTipo() == normal {
		colanormal= append(colanormal,reg)
	}else if in.GetTipo() == prioritario{
		colaprioritario= append(colaprioritario,reg)
	}else{
		colaretail= append(colaretail,reg)
	}
	fmt.Println(colanormal)
	fmt.Println(colaprioritario)
	fmt.Println(colaretail)
  //-----------------------------------escribiendo registro-----------------------------

	return &pb.PacketReply{Message: in.GetId(), Nseg: seguimiento3,}, nil
}
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
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
