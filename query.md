# GraphQLのクエリ例
スキーマについては[こちら](./graph/schema/)を見てください
## ユーザー登録
- リクエスト
```
mutation 任意の名前($i: Upload!){
  registerUser(input:{name:"テストユーザー",password:"test1234",description:"これはテストユーザーです",profieImage:$i}){
    userID
    name
    password
    profieImageURL
  }
}
```
- レスポンス
```
{
  "data": {
    "registerUser": {
      "userID": 1,
      "name": "テストユーザー",
      "password": "$2a$10$O8yDtapGPXflkyl.8./Obehe90bxQ4Zc6OnfZQKgng3XLMnyzazWa",
      "profieImageURL": "https://バケット名.s3.リージョン.amazonaws.com/生成されたuuid.ファイルの拡張子"
    }
  }
}
```

## ログイン
- リクエスト
```
mutation 任意の名前{
  auth{
    login(input:{name:"テストユーザー",password:"test1234"}){
      success
      token
    }
  }
}
```
- レスポンス
```
{
  "data": {
    "auth": {
      "login": {
        "success": true,
        "token": "JWTのトークン"
      }
    }
  }
}
```
このtokenにあるJWTのトークンを使います
## ユーザー情報を更新
`注意 このリクエストにはJWTトークンをHeadersに追加してください`
```
{
  "Authorization":"Bearer JWTトークン"
}
```
名前とプロフィール画像を更新する場合
リクエスト
```
mutation 任意の名前($i: Upload!){
  updateUser(input:{name:"変更したテストユーザー",profieImage:$i}){
    name
    profieImageURL
    password
  }
}
```
レスポンス
```
{
  "data": {
    "updateUser": {
      "name": "変更したテストユーザー",
      "profieImageURL": "https://バケット名.s3.リージョン.amazonaws.com/変更された画像の生成されたuuid.ファイルの拡張子",
      "password": "$2a$10$O8yDtapGPXflkyl.8./Obehe90bxQ4Zc6OnfZQKgng3XLMnyzazWa"
    }
  }
}
```
`注意 このmutationではパスワードは変更できません`
## ユーザーの削除
`注意 このリクエストにはJWTトークンをHeadersに追加してください`
```
{
  "Authorization":"Bearer JWTトークン"
}
```
リクエスト
```
mutation 任意の名前{
  deleteUser{
    success
  }
}
```
レスポンス
```
{
  "data": {
    "deleteUser": {
      "success": true
    }
  }
}
```
ユーザーを削除をするとユーザーの持っているアニメランキングも削除します

## 全てのユーザー情報を取得
リクエスト
```
query 任意の名前{
  GetAllUserInformation{
    name
    password
    userID
    haveAnime{
      rank
      title
      animeID
      description
      animeImageURL
    }
  }
}
```
レスポンス
```
{
  "data": {
    "GetAllUserInformation": [
      {
        "name": "テストユーザー",
        "password": "$2a$10$O8yDtapGPXflkyl.8./Obehe90bxQ4Zc6OnfZQKgng3XLMnyzazWa",
        "userID": 1,
        "haveAnime": [
          {
            "rank": 1,
            "title": "テストアニメ",
            "animeID": 1,
            "description": "テストアニメの説明",
            "animeImageURL": "https://バケット名.s3.リージョン.amazonaws.com/画像の生成されたuuid.ファイルの拡張子""
          },
          {
            "rank": 22
            "title": "テストアニメ2",
            "animeID": 2,
            "description": "",
            "animeImageURL": "S3のURL"
          }
        ]
      },
      {
        "name": "ユーザー2",
        "password": "$2a$10$UqHOT.2PF1EXEkc8Bbxbo.dlXgxL/.fEdDnxey0TtSCYc/lZS8HMW",
        "userID": 2,
        "haveAnime": [
          {
            "rank": 1,
            "title": "面白いやつ",
            "animeID": 3,
            "description": "ホゲーー",
            "animeImageURL": "S3のURL"
          }
        ]
      }
    ]
  }
}
```

## ユーザーの情をIDから取得
リクエスト
```
query 任意の名前{
  user(id:1){
    name
    password
    userID
    haveAnime{
      rank
      title
      animeID
      description
      animeImageURL
    }
  }
}
```
レスポンス
```
{
  "data": {
    "user": {
      "name": "テストユーザー",
      "password": "$2a$10$O8yDtapGPXflkyl.8./Obehe90bxQ4Zc6OnfZQKgng3XLMnyzazWa",
      "userID": 1,
      "haveAnime": [
        {
          "rank": 1,
          "title": "テストアニメ",
          "animeID": 1,
          "description": "hogege",
          "animeImageURL": "https://バケット名.s3.リージョン.amazonaws.com/変更された画像の生成されたuuid.ファイルの拡張子"
        }
      ]
    }
  }
}
```

## アニメランキングを作成
`注意 このリクエストにはJWTトークンをHeadersに追加してください`
```
{
  "Authorization":"Bearer JWTトークン"
}
```

リクエスト
```
mutation 任意の名前($i: Upload!){
    createAnimeRanking(input:{title:"テストアニメ",description:"テストアニメ説明",rank:1,animeImage:$i}){
    animeID
    title
    description
    animeImageURL
    rank
  }
}
```
レスポンス
```
{
  "data": {
    "createAnimeRanking": {
      "animeID": 1,
      "title": "テストアニメ",
      "description": "hogege",
      "animeImageURL": "https://バケット名.s3.リージョン.amazonaws.com/生成されたuuid.ファイルの拡張子",
      "rank": 1
    }
  }
}
```

## アニメランキングをanimeIDにより更新
`注意 このリクエストにはJWTトークンをHeadersに追加してください`
```
{
  "Authorization":"Bearer JWTトークン"
}
```
タイトルと画像を更新する場合
リクエスト
```
mutation 任意の名前($i: Upload!){
  updateAnimeRanking(id:1,input:{title:"変更後のテストアニメ",animeImage:$i}){
    animeID
    title
    description
    animeImageURL
    rank
  }
}
```
レスポンス
```
{
  "data": {
    "updateAnimeRanking": {
      "animeID": 1,
      "title": "変更後のテストアニメ",
      "description": "テストアニメ説明",
      "animeImageURL": "https://バケット名.s3.リージョン.amazonaws.com/変更された画像の生成されたuuid.ファイルの拡張子",
      "rank": 1
    }
  }
}
```

## アニメランキングをanimeIDにより削除
`注意 このリクエストにはJWTトークンをHeadersに追加してください`
```
{
  "Authorization":"Bearer JWTトークン"
}
```

リクエスト
```
mutation 任意の名前{
  deleteAnimeRanking(animeID:1){
    success
  }
}
```
レスポンス
```
{
  "data": {
    "deleteAnimeRanking": {
      "success": true
    }
  }
}
```
## animeIDによりアニメランキングを取得
リクエスト
```
query 任意の名前{
  getAnimeRanking(id:1){
    rank
    title
    animeID
    description
    animeImageURL
  }
}
```
レスポンス
```
{
  "data": {
    "getAnimeRanking": {
      "rank": 1,
      "title": "テストアニメ説明",
      "animeID": 1,
      "description": "テストアニメ説明",
      "animeImageURL": "https://バケット名.s3.リージョン.amazonaws.com/生成されたuuid.ファイルの拡張子"
    }
  }
}
```