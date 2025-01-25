# Go for Beginner

[Notion](https://www.notion.so/116a1df91a078040a95be13c1dc71249?pvs=4)

1. [A Tour of Go](https://go-tour-jp.appspot.com/methods/1)

2. [A Tour of Go](https://go-tour-jp.appspot.com/flowcontrol/1)

[golang](https://hub.docker.com/_/golang)

## Docker
Docker Image

```dockerfile
FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
```

------

マルチステージビルドは、Dockerで複数のステージを使用してコンテナイメージを構築する方法です。

主な特徴：
- ビルド環境と実行環境を分離
- 最終イメージサイズを最小化
- ビルドツールや依存関係を実行環境に含めない

基本構造：
1. ビルドステージ：ソースコードのコンパイルと依存関係の解決
2. 実行ステージ：必要な成果物のみをコピーして実行環境を構築

Dockerfile例:
```dockerfile
# ビルドステージ
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o main

# 実行ステージ
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
```

利点:
- 最終イメージが小さい
- ビルドツールを含まない
- セキュリティリスク低減

```shell
docker build -t main-app .
```

```shell
docker run -it --rm --name my-running-app main-app
```