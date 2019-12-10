## ハマったこと
1. go getしてもsrcにソースが落ちてこない
`GO111MODULE=off`をしないと`go modules`を利用して開発を行っている場合、`$GOPATH/src`配下へのダウンロードは省かれるので、
`go get`しても`src`には何もないという状況になる。
`golang`に入門した時に`go modules`を採用していないと、今までsrcにも落とされていて、`go modules`を利用して落としたのとそうでないので、ごちゃ混ぜになっているので、ここら辺の挙動に気付きづらい。

2. protocでコンパイルが上手くいかない。
以下のようにコマンドが通らなかった。
```
protoc -I/usr/local/bin -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis=plugins=grpc:. ./api/echo_api/sample.proto
google/protobuf/descriptor.proto: File not found.
google/api/annotations.proto:20:1: Import "google/protobuf/descriptor.proto" was not found or had errors.
google/api/annotations.proto:28:8: "google.protobuf.MethodOptions" is not defined.
api/echo_api/sample.proto:4:1: Import "google/api/annotations.proto" was not found or had errors.
```
問題は`-I/usr/local/include`にprotoのファイルがある想定でこのコマンドは成り立っていたが、`-I/usr/local/include`に.protoが配置されていなかったことが原因だった。このオプションは、指定したパスをprotoファイルの対象検索ディレクトリとして登録するもの。protobufのディレクトリ群をインストール時に`/usr/local/include`へ配置しなかったため起きた。