FROM golang:alpine

WORKDIR /app

COPY service2.go .

RUN go mod init service2 && go build -o service2 .

CMD ["./service2"]
