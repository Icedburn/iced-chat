FROM golang:1.14

WORKDIR /go/src/app

COPY . .

RUN go build -i -o ./bin/IcedChat ./...

CMD ["./bin/IcedChat"]