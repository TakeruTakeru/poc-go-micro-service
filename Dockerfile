FROM golang:latest

ENV GOARCH="amd64"
ENV GOOS="linux"
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /go/src/github.com/TakeruTakeru
COPY . .

RUN go build -o $GOPATH/bin/main-server ./cmd/web/fileserver/grpc-server
RUN go build -o $GOPATH/bin/proxy-server ./cmd/web/fileserver/restserver