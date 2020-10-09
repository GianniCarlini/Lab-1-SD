package main

import (
	"context"
	"log"
	"net"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *server) SendPacket(ctx context.Context, in *pb.PacketRequest) (*pb.PacketReply, error) {
	log.Printf("%v", in.GetPacket())
	return &pb.PacketReply{Message: "Su packete es:" + in.GetPacket()}, nil
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
