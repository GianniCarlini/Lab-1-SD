package main

import (
  "log"

  "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}
type PaqueteCola struct{
	IdPaquete string
	Seguimiento int64
	Tipo string
	Valor int64
	Intentos int64
	Estado int64 //0: En bodega / 1: En Camino / 2: Recibido / 3: No Recibido
}

var completados int64
var total float64
//saco un paquete de la cola
func CalculoCompletados(paquete PaqueteCola){
	if paquete.Estado == 2{
		completados +=1
		fmt.Println("Me complete")
	}
}
func GananciaPaquete(paquete PaqueteCola){
	var ganancia float64

	valor := float64(paquete.Valor)
	intentos := float64(paquete.Intentos)

	if paquete.Tipo == "normal"{
		if paquete.Estado == 3{
			ganancia = 0
		}else{
			ganancia = valor-((intentos-1)*10)
		}
	}else if paquete.Tipo == "prioritario"{
		if paquete.Estado == 3{
			ganancia = 0.3*valor-((intentos-1)*10)
		}else{
			ganancia = valor-((intentos-1)*10)
		}
	}else if paquete.Tipo == "retail"{
		ganancia = valor-((intentos-1)*10)
	}
	fmt.Println("tengo esto de ganancia: %v",ganancia)
	total += ganancia
}


func main(){
	conn, err := amqp.Dial("amqp://mqadmin:mqadminpassword@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()

q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")

msgs, err := ch.Consume(
	q.Name, // queue
	"",     // consumer
	true,   // auto-ack
	false,  // exclusive
	false,  // no-local
	false,  // no-wait
	nil,    // args
  )
  failOnError(err, "Failed to register a consumer")
  
  forever := make(chan bool)
  
  go func() {
	for d := range msgs {
    log.Printf("Received a message: %s", d.Body)

    var m PaqueteCola

    _ = json.Unmarshal(d.Body, &m)
    CalculoCompletados(aux)
    GananciaPaquete(aux)
    fmt.Println(total)
	}
  }()
  
  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}