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

## Directory structure
[see here](https://github.com/golang-standards/project-layout)