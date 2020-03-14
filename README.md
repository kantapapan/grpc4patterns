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

## 参考URL
https://grpc.io/docs/guides/concepts/

https://christina04.hatenablog.com/entry/2017/11/13/190000

https://qiita.com/tomo0/items/310d8ffe82749719e029#server-streaming-rpc-1
