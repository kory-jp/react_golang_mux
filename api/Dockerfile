FROM golang:1.18.2-alpine3.14 as build
WORKDIR /app/api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o apiBuilder main.go

# FROM gcr.io/distroless/static:nonroot
# コンテナログインのためalpineを使用
FROM alpine as prod
WORKDIR /app
ENV GO_ENV production
COPY --from=build /app/api/apiBuilder /app/apiBuilder
COPY ./env/production.env /app/production.env
EXPOSE 8000
ENTRYPOINT [ "/app/apiBuilder" ]