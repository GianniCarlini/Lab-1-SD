grpc:
	export GO111MODULE=on
	go get github.com/golang/protobuf/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=plugins=grpc:chat helloworld.proto

cliente:
	go run client.go


logistica:
	go run server.go


camion:
	go run camiones.go		


finanza:
	go run rabbit/finanzas.go	