# grpc4patterns

- gRPCの概要は割愛。
- gRPCの4種類の通信方式をお手軽に試せるコードです

# gRPCの通信方式は4種類

## Unary RPC
- シンプルな1 Request - 1 Response方式
- 用途: 一般的なRESTと同じケース

## Server Streaming RPC
- 1 Request - N Response方式
- 用途: サーバサイドプッシュ・フィードなど
- 用途：サーバから任意のタイミングでクライアントに通知させたいケース

## Client Streaming RPC
- N Request - 1 Response方式
- 用途：データアップロードや、クライアントから多くのデータを送るケース

## Bidirectional Streaming RPC
- 単一TPCコネクションの中で、ReqestとResponseの送受信を任意数繰り返す
- 用途: チャットシステムなど（従来のWebSocketみたいな使い方)
- 用途: オンライン対戦ゲームなど

## ソース構成

```
grpc4patterns
|
├── bidirectionalstreaming   # チャットみたいな双方向通信のサンプル
│   ├── client
│   │   └── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── pb
│   │   └── chat
│   │       └── chat.pb.go
│   ├── proto
│   │   └── chat.proto
│   └── server
│       └── main.go
├── clientstreaming　# データアップロードや、クライアントから多くのデータを送るサンプル
│   ├── client
│   │   └── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── pb
│   │   └── upload
│   │       └── upload.pb.go
│   ├── proto
│   │   └── upload.proto
│   └── server
│       └── main.go
├── srvstreaming   #プッシュ通知のようなサーバから任意のタイミングでクライアントに通知させたいときのサンプル
│   ├── client
│   │   └── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── pb
│   │   └── notification
│   │       └── notification.pb.go
│   ├── proto
│   │   └── notification.proto
│   └── server
│       └── main.go
└── unary  # 一般的なRESTと同じようなサンプル
    ├── client
    │   └── main.go
    ├── go.mod
    ├── go.sum
    ├── pb
    │   └── calc
    │       └── calc.pb.go
    ├── proto
    │   └── calc.proto
    └── server
        └── main.go
```
---

# 環境構築

CentOS7 を開発環境と想定して必要なライブラリの導入手順を書いておきます。
Server/ClientともにGoで書く前提です。

    CentOS Linux release 7.6.1810 (Core)

##  Go gRPCライブラリ

    go get -u google.golang.org/grpc


##  protocol buffer コンパイラ

- protocol buffer rpmを取得してインストールします。
- protoファイルからコード生成をするコンパイラ(protoc)

```
wget http://mirror.centos.org/centos/7/os/x86_64/Packages/emacs-filesystem-24.3-22.el7.noarch.rpm
wget https://cbs.centos.org/kojifiles/packages/protobuf/3.6.1/4.el7/x86_64/protobuf-3.6.1-4.el7.x86_64.rpm
wget https://cbs.centos.org/kojifiles/packages/protobuf/3.6.1/4.el7/x86_64/protobuf-compiler-3.6.1-4.el7.x86_64.rpm

sudo rpm -ivh *.rpm

```

## Protocol Buffersのプラグインをインストール

    go get github.com/golang/protobuf/protoc-gen-go


## gRPCミドルウェア

    go get -u github.com/grpc-ecosystem/go-grpc-middleware
    go get -u github.com/sirupsen/logrus



## 参考URL
https://grpc.io/docs/guides/concepts/

https://christina04.hatenablog.com/entry/2017/11/13/190000

https://qiita.com/tomo0/items/310d8ffe82749719e029#server-streaming-rpc-1
