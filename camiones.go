package main

import (
	"context"
	//"encoding/csv"
	"fmt"
	//"io"
	"log"
	//"os"
	//"strconv"
	//"time"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)
type PaqueteCola struct{
	id_paquete string
	seguimiento int64
	tipo string
	valor int64
	intentos int64
	estado int64 //0: En bodega / 1: En Camino / 2: Recibido / 3: No Recibido
}

func ordenar(p1 PaqueteCola, p2 PaqueteCola) (PaqueteCola, PaqueteCola){
	if p1.valor < p2.valor {
		return p2,p1
	} else {
		return p1,p2
	}
}
/*func reparto{

}

func simulacion{

}*/

func main() {
	var colaenvios[6]PaqueteCola //max 6 paquetes para los camiones, cola para asignacion en planta(?)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn err: %v", err)
	}
	
	client := pb.NewPacketClient(conn)
	stream, err := client.Camion(context.Background())
	waitc := make(chan struct{})

	msg := &pb.CamionRequest{Status: "Envio a reply"}
	go func() {
		for i := 1; i <= 1; i++ {
			stream.Send(msg)
		}
	}()

	go func() {
		cont := 0
		for {
			resp, err := stream.Recv()
			paquetazo := PaqueteCola{
				id_paquete: resp.IdPaquete,
				seguimiento: resp.Seguimiento,
				tipo: resp.Tipo,
				valor: resp.Valor,
				intentos: resp.Intentos,
				estado: resp.Estado,
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			colaenvios[cont] = paquetazo
			fmt.Println(colaenvios)
			cont += 1
			if cont > 5{
				cont = 0
			}
		}
	}()


	<-waitc
	stream.CloseSend()
}
