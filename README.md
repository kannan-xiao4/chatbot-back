# Chatbot using Prediction API

# Overview（以下は想定）
- APIリクエストにメッセージを送ることで、PredictionAPIで予測された内容をレスポンスとして返す
- ユーザー登録、返答メッセージの保存機能

## Description
- TBD

## Demo
- TBD

## Requirement
- Go - https://github.com/golang/go
- Dep - https://github.com/golang/dep
- Gin - https://github.com/gin-gonic/gin

## Usage
- 以下のコマンドを実行
`go run main.go`

- `http://localhost:3000/welcome` で起動確認

## DB initialize
- provision : 現存のDBを破棄して作り直す
```
go run command/provision.go
```
- migration : すでにDBがある状態でテーブルを作成する
```
go run command/migrate.go
```

## Install
- ディレクトリ構成は下記で
```$xslt
└── chatbot
    └── src
        └── chatbot-back
            └── main.go
```
- GoLandの設定
```$xslt
GOROOT: インストール済みのものをプルダウンから選択
GOPATH: 上記のディレクりのchatbot を選択
dep: depを有効にし、プルダウンから選択
```
- depで依存解決
    - `dep ensure`が実行できるので行う

