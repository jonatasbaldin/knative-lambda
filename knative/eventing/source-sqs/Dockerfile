FROM golang:1.13-alpine

WORKDIR /go/src/app
COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go build -o /go/bin/consumer-sqs

ENTRYPOINT ["/go/bin/consumer-sqs"]