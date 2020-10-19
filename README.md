# Lab-1-SD
Laboratorio numero 1 de sistemas distribuidos Golang-grpc-RabbitMQ
Gianni Carlini 201773105-2

Client.go Pertenece a la logica del Cliente
server.go Pertenece a la logica del Logistica
camiones.go Pertenece a la logica del Camiones
En la carpeta rabbit se encuentra finanzas.go que pertenece a la logica del Finanzas

El .proto para correr el servidor de logistica esta compilado

Los tiempo ingresados al inicio del sistema son en segundos y solo debe mandar 1 numero desde teclado

Máquina 1 (10.10.28.67) - cliente
Máquina 2 (10.10.28.68) - logistica
Máquina 3 (10.10.28.69) - camiones
Máquina 4 (10.10.28.7)  - finanzas

Primero se debe inciar el servidos de logistica con el comando "go run server.go", una ves iniciado se puede proceder a iniciar el cliente "go run client.go"
por otro se puede iniciar camiones aunque sus ordener empezaran cuando lleguen peticiones desde cliente.

Finanzas al ser asincronico le llegaran las peticiones y al iniciar el codigo "go run finanzas.go" desde la carpeta rabbit empezara a hacer los calculos pedidos.

El sistema cliente tiene 3 opciones default las primeras 2 son de comportamiento con los archivos de ejemplo fijos y la 3ra es para enviar un codigo de seguimiento, este codigo de seguimiento es por teclado y solo un numero.

Las maquinas 2 y 4 se les instalo go 1.15 para que los codigos de rabbitmq funcionaran.
En caso de no funcionar go:
    - usar export PATH=$PATH:/usr/local/go/bin
    - o el caso extremo:-export GOROOT=/usr/local/go
                        -export GOPATH=$HOME/go
                        -export GOBIN=$GOPATH/bin
                        -export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
