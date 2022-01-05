# react_golang_mux

mux を利用して react と golang の通信、cors 設定
docker を利用しての環境構築

## 環境構築

- ディレクトリ構成

```
./-- api //golang
  |_client //react-tyepscript
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
