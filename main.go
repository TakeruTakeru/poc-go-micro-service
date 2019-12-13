package main

import (
	"fmt"
	"os/exec"
)

func CreateInterfaceDoc() {
	_, err := exec.Command(
		"protoc",
		"-I/usr/local/include",
		"-I.",
		"-I$GOPATH/src",
		"-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis",
		"--doc_out=html,index.html:./api/fileservice",
		"./api/fileservice/*.proto",
	).Output()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func main() {
	CreateInterfaceDoc()
}
