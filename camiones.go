package main

import (
	"context"
	//"encoding/csv"
	//"fmt"
	//"io"
	"log"
	//"os"
	//"strconv"
	//"time"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn err: %v", err)
	}
	
	client := pb.NewPacketClient(conn)
	stream, err := client.Camion(context.Background())
	waitc := make(chan struct{})

	msg := &pb.CamionRequest{Status: "Envio a reply"}
	go func() {
		for i := 1; i <= 10; i++ {

			stream.Send(msg)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			log.Printf(resp.Algo)
		}
	}()


	<-waitc
	stream.CloseSend()
}
