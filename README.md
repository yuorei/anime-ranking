# my-anime-ranking
実際のアプリは[こちら](https://my-anime-ranking-front.vercel.app/)  ＊現在はバックエンドが動いておりません。
フロントエンドのリポジトリは[こちら](https://github.com/yuorei/my-anime-ranking-front)  
## 概要
- こちらはマイアニメランキングアプリのバックエンドです。  
- 自分の好きなアニメのランキングを共有できるサービスです。  
- ユーザー登録をすることでランキングにアニメを登録することができます。  
- このサービスでいろんなユーザーのアニメのランキングを見ることができます。  
- ユーザー情報には名前、説明、プロフィール画像を含められています。  
- 名前とパスワードでログインをすることができます。  
- ログインすることでアニメのランキングを登録、更新、削除ができるようになります。  
- アニメの情報にはタイトル、ランク、説明、画像を含められています。  
- 現在はバックエンドのみの実装となっておりますので任意のPOSTを送信できるツールでお試しください。  
## 技術スタック
### 言語
- Go
### APIアーキテクチャ
- GraphQL API
### フレームワーク
- gqlgen  
GraphQL APIを扱うためのフレームワークです。
### データベース
- MySQL  
ORMとしてgormを使用しています。
### 画像アップロード先
- Cloud Strage
## 開発していた環境
- Apple M1 Mac OS 13.2.1
## 使い方
ローカルで立ち上げる場合
```
git clone git@github.com:yuorei/anime-ranking.git
```
クローンしてきてください
go をインストールしてください
```
go version
```
これでバージョンが出れば大丈夫です。

モジュールをインストールできていない場合は
```
go get hogehoge
```
でインストールしてください

webp ツールもインストールしてください。  
https://github.com/kolesa-team/go-webp

MySQLのインストールをしてください
MySQLのサーバーを起動して
データベースを作ってください

`.env`をmain.goと同じディレクトリにおいてください
中身については以下を参考にしてください
```
DB_USER=データベースのユーザー名
DB_PASS=データベースのパスワード
DB_NAME=データベース名
DB_PORT=127.0.0.1:3306 // これはローカル
DB_TZ=True&loc=Local

JWT_SECRET=何らかの文字列
S3_BUCKET_NAME=S3のバケット名
S3_REGION=S3のリージョン
```
すべて完了した場合は
```
go run main.go
```

POSTで`/query`にGraphQLのクエリを送信してください

送信できるクエリについては[こちら](./query.md)を確認ください。

## 工夫した点について
- GraphQLクエリでUserにネストされているhaveAnimeを要求しない時に
不要なDB処理を行わないためにhaveAnimeに`@goField(forceResolver: true)}`をつけることにより機能を使って分割しました。

- resolverの中の処理が複雑な場合はapplicationパッケージにて行うようにしました。
DB処理はmysqlパッケージにて行いました。
これによりresolverはなにかしらの処理をされたものを返すものとして簡潔にみることができます。  

IDのNodeを実装したことにより事でグローバルなIDを返すようにしました。  
`user:1`,`anime:1`
