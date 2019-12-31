package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	auth "github.com/TakeruTakeru/poc-go-micro-service/api/auth"
	_ "github.com/TakeruTakeru/poc-go-micro-service/pkg/logger"
	_ "github.com/TakeruTakeru/poc-go-micro-service/configs"
	fileInterface "github.com/TakeruTakeru/poc-go-micro-service/api/fileservice"
	fileService "github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
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
			// 認証、認可の実装が追いつきそうにないのでパス。
			// authUnaryInterceptor(),
		),
	)
	service := fileService.NewFileService()

	fileInterface.RegisterFileServiceServer(server, service)

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}

// curl -XGET -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFp.q3MvA6NaAa8m9Yig6EcEe-Dy70ltqY2d8ywhcsu4Coo'  http://localhost:5555/v1/gdrive/list
func authUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "Invalid Request")
		}
		connection, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
		if err != nil {
			return nil, status.Error(codes.Internal, "Connection to auth server denied")
		}
		defer connection.Close()
		context, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		token, hasToken := md["token"]
		if !hasToken {
			client := auth.NewAuthenticationServiceClient(connection)
			uname, userOk := md["user"]
			pass, passOk := md["pass"]
			if !(userOk && passOk) {
				return nil, status.Error(codes.Unauthenticated, "Invalid Request<invalid user or invalid pass>")
			}
			if len(uname) != 1 || len(pass) != 1 {
				return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("Invalid Request <uname and pass is too much or less. (user: %d, pass: %d)", len(uname), len(pass)))
			}
			response, err := client.Authenticate(context, &auth.AuthenticationRequest{Username: uname[0], Password: pass[0]})
			if err != nil {
				return nil, status.Error(codes.Internal, fmt.Sprintf("unknown error. user: %s, pass: %s, err: %s", uname[0], pass[0], err.Error()))
			}
			if !response.GetAuthorized() {
				return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("Failed Authentication. user: %s, pass: %s", uname[0], pass[0]))
			}
			newToken := response.GetToken()
			return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("Please Use token. %s", newToken))
		}
		client := auth.NewAuthorizeServiceClient(connection)
		if len(token) != 1 {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Invalid token. token %v", token))
		}
		response, err := client.Authorize(context, &auth.AuthorizeRequest{Token: token[0], ExpiredAt: nil})
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("unknown error. token: %s, err: %s", token[0], err.Error()))
		}
		if !response.GetAuthorized() {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Invalid token. token %v", token))
		}
		fmt.Printf("%+v", response)

		m, err := handler(ctx, req)
		if err != nil {
		}
		return m, err
	}
}
