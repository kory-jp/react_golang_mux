# react_golang_mux

## 概要/開発経緯

`Golang` の基本的な知識を習得するために開発している Todo アプリになります。
基礎的なところから理解を図るために極力、フレームワークを利用せずに開発に努めております。
また合わせて、`React+TypeScript+Redux` の組み合わせでアプリを作成することも目標としており、こちらも最低限のパッケージで開発を進めております。

<br>
<br>

## こだわったところ

- API(golang)を`Clean Architecture`にて設計

以前に `Ruby on Rails` にて作成したポートフォリオは `Rails`の設計である `MVC`に準じてアプリケーションを作成しました。

[Rails のポートフォリオ](https://github.com/kory-jp/proto2)

今回は別の視点からアプリ開発の知識を得ることを目標に設計に`Clean Architecture`を採用しました。保守性などの観点から`MVC`以上に各担当役割を細かく分けている設計思想なので、コーディングの際は常にファイルの役割を考えつつ多少コードの記載量が増えても役割ごとにファイル分割を意識しました。

</br>

- データ整合性を維持するデータベース処理`Transaction`を、上位レイアのアプリケーション層で実行できるように`interface`を用いたダックタイピングにて実現

`Todo`をタグ分類するために、モデルとして`Todo`,`Tag`中間テーブルとして`Todo_Tag_Relations`を作成。`Todo`と`Todo_Tag_Relations`を保存する際に、どちらか一方が保存処理に失敗した場合、`Rollback`をするように設計しました。ポイントとして、`transaction`にまつわるデータベース操作は`Clean Architecture`において`interfaces/database`層が制御することになりますが、アプリにおけるデータの保存や廃棄の判断はより上位レイアの`usecase/interactor`(アプリケーション)層が関心をもつ必要があります。一方、上位レイア(アプリケーション層)は下位レイア(`interfaces/database`)に依存することは`Clean Architecture`においては設計思想に沿いません。そのため、DIP（依存関係逆転の原則）の考えのもと`interface`を用いて、いわゆるダックタイピングによりアプリケーション層にて`Transaction`操作を実現しました。

</br>

- API(golang)側の`Mock`を利用したテストコード

上述の通り、API 側の設計として、`Clean Architecture`を採用しており各種役割ごとに階層化されており、効率的にテストするために初めて`goMock`を採用しました。
とくにアプリケーションロジックの役割をになっている `api/usecase/interactor` をテストする際に、通信やデータベース処理おいて発生したエラーなどアプリケーションロジックに関係ないエラーを排除できるような設計にできました。

</br>
</br>

## 環境構築

下記のサイトをもとに、`golang`の開発効率の向上を図るため`.air`によるホットリロードを実現

`MySQL`において発行された`SQL`を `log` ディレクトリ以下に記録

[参考サイト 1](https://qiita.com/takuya911/items/2447c97525d4c48b72a2)

- ディレクトリ構成

```
./-- api //golang
  |_client //react-tyepscript
  |_nginx/nginx.conf
  |_mysql/Dockerfile
      |_conf.d/my.conf
```

<br>
<br>

## 使用技術

バックエンド

- Golang(Go Modules)
- Nginx
- MySQL

フロントエンド

- React(create-react-app/Material-ui/Redux)
- TypeScript
- HTML/CSS

Golang 主要パッケージ

- go-playground/validator
- gorilla/mux
- gorilla/sessions
- joho/godotenv
- rs/cors
- ini.v1
- stretchr/testify

その他

- docker(開発環境構築)
- .air(Golang のホットリロード)
- deleve(Golang におけるデバッグ)

<br>
<br>

## 技術選定理由

### バックエンド: Golang

- 前回、`Ruby on Rails`でアプリを開発しており、次は静的型付け言語でフレームワーク無しの開発を考えていた
  [Rails のポートフォリオ](https://github.com/kory-jp/proto2)
- コードがシンプルで静的型付け言語においては学習難易度が低い
- アプリを作成するためのライブラリが豊富
- 近年の言語としての勢い

### フロントエンド: React/TypeScript

- コンポーネント分割に重きをおくライブラリで、ある程度の分割が出来ていれば可読性が上がるのと同時にメンテナンス性が上がるため
- `TypeScript`を併用することで型によるさらなる安全性の向上
- `Vue` との比較で、より深く `JavaScript` の理解が必要な `React` にキャリア初期から触れることで自己成長につながるのではないかと考えたため

<br>
<br>

## アーキテクチャ設計

_api(golang)/Clean Architecture を採用_

![](https://blog.tai2.net/images/CleanArchitecture.jpg)

以下の理由で採用

- 再利用性の高い設計になり生産性が向上する
- コードの可読性が上がり、メンテナンスが容易になる
- 変化に強い設計になる

[参考サイト 2](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)

<br>

_client(react)/Atomic Design を採用_

![](https://cdn.codegrid.net/2017-atomic-design/img/hierarchy.png)

以下の理由で採用

- コードの再利用性が高いの効率的に開発が可能
- メンテナンスが容易になる

[参考サイト 3](https://www.codegrid.net/articles/2017-atomic-design-1/)

<br>
<br>

## 主要機能

User

- 新規登録
- ログイン
- ログアウト

\*session と cookie を利用することでにログイン状態の確認

Todo

- 新規投稿
- 一覧表示
- 詳細表示
- 編集
- 削除
- 完了/未完了切り替え
- タグ分類
- タグ検索
- 重要度で分類
- 緊急度で分類
- 重要度、緊急度で検索

その他

画像投稿/保存/配信

Toast による通知機能

Todo が 5 以上投稿されるとページネーションにより、データの一部取得

コマンドを実行することでテストデータを自動作成

一部、ファイルに対してテストコードを作成

<br>
<br>

## 基本操作

`/`からテストユーザーのメールアドレスとパスワードでログイン。
または`/registration`からユーザー名、メールアドレス、パスワードを入力してユーザー登録とログインが可能

ログインすると過去に投稿した Todo 一覧(`todo`)が表示される
ヘッダー部分には新規 Todo 投稿画面(`todo/new`)へのリンクとログアウトボタンを設置

一覧画面から各 Todo の詳細情報のページに遷移することや、完了未完了の操作、削除の操作が可能

詳細画面(`todo/show/:id`)からはより詳細の Todo 情報を確認でき編集や削除の操作も可能

<br>
<br>

## ローカルでの起動方法

- クローン

```
git clone git@github.com:kory-jp/react_golang_mux.git
```

- ルートディレクトリへ移動

```
cd golang_react_mux
```

- 環境変数設定

React 側

```
touch client/client-app/.env
```

client/client-app/.env へキー(環境変数)を入力

```
REACT_APP_API_URL="http://localhost:8000/api/"
```

golang 側

```
touch api/env/dev.env
```

api/env/dev.env へキー(環境変数)を入力

```
SESSION_KEY = ********
```

- docker 起動

```
docker-compose build
docker-compose up -d
```

- サイトアクセス

[Todo](http://localhost:8080/)

- コンテナログイン

```
docker exec -it go_container /bin/sh
docker exec -it react_container /bin/sh
docker exec -it mysql_container /bin/sh
```

<br>
<br>

## テストデータ投入

```
docker exec -it go_container /bin/sh
go run infrastructure/seeds/development/seeder.go
```

テストユーザーによるログイン

```
メールアドレス:
sam@exm.com
パスワード
password
```

<br>
<br>

## 各種ログ確認

golang のログを確認

```
docker logs go_container
```

MySQL のクエリを確認

`/mysql/lgos/mysqld.log`

<br>
<br>

## テストコード実行

- 実装したテストコード

```
api/interfaces/controllers/users/user_controller_test.go
api/interfaces/controllers/sessions/session_controller_test.go
api/interfaces/controllers/todos/todo_controller_test.go
api/usecase/user/user_interactor_test.go
api/usecase/session/session_interactor_test.go
api/usecase/todo/todo_interactor_test.go
```

- テストコード実行

```
cd api
sh test.sh
```

- 成功した場合

```
$sh test.sh
TEST ALL COMPLETED
```

- 失敗した場合

エラーログの詳細が表示

<br>
<br>

## モデルデザイン

![ER](https://user-images.githubusercontent.com/66899822/165004467-2e4eb863-3025-49f5-8bfe-a32f3544ea75.png)

<br>
<br>

## 各種 API 仕様詳細

以下のサイトを参考にデザイン

[［API］ API 仕様書の書き方](https://qiita.com/sunstripe2011/items/9230396febfab2eae2c2)

- [User API](https://github.com/kory-jp/react_golang_mux/tree/main/api/interfaces/controllers/users)
- [Session API](https://github.com/kory-jp/react_golang_mux/tree/main/api/interfaces/controllers/sessions)
- [Todo API](https://github.com/kory-jp/react_golang_mux/tree/main/api/interfaces/controllers/todos)
- [Tag API](https://github.com/kory-jp/react_golang_mux/tree/main/api/interfaces/controllers/tags)

<br>
<br>
