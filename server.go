package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"
	"fmt"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}
var seguimiento = 0 // variable para el codigo de seguimieto
var registros [][8]string

func (s *server) SendPacket(ctx context.Context, in *pb.PacketRequest) (*pb.PacketReply, error) {
	log.Printf("Peticion: %v, Producto: %v", in.GetId(), in.GetProducto())
	seguimiento += 1
	seg := strconv.Itoa(seguimiento)
	Value := strconv.FormatInt(in.GetValor(), 10)
	t := time.Now().Format(time.ANSIC)
	registro := [8]string{t, in.GetId(), "tipo", in.GetProducto(), Value, in.GetTienda(), in.GetDestino(), seg}
	registros = append(registros, registro)
	fmt.Println(registro)
	return &pb.PacketReply{Message: "El paquete " + in.GetId()+ " se recibio correctamente con id de seguimiento: "+ seg}, nil
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
}
