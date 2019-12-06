package main

import (
	"context"
	"log"
	"time"

	echoInterface "github.com/TakeruTakeru/poc-go-micro-service/api/echo_api"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect: %s", err)
	}
	defer connection.Close()

	client := echoInterface.NewEchoServiceClient(connection)

	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.EchoMessage(context, &echoInterface.EchoRequest{Message: "Hello, World"})
	if err != nil {
		log.Println(err)
	}

	log.Printf("Server responsed with: %s", response.GetMessage())
}
