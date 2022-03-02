# react_golang_mux

## 概要/開発経緯

Golang の基本的な知識を習得するために開発している Todo アプリになります。
基礎的なところから理解を図るために極力、フレームワークを利用せずに開発に努めております。
また合わせて、React+TypeScript+Redux の組み合わせでアプリを作成することも本のアプリの目標としており、こちらも最低限のパッケージで開発を進めております。

## 環境構築

[参考サイト 1](https://qiita.com/takuya911/items/2447c97525d4c48b72a2)

- ディレクトリ構成

```
./-- api //golang
  |_client //react-tyepscript
  |_nginx/nginx.conf
  |_mysql/Dockerfile
      |_conf.d/my.conf
```

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

その他

- docker(開発環境構築)
- .air(Golang のホットリロード)
- deleve(Golang におけるデバッグ)

## アーキテクチャ設計

api(golang)/Clean Architecture を採用

[参考サイト 2](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)

client(react)/Atomic Design を採用

[参考サイト 3](https://www.happylifecreators.com/blog/20220113/)

## ローカルでの起動方法

クローン

```
git clone git@github.com:kory-jp/react_golang_mux.git
```

ルートディレクトリへ移動

```
cd golang_react_mux
```

環境変数設定

```
touch api/env/dev.env
```

api/env/dev.env へキー(環境変数)を入力

```
SESSION_KEY = ********
```

docker 起動

```
docker-compose build
docker-compose up -d
```

サイトアクセス

[Todo](http://localhost:8080/)

コンテナログイン

```
docker exec -it go_container /bin/sh
docker exec -it react_container /bin/sh
docker exec -it mysql_container /bin/sh
```

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

## 各種ログ確認

golang のログを確認

```
docker logs go_container
```

MySQL のクエリを確認

`/mysql/lgos/mysqld.log`

## 主要機能

User

- 新規登録
- ログイン
- ログアウト

Todo

- 新規投稿
- 一覧表示
- 詳細表示
- 編集
- 削除
- 完了/未完了切り替え

その他

画像投稿/保存/配信

Toast による通知機能

Todo が 5 以上投稿されるとページネーションにより、データの一部取得

## モデルデザイン

User

```
create table if not exists users (
	id integer primary key auto_increment,
	name varchar(50) NOT NULL,
	email varchar(50) NOT NULL UNIQUE,
	password varchar(50) NOT NULL,
	created_at datetime default current_timestamp
);
```

Todo

```
create table if not exists todos (
	id integer primary key auto_increment,
	user_id integer NOT NULL,
	title varchar(50) NOT NULL,
	content text NOT NULL,
	image_path varchar(100),
	isFinished boolean NOT NULL,
	created_at datetime default current_timestamp
);
```

## 基本操作

`/`からテストユーザーのメールアドレスとパスワードでログイン。
または`/registration`からユーザー名、メールアドレス、パスワードを入力してユーザー登録とログインが可能

ログインすると過去に投稿した Todo 一覧(`todo`)が表示される
ヘッダー部分には新規 Todo 投稿画面(`todo/new`)へのリンクとログアウトボタンを設置

一覧画面から各 Todo の詳細情報のページに遷移することや、完了未完了の操作、削除の操作が可能

詳細画面(`todo/show/:id`)からはより詳細の Todo 情報を確認でき編集や削除の操作も可能
