package main

import (
	"context"
	"fmt"
	"log"
	"time"

	echoInterface "github.com/TakeruTakeru/poc-go-micro-service/api/echo_api"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err.Error())
	}
	defer connection.Close()

	client := echoInterface.NewEchoServiceClient(connection)

	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.EchoMessage(context, &echoInterface.EchoRequest{Message: "Hello, World"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Server responsed with: %s", response.GetMessage())
}
