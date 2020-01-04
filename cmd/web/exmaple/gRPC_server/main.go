package main

import (
	"log"
	"net"
	"os"
	"time"

	echoInterface "github.com/TakeruTakeru/poc-go-micro-service/api/echo_api"
	echoService "github.com/TakeruTakeru/poc-go-micro-service/internal/app/sample_service"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalln(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{})

	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	grpc_logrus.ReplaceGrpcLogger(logger)

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logger, opts...),
		),
	)
	service := echoService.NewEchoService()

	echoInterface.RegisterEchoServiceServer(server, service)

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
