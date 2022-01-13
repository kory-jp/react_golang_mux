FROM golang:alpine
WORKDIR /app/api
RUN export GOPATH= && export GO111MODULE=on
COPY ./api /app/api
RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air
CMD ["air", "-c", ".air.toml"]
