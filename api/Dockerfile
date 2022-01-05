FROM golang:alpine
COPY ./api /app/api
WORKDIR /app/api
RUN go mod init github.com/kory-jp/react_golang_mux/api && go build -o main .
CMD ["/app/api/main"]