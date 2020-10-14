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
	port2 = ":9000"
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

//------------------------------ generar paquetes para camiones-------------------------------
/*func GenerarEnvio(){
	var enviocamion [6]Registro //6 paquetes que se envian al camion
	aux := PaqueteCola{} //paquete vacio para hacer comparacion
	
	for _,p:= range colaretail{
		if (!(reflect.DeepEqual(enviocamion[3],void))){
			break
		}
		for i := 0; i < 4 ; i++{
			if reflect.DeepEqual(enviocamion[i],p){
				break
			}
			if reflect.DeepEqual(enviocamion[i],void){
				enviocamion[i]=p
				break
			}
		}
	}

	for _,p:= range colaprioritario{
		if (!(reflect.DeepEqual(enviocamion[5],void))){
				break
		}
		for j := 0 ; j < 6 ; j++{
			if reflect.DeepEqual(p[j],p){
				break
			}
			if reflect.DeepEqual(enviocamion[j],void){
				enviocamion[j]=p
				break
			}
		}
		
	}

	for _,p:= range colanormal{
		if (!(reflect.DeepEqual(enviocamion[5],void))){
				break
		}
		for k := 4 ; k < 6 ; k++{
			if reflect.DeepEqual(enviocamion[k],p){
				break
			}
			if reflect.DeepEqual(enviocamion[k],void){
				enviocamion[k]=p
				break
			}
		}
	}
	return enviocamion
}*/
//-------------------------------------------server logistica---------------------------------

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
	fmt.Println(colanormal)  //borrar
	fmt.Println(colaprioritario) //borrar
	fmt.Println(colaretail) //borrar
  //-----------------------------------escribiendo registro-----------------------------

	return &pb.PacketReply{Message: in.GetId(), Nseg: seguimiento3,}, nil
}
 //-----------------------------------servidor camiones---------------------------------
 func (s *server) Camion(stream pb.Packet_CamionServer) error {
	
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Fatalf("RPC failed: %v", err)
		}
		fmt.Println(req)
		resp := pb.CamionReply{Algo: "Envio a request"}
		if err := stream.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}

	}
	
	
}

func main() {
	fmt.Println("Bienvenido al servidor de logistica de PrestigioExpress")
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
