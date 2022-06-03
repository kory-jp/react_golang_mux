FROM golang:1.18.2-alpine3.14 as build
RUN export GOPATH= && export GO111MODULE=on
WORKDIR /app/api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o apiBuilder main.go

# サンプルデータ投入のためのビルド
# サンプルデータの投入実行はコンテナにログインして手動で実行
# ./seedBuilder
WORKDIR /app/api/infrastructure/seeds/development
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o seedBuilder seeder.go


# FROM gcr.io/distroless/static:nonroot
# コンテナログインのためalpineを使用
FROM alpine as prod
WORKDIR /app
ENV GO_ENV production
COPY --from=build /app/api/apiBuilder /app/apiBuilder
COPY --from=build /app/api/infrastructure/seeds/development/seedBuilder /app/seedBuilder
COPY ./env/production.env /app/production.env
EXPOSE 8000
ENTRYPOINT [ "/app/apiBuilder" ]