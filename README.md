# react_golang_mux

mux を利用して react と golang の通信、cors 設定
docker を利用しての環境構築

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

クリーンアーキテクチャを採用

[参考サイト 2](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)

## ローカルでの起動方法

クローン

```
git clone git@github.com:kory-jp/react_golang_mux.git
```

ルートディレクトリへ移動

```
cd golang_react_mux
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

golang のログを確認

```
docker logs go_container
```
