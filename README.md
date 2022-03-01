# react_golang_mux

## 概要

フレームワークを利用しない golang,Typescript を利用した Redux 開発の練習のための TODO 開発

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

```
SESSION_KEY = ********
```

docker 起動

```
docker-compose build
docker-compose up -d
```

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
