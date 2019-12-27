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

3. herokuにgoogle apiのconfig fileを置けない。
google apiのためにidやtoken情報を含んだ証明書ファイルが必要だが、herokuにgit経由以外でファイルをアップすることができないため、
今回のシステム上、セキュリティ面で証明書ファイルのgit経由でのアップはできない。
- 対策として、この証明書ファイルを環境変数に設定して認証をパスする。以下のように環境変数を設定し、暗黙的にgoogleのライブラリが自動で読み取ってくれるようにする。
`heroku config:set GOOGLE_CLOUD_KEYFILE_JSON="$(< keyfile.json)"`

4. Goのtestパッケージを使ったテストでのハマりどころ。
- 異常系のテストで、エラー発生時になぜかテストが終了してしまった。 
    - 原因: logパッケージの`log.Fatal()`をよく読まずに使っていた。testを行ったメソッドでlog.Fatalを使用していて、実はlog.Fatalを実行すると`os.Exit(1)`されてしまう。

5. GoogleStorageの仕様
- bucketの名前空間はプロジェクトでグローバルなのではなく、google storage全体でグローバルである。
直感的でない事象として、同じ名前のバケットを作ってしまうとapiのエラーログにただ単にconflictと出てくる。
自分のプロジェクト内で同じ名前のバケットがなかったので、何かをキャッシュしていたのかと思ってしまった。
また、バケットの削除も権限がないとエラーが出てくるが、それは他のプロジェクトを消しに行っている。

6. circle ci のテストでexit codeが2になってしまい、テストをパスできない。
- 原因: `log.Printf`を呼ぶと`exit code`が2になってしまう。

7. テスト環境と開発環境のdockerイメージが別にしてしまっていたのでとっとと統一すべきだった。また、各環境で環境変数の設定をうまいこと取れるような仕組みが欲しい(Go moduleのオンオフとかherokuのトークンとか。。)。

8. 認証の設計、実装案(要検討)
grpcではインターセプターと呼ばれるミドルウェアがあり、通信の前後もしくは最中に干渉できる機能を実装できる仕組みがある。
今回の認証機能の実装ではミドルウェアとして、プロキシサーバで認証情報をコンテキストにキーバリューで仕込み、サーバ側で取り出して検証する仕組みにする。
