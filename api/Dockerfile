FROM golang:alpine
WORKDIR /app/api
COPY . /app/api
RUN export GOPATH= && export GO111MODULE=on 
RUN apk upgrade --update && \
  apk --no-cache add git gcc musl-dev
RUN go mod download

RUN go get -d -v
RUN go install -v

EXPOSE 8000

CMD ["./main"]