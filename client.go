package main

import (
	"context"
	"encoding/csv"

	"io"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

type Retail struct {
	id       string
	producto string
	valor    int64
	tienda   string
	destino  string
}

const (
	address = "localhost:50051"
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
	file, err := os.Open("csv/retail.csv")
	if err != nil {
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','         //Como se separan las columnas
	reader.FieldsPerRecord = 5 //Cuanta columnas hay, si no se valor = -1

	var rawData []Retail
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error al leer: %v", err)
		}
		c := Retail{
			id:       record[0],
			producto: record[1],
			tienda:   record[3],
			destino:  record[4],
		}
		if record[2] != "" {
			i, err := strconv.ParseInt(record[2], 10, 64)
			if err != nil {
				log.Printf("Intentando procesar la edad: %v", err)
				continue
			}
			c.valor = i
		}
		rawData = append(rawData, c)
	}

	for i := 0; i < len(rawData); i++ {
		id := rawData[i].id
		producto := rawData[i].producto
		valor := rawData[i].valor
		tienda := rawData[i].tienda
		destino := rawData[i].destino

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.SendPacket(ctx, &pb.PacketRequest{Id: id, Producto: producto, Valor: valor, Tienda: tienda, Destino: destino}) 
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("%s", r.GetMessage())
	}
}
