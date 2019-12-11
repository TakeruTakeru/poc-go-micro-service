# POC-Go-micro-service

## what is this.

- this is web server with some simple micro service.

## Detail

- Language: `Go`
- type: `Web`
- Protocol: `gRPC`
- CI/CD: `Circle CI`
- Server: `linux`
- PaaS or IaaS: `Heroku or GCP`

## For Developers
0. Install Protocol Buffer(Mac) `brew install protobuf`
1. Install gRPC Library: `go get -u google.golang.org/grpc`
2. Install Protocol Buffers plugin: `go get -u github.com/golang/protobuf/protoc-gen-go`
3. Install Protocol Buffers plugin: `go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc`
4. Install Middleware(options): `go get -u github.com/grpc-ecosystem/go-grpc-middleware` `go get -u github.com/sirupsen/logrus`

- Make Interface on project root: `protoc --go_out=plugins=grpc:. ./api/*.proto`
- Make Interface Docs: `protoc --doc_out=html,index.html:./api ./api/*.proto`

- Generate gRPC stub: `protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. path/to/your_service.proto`

- Generate reverse-proxy: `protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. path/to/your_service.proto`

## Testing REST API
1. Initiate gRPC Server: `go run cmd/web/gRPC_client/main.go`
2. Initiate reverse-proxy: `go run cmd/web/restserver/main.go`
3. Test POST method: `curl -X POST http://localhost:8080/v1/example/echo -d '{"message":"this is test"}'`

## Directory structure
[see here](https://github.com/golang-standards/project-layout)