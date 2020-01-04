package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	gw "github.com/TakeruTakeru/poc-go-micro-service/api/fileservice"
)

var (
	port     = flag.String("port", "8080", "listen port")
	endpoint = flag.String("end_point", "localhost:5555", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithForwardResponseOption(CORSFilter))
	err := gw.RegisterFileServiceHandlerFromEndpoint(ctx, mux, *endpoint, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}
	fmt.Printf("Server started PORT: %s. Proxy to %s.", *port, *endpoint)
	return http.ListenAndServe(":"+*port, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func ClientIAuthInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) (err error) {
		md, _ := metadata.FromOutgoingContext(ctx)
		basicAuth, ok := md["grpcgateway-authorization"]
		if ok {
			md := metadata.Pairs("token", basicAuth[0])
			ctx = metadata.NewOutgoingContext(ctx, md)
			err = invoker(ctx, method, req, reply, cc, opts...)
		}
		return
	}
}

func CORSFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return nil
}
