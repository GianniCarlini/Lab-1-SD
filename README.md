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

Primero se debe inciar el servidor de logistica en la maquina 2 con el comando make logistica o en caso de no funcionar "go run server.go", una ves iniciado se puede proceder a iniciar el cliente con make cliente en la maquina 1 o en su defecto si falla "go run client.go", a su ves se puede iniciar camiones aunque sus ordener empezaran a enviarce solo si hay peticiones del cliente.

Finanzas al ser asincronico le llegaran las peticiones y al iniciar el codigo make finanza se inciara el sistema si llega a fallar el make "go run finanzas.go" desde la carpeta rabbit empezara a hacer los calculos pedidos.

El sistema cliente tiene 3 opciones default las primeras 2 son de comportamiento con los archivos de ejemplo fijos y la 3ra es para enviar un codigo de seguimiento, este codigo de seguimiento es por teclado y solo un numero.

El sistema camion tiene la opcion al inicio de inciar el tiempo.

Finanzas se termina con CTRL+C al igual que los demas servicios.

El sistema se basa en que el servidor de logistica toma peticiones del cliente y las envia a colas para ser pasadas a camion, luego cuando camion responde con los estados le envia esa informacion a finanzas.

Las maquinas 2 y 4 se les instalo go 1.15 para que los codigos de rabbitmq funcionaran.
En caso de no funcionar go:
    - usar export PATH=$PATH:/usr/local/go/bin
    - o el caso extremo:-export GOROOT=/usr/local/go
                        -export GOPATH=$HOME/go
                        -export GOBIN=$GOPATH/bin
                        -export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

Si desea compilar el proto es con make grpc o en su defecto con los comandos:
    -export GO111MODULE=on
    -go get github.com/golang/protobuf/protoc-gen-go
    -go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0
    -export PATH="$PATH:$(go env GOPATH)/bin"
    -protoc --go_out=plugins=grpc:chat helloworld.proto
