FROM golang:alpine
WORKDIR /app/api
RUN export GOPATH= && export GO111MODULE=on 
COPY ./api /app/api
RUN apk upgrade --update && \
  apk --no-cache add git gcc musl-dev
RUN go get github.com/go-sql-driver/mysql && \
  go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air
RUN go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv
CMD ["air", "-c", ".air.toml"]
