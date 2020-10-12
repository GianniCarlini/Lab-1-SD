package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"
	//"fmt"
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
  //-----------------------------------escribiendo registro-----------------------------

	return &pb.PacketReply{Message: in.GetId(), Nseg: seguimiento3,}, nil
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
