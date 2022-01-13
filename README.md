# react_golang_mux

mux を利用して react と golang の通信、cors 設定
docker を利用しての環境構築

## 環境構築

[参考サイト 1](https://qiita.com/Blueman81/items/72ca43681d16d44e21ad)

- ディレクトリ構成

```
./-- api //golang
  |_client //react-tyepscript
  |_nginx/nginx.conf
  |_mysql/Dockerfile
      |_conf.d/my.conf
```

### React

- Dockerfile

```
FROM node:16-alpine
WORKDIR /app/client
RUN npm install --save prop-types
RUN npm install -g create-react-app
```

- docker-compose.yml

```
version: '3'
services:
  client:
    build:
      context: .
      dockerfile: ./react_golang_mux/client/Dockerfile
    container_name: react_container
    tty: true
    volumes:
    - ./client:/app/client
    command: sh -c "cd /app/client/client-app  && yarn start"
    ports:
      - 3000:3000
```

- コマンド

```
docker-compose run --rm client sh -c 'create-react-app client-app --template typescript'
```

- rm: 停止済みのサービス・コンテナを削除します。/コンテナを停止した段階で削除してくれる。
- sh -c: yml ファイルの command を実行？

```
docker-compose build
docker-compose up -d
```

### Golang

```
cd api
go mod init github.com/kory-jp/react_golang_mux/api && touch go.sum
```

- Dockerfile

```
FROM golang:alpine
COPY ./api /app/api
WORKDIR /app/api
RUN go mod init github.com/kory-jp/react_golang_mux/api && go build -o main .
CMD ["/app/api/main"]
```

- docker-compose.yml

```
  api:
    build:
      context: .
      dockerfile: ./api/Dockerfile
    container_name: go_container
    ports:
      - 8000:8000
    tty: true
```

- ファイル作成

api/main.go 作成

- コマンド&エラー

```
docker-compose build

 => ERROR [4/4] RUN go build -o main .                                                            0.4s
------
 > [4/4] RUN go build -o main .:
#9 0.336 go: go.mod file not found in current directory or any parent directory; see 'go help modules'
------
executor failed running [/bin/sh -c go build -o main .]: exit code: 1
ERROR: Service 'api' failed to build : Build failed
```

GoMudules を利用するために Dockerfile に下記に修正し起動確認
空文字でもよいので[GOPATH= ]を指定
[参考動画](https://www.youtube.com/watch?v=rHontd51R3A&t=1219s)

```
FROM golang:alpine
WORKDIR /app/api
RUN export GOPATH= && export GO111MODULE=on
COPY ./api /app/api
RUN go build -o main .
CMD ["/app/api/main"]
```

?? パッケージをインストールする際はコンテナにログインして goget を実行すれば正しくパッケージ管理されているが、ホスト上で goget をするとパッケージが管理外になっている

- コンテナログイン

```
docker exec -it go_container /bin/sh
```

### MySQL

各種ファイル新規作成

gitignore

```
mysql/mysql_data
mysql/initdb.d
```

[参考サイト 1](https://michinoku-se.org/docker-mysql-workbench/)

上記の参考サイトから MySQLWorkBench からデータベース詳細を確認できる

## react 構築

### package

material-ui

```
yarn add @mui/material @emotion/react @emotion/styled
yarn add @mui/material @mui/styled-engine-sc styled-components
```

redux

```
yarn add connected-react-router history react-redux react-router redux redux-actions redux-logger redux-thunk reselect axios
```
