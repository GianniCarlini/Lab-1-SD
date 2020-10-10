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
type Pymes struct {
	id       string
	producto string
	valor    int64
	tienda   string
	destino  string
	tipo 	 int64
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

 //------------------------- Lectura archivo retail --------------------------------------------
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
				log.Printf("Intentando procesar el valor: %v", err)
				continue
			}
			c.valor = i
		}
		rawData = append(rawData, c)
	}
 //------------------------- Envio paquete retail--------------------------------------------
	for i := 0; i < len(rawData); i++ {
		id := rawData[i].id
		producto := rawData[i].producto
		valor := rawData[i].valor
		tienda := rawData[i].tienda
		destino := rawData[i].destino
		tipo := "retail"

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.SendPacket(ctx, &pb.PacketRequest{Id: id, Producto: producto, Valor: valor, Tienda: tienda, Destino: destino, Tipo: tipo}) 
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("%s", r.GetMessage())
	}
 //------------------------- Lectura archivo Pymes --------------------------------------------
	file2, err := os.Open("csv/pymes.csv")
	if err != nil {
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer file2.Close()

	reader2 := csv.NewReader(file2)
	reader2.Comma = ','         //Como se separan las columnas
	reader2.FieldsPerRecord = 6 //Cuanta columnas hay, si no se valor = -1

	var rawData2 []Pymes
	for {
		record2, err := reader2.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error al leer: %v", err)
		}
		c2 := Pymes{
			id:       record2[0],
			producto: record2[1],
			tienda:   record2[3],
			destino:  record2[4],
		}
		if record2[2] != "" {
			i, err := strconv.ParseInt(record2[2], 10, 64)
			if err != nil {
				log.Printf("Intentando procesar el valor: %v", err)
				continue
			}
			c2.valor = i
		}
		if record2[5] != "" {
			i, err := strconv.ParseInt(record2[5], 10, 64)
			if err != nil {
				log.Printf("Intentando procesar el tipo: %v", err)
				continue
			}
			c2.tipo = i
		}
		rawData2 = append(rawData2, c2)
	}
 //------------------------- Envio paquete Pymes--------------------------------------------
	for i := 0; i < len(rawData2); i++ {
		id := rawData2[i].id
		producto := rawData2[i].producto
		valor := rawData2[i].valor
		tienda := rawData2[i].tienda
		destino := rawData2[i].destino
		var t string
		if rawData2[i].tipo == 0{ //suponemos 0 = normal, 1 = prioritario
			t = "normal"
		}else {
		t = "prioritario"
		}
		tipo := t

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.SendPacket(ctx, &pb.PacketRequest{Id: id, Producto: producto, Valor: valor, Tienda: tienda, Destino: destino, Tipo: tipo}) 
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("%s", r.GetMessage())
	}
}
