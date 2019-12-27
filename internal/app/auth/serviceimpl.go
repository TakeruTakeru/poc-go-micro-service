package auth

import (
	"context"

	"github.com/TakeruTakeru/poc-go-micro-service/api/auth"
)

type AuthService struct{}

func (as *AuthService) Authenticate(ctx context.Context, req *auth.AuthenticationRequest) (res *auth.AuthenticationResponse, err error) {
	uname := req.GetUsername()
	pass := req.GetPassword()

}

func (as *AuthService) ResetToke(ctx context.Context, req *auth.AuthenticationRequest) (res *auth.AuthenticationResponse, err error) {

}

func Varify(uname string, pass string)

// curl -XGET -H 'Authorization:takeru-pass' -H 'Content-Type:application/json' -H 'User-Agent:iPhone' -H 'Accept-Encoding:gzip,deflate' -d "{"key""val","key2":",val2"}" http://localhost:5555/v1/gdrive/list
