package main

import (
	"context"
	//"encoding/csv"
	"fmt"
	//"io"
	"log"
	//"os"
	//"strconv"
	"time"
	"math/rand"
	"reflect"

	pb "github.com/GianniCarlini/Lab-1-SD/chat"
	"google.golang.org/grpc"
)

const (
	address = "10.10.28.68:50051"
)
type PaqueteCola struct{
	id_paquete string
	seguimiento int64
	tipo string
	valor int64
	intentos int64
	estado int64 //0: En bodega / 1: En Camino / 2: Recibido / 3: No Recibido
	fecha string
}

func ordenar(p1 PaqueteCola, p2 PaqueteCola) (PaqueteCola, PaqueteCola){
	if p1.valor < p2.valor {
		return p2,p1
	} else {
		return p1,p2
	}
}
func reparto(r [2]PaqueteCola)[2]PaqueteCola{
	aux := PaqueteCola{estado:1}
	normal := "normal"
	prioritario := "prioritario"
	retail := "retail"
	for iteracion,p := range r{
		if reflect.DeepEqual(p,aux){ //caso base reviso si esta vacio
			continue
		}
		// cantidad de intentos son 3 (retail 3 si o si y pyme 1+2 reintentos si es valor positivo)
		for k := 0 ; k < 4 ; k++{
			var t int
			if k==3{ //paquete en intento numero 4
				p.estado = 3
				p.intentos = int64(k)
				fmt.Println("no me quieren")
				break
			}
			if p.tipo == normal{
				t = 1
			}
			if p.tipo == prioritario{
				t = 2
			}
			if p.tipo == retail{
				t = 3
			}
			fmt.Println(p.tipo)
			if t<=2 && int64(k)*10>p.valor{
				p.estado = 3
				p.intentos = int64(k)
				fmt.Println("No me quieren")
				break
			}
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			exito := r1.Intn(101)

			if exito < 80{
				p.estado = 2
				p.fecha = time.Now().Format("02-01-2006 15:04:05")
				p.intentos = int64(k)
				fmt.Println("Me recibieron")
				break
			}
			
		} 
		r[iteracion]=p
	}
	return r
}

func simulacion(envios [6]PaqueteCola)[6]PaqueteCola{
	package1,package2 := ordenar(envios[0],envios[1]) //ordenos envios camion 1
	package3,package4 := ordenar(envios[2],envios[3])
	package5,package6 := ordenar(envios[4],envios[5])
	//creamos los camiones con sus 2 paquetes asignados
	camionretail1 := [2]PaqueteCola{package1,package2}
	camionretail2 := [2]PaqueteCola{package3,package4}
	camionnormal := [2]PaqueteCola{package5,package6}
	// hacemos la simulacion de camion
	camionretail1 = reparto(camionretail1)
	camionretail2 = reparto(camionretail2)
	camionnormal = reparto(camionnormal)
	resultado := [6]PaqueteCola{camionretail1[0],camionretail1[1], camionretail2[0], camionretail2[1], camionnormal[0], camionnormal[1]} 
	//------------------------aca tengo que escribir los registros-------------------------
	/*var registros13 [][]string
	file3, err := os.OpenFile("rcamion1.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file3.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter := csv.NewWriter(file3)
	Value11 := strconv.FormatInt(camionretail1[0].valor, 10)
	registros11 := []string{camionretail1[0].id_paquete,strconv.FormatInt(camionretail1[0].seguimiento, 10),camionretail1[0].tipo,Value11,strconv.FormatInt(camionretail1[0].intentos, 10),strconv.FormatInt(camionretail1[0].estado, 10),camionretail1[0].fecha}
	registros13 = append(registros13, registros11)
	Value12 := strconv.FormatInt(camionretail1[1].valor, 10)
	registros12 := []string{camionretail1[1].id_paquete,strconv.FormatInt(camionretail1[1].seguimiento, 10),camionretail1[1].tipo,Value12,strconv.FormatInt(camionretail1[1].intentos, 10),strconv.FormatInt(camionretail1[1].estado, 10),camionretail1[1].fecha}
	registros13 = append(registros13, registros12)
	strWrite := registros13
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
	registros11 = nil
	registros12 = nil
	registros13 = nil
	//-------------------------------------------------------------------------------------
	var registros23 [][]string
	file5, err := os.OpenFile("rcamion2.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file5.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter2 := csv.NewWriter(file5)
	Value21 := strconv.FormatInt(camionretail2[0].valor, 10)
	registros21 := []string{camionretail2[0].id_paquete,Value21}
	registros23 = append(registros23, registros21)
	Value22 := strconv.FormatInt(camionretail2[1].valor, 10)
	registros22 := []string{camionretail2[1].id_paquete,Value22}
	registros23 = append(registros23, registros22)
	strWrite2 := registros23
	csvWriter2.WriteAll(strWrite2)
	csvWriter2.Flush()
	registros21 = nil
	registros22 = nil
	registros23 = nil
	//-------------------------------------------------------------------------------------
	var registros33 [][]string
	file2, err := os.OpenFile("rcamion3.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file2.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter3 := csv.NewWriter(file2)
	Value31 := strconv.FormatInt(camionnormal[0].valor, 10)
	registros31 := []string{camionnormal[0].id_paquete,Value31}
	registros33 = append(registros33, registros31)
	Value32 := strconv.FormatInt(camionnormal[1].valor, 10)
	registros32 := []string{camionnormal[1].id_paquete,Value32}
	registros33 = append(registros33, registros32)
	strWrite3 := registros33
	csvWriter3.WriteAll(strWrite3)
	csvWriter3.Flush()
	registros31 = nil
	registros32 = nil
	registros33 = nil*/
	//----------------------------------------------------------------------

	return resultado
}

func main() {
	var colaenvios[6]PaqueteCola //max 6 paquetes para los camiones, cola para asignacion en planta(?)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn err: %v", err)
	}
	
	client := pb.NewPacketClient(conn)
	stream, err := client.Camion(context.Background())
	waitc := make(chan struct{})

	msg := &pb.CamionRequest{
		IdPaquete: "",
		Seguimiento: 0,
		Tipo: "",
		Valor: 0,
		Intentos: 0,
		Estado: -9999,}
	go func() {
		for i := 1; i <= 1; i++ {
			stream.Send(msg)
		}
	}()

	go func() {
		var r [6]PaqueteCola
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
			cont += 1
			if cont > 5{
				cont = 0
				r = simulacion(colaenvios)
				for _, pp := range r{
					m := &pb.CamionRequest{
						IdPaquete: pp.id_paquete,   
						Seguimiento: pp.seguimiento,
						Tipo: pp.tipo,        
						Valor: pp.valor,       
						Intentos: pp.intentos,    
						Estado: pp.estado }
					stream.Send(m)
				}
				stream.Send(msg)
				cont = 0
			}
			time.Sleep(1 * time.Second)
		}
	}()


	<-waitc
	stream.CloseSend()
}
