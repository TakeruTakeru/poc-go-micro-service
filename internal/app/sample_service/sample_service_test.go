package sample_service

import (
	"context"
	"testing"

	echo "github.com/TakeruTakeru/poc-go-micro-service/api/echo_api"
)

func TestEchoService(t *testing.T) {
	service := NewEchoService()
	testCases := []struct {
		req  *echo.EchoRequest
		want string
	}{
		{&echo.EchoRequest{Message: "hoge"}, "hoge"},
		{&echo.EchoRequest{Message: ""}, "Say something"},
	}

	for _, testCase := range testCases {
		res, err := service.EchoMessage(context.Background(), testCase.req)
		if err != nil {
			t.Errorf("Failed: Runtime error.")
		}
		if res.GetMessage() != testCase.want {
			t.Errorf("Failed: Expected %s, but Got %s", testCase.want, res.GetMessage())
		}
	}
}
