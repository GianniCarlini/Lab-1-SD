package main

import (
	"context"
	"log"
	"time"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

const (
	address       = "localhost:50051"
	defaultPacket = "SA5698PO,Televisor,110,tienda-C,casa-C"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPacketClient(conn)

	// Contact the server and print out its response.
	packet := defaultPacket
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendPacket(ctx, &pb.PacketRequest{Packet: packet})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r.GetMessage())
}
