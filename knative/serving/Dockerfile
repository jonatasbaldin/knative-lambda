FROM golang:1.13-alpine

WORKDIR /go/src/app
COPY main.go .

RUN go build -o /go/bin/quarantine

CMD ["/go/bin/quarantine"]