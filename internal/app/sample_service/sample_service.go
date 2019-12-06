package sample_service

import (
	"context"

	echo "github.com/TakeruTakeru/poc-go-micro-service/api/echo_api"
)

type EchoService struct {
	b string
}

func (e *EchoService) EchoMessage(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	message := req.GetMessage()
	if message == "" {
		message = "Say something"
	}
	return &echo.EchoResponse{Message: message}, nil
}

func NewEchoService() *EchoService {
	return &EchoService{}
}
