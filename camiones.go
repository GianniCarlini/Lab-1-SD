package main

import (
	"context"
	"encoding/csv"
	"fmt"
	//"io"
	"log"
	"os"
	"strconv"
	"time"
	"math/rand"
	"reflect"

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
	fecha string
}
var camion1 [][] string
var camion2 [][] string
var camion3 [][] string

//-----------------------------funcion para ordenar los paquetes de cada camion por valor---------------------------
func ordenar(p1 PaqueteCola, p2 PaqueteCola) (PaqueteCola, PaqueteCola){ 
	if p1.valor < p2.valor {
		return p2,p1
	} else {
		return p1,p2
	}
}
//----------------------------------funcion que realiza la logica de reparto---------------------------------------
func reparto(r [2]PaqueteCola)[2]PaqueteCola{ 
	aux := PaqueteCola{estado:1} //aux para verificar errores en paquetes
	normal := "normal"
	prioritario := "prioritario"
	retail := "retail"
	for iteracion,p := range r{
		if reflect.DeepEqual(p,aux){ //caso base reviso si esta vacio (estado en camino)
			continue
		}
		// cantidad de intentos son 3 (retail 3 si o si y pyme 1+2 reintentos si es valor positivo)
		for k := 0 ; k < 4 ; k++{
			var t int
			if k==3{ //paquete en intento numero 4
				p.estado = 3
				p.intentos = int64(k)
				p.fecha = 0
				fmt.Println("No me quieren")
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
			if t<=2 && int64(k)*10>p.valor{ //condiciones para pyme
				p.estado = 3
				p.intentos = int64(k)
				p.fecha = 0
				fmt.Println("No me quieren")
				break
			}
			//abajo genero numero random entre 0-100 para el  80%
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			exito := r1.Intn(101)

			if exito < 80{ //si el numero es entre 0-80 exito en la entrega
				p.estado = 2
				p.fecha = time.Now().Format("02-01-2006 15:04:05")
				p.intentos = int64(k)
				fmt.Println("Me recibieron")
				break
			}
			
		} 
		r[iteracion]=p
	}
	return r //retorno el mismo camion(array) que entro
}

//------------------------------- funcion para generar los registros por camion-------------------------------
func registrocamion(entrega [6]PaqueteCola){
	conn, err := grpc.Dial(address, grpc.WithInsecure()) //coneccion grpc para pedir el origen y destino ya que no lo pedian en pdf para la primera conexion
		
		if err != nil {
			log.Fatalf("Conn err: %v", err)
		}
		c := pb.NewPacketClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

	for i,p := range entrega{ //itero sobre los 6 paquetes entregados
		var nombre string

		r, err := c.Od(ctx, &pb.OdRequest{Id: p.id_paquete}) //peticion con id paquete
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Println(r.GetOrigen())
		fmt.Println(r.GetDestino())
		reg:= []string{p.id_paquete,p.tipo,strconv.FormatInt(p.valor,10), r.GetOrigen(), r.GetDestino(),strconv.FormatInt(p.intentos,10),p.fecha} //creo array registro

	if i == 0 || i == 1{ //casos camion 1
		nombre = "rcamion1.csv"
		camion1=append(camion1,reg)
		file,err:= os.OpenFile(nombre,os.O_CREATE|os.O_WRONLY,0777)
		defer file.Close()
		if err !=nil{
			os.Exit(1)
		}
		csvWriter:= csv.NewWriter(file)
		csvWriter.WriteAll(camion1)
		csvWriter.Flush()
	}else if i == 2 || i == 3 { //casos camion 2
		nombre = "rcamion2.csv"
		camion2=append(camion2,reg)
		file,err:= os.OpenFile(nombre,os.O_CREATE|os.O_WRONLY,0777)
		defer file.Close()
		if err !=nil{
			os.Exit(1)
		}
		csvWriter:= csv.NewWriter(file)
		csvWriter.WriteAll(camion2)
		csvWriter.Flush()
	}else if i == 4 || i == 5 { //casos camion 3
		nombre ="rcamion3.csv"
		camion3=append(camion3,reg)
		file,err:= os.OpenFile(nombre,os.O_CREATE|os.O_WRONLY,0777)
		defer file.Close()
		if err !=nil{
			os.Exit(1)
		}
		csvWriter:= csv.NewWriter(file)
		csvWriter.WriteAll(camion3)
		csvWriter.Flush()
		}
	}
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
	// array con las 6 entregas completadas
	entrega := [6]PaqueteCola{camionretail1[0],camionretail1[1], camionretail2[0], camionretail2[1], camionnormal[0], camionnormal[1]} 
	//hago el registro por camion en csv
	registrocamion(entrega)

	return entrega //retorno un viaje completado de los 3 camiones
}

func main() {
	var colaenvios[6]PaqueteCola //max 6 paquetes para los camiones, cola para asignacion en planta(?)
//--------------------------------------conexion grpc---------------------------------------
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

//----------------- go routine para enviar paquete fantasma-------------------------
	go func() {
		for i := 1; i <= 1; i++ {
			stream.Send(msg)
		}
	}()
//-------- goruoutine para enviar paquetes completados y hacer simulacion------------
	go func() {
		var r [6]PaqueteCola
		cont := 0
		for {
			resp, err := stream.Recv() //recivo los paquetes desde logistica
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
			colaenvios[cont] = paquetazo //creo una cola de envios para realizar la simulacion
			cont += 1
			if cont > 5{ // me llegaron 6 paquetes hago similacion
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
					stream.Send(m) //envio resultados a logisitca
				}
				stream.Send(msg) //mando paquete fantasma para que me envien mas paquetes a repartir
				cont = 0
			}
			time.Sleep(1 * time.Second)
		}
	}()
	<-waitc
	stream.CloseSend()
}
