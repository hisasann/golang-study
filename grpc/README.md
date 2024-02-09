
## Protocol buffer compilerをインストールする

[Protocol Buffer Compiler Installation | gRPC](https://grpc.io/docs/protoc-installation/)

```bash
brew install protobuf
protoc --version  # libprotoc 25.2
```

## 必要なモジュールのインストール

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
cd grpc
go mod init github.com/hello-grpc
```

## protoファイルからprotoc-gen-go/protoc-gen-go-grpcファイル生成

```bash
cd proto
protoc  --go_out=. --go-grpc_out=require_unimplemented_servers=false:. hello-grpc.proto
```

## serverとclientを作成する

server と client ディレクトリを参照。

## パッケージをダウンロードする

```bash
go mod tidy
```

## サーバーを起動する

```bash
go run server/main.go
```

## クライアントを起動する

```bash
go run client/main.go
```

## 参考記事

[Go言語gRPCを爆速で試す方法 #Go - Qiita](https://qiita.com/totoaoao/items/6bf533b6d2164b74ac09)

[Quick start | Go | gRPC](https://grpc.io/docs/languages/go/quickstart/)

[はじめに｜作ってわかる！ はじめてのgRPC](https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/intro)