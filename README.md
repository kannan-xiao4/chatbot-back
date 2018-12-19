# Chatbot using SmartTalk API

# Overview（以下は想定）
- APIリクエストにメッセージを送ることで、SmartTalkAPIで予測された内容をレスポンスとして返す
- ユーザー登録、返答メッセージの保存機能


## Requirement
- Go - https://github.com/golang/go
- Dep - https://github.com/golang/dep
- Gin - https://github.com/gin-gonic/gin

## Prepare
### Firebase Admin SDK
- Firebaseからadminsdkのjsonを取得する
- 以下の導入を参照、そのまま
    - https://firebase.google.com/docs/admin/setup?hl=ja
- 配置場所は envfiles以下に置く

### envの作成
- 以下の形式のenvファイルをenvfiles 以下に置く
```
FIREBASE_ADMIN_SDK_FILENAME: envfiles/adminsdk_filename.json
RECRUIT_SMARTTALK_API_KEY: AAAAAAAAAAAAAAAAAAAAAAAAAAAAA
RECRUIT_SMARTTALK_ENDPOINT: https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk
DATABASE_USER: root
DATABASE_PASSWARD: secret
DATABASE_HOST_NAME: 127.0.0.1:3306
FIREBASE_DATABASE_URL: https://projectname.firebaseio.com/

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

## DB initialize
- provision : 現存のDBを破棄して作り直す
```
go run command/provision.go
```
- migration : すでにDBがある状態でテーブルを作成する
```
go run command/migrate.go
```

## Usage
- 以下のコマンドを実行
`go run main.go`

- `http://localhost:3000/welcome` で起動確認
