FROM golang:1.20.4

WORKDIR /app

ADD . .

WORKDIR /app/cmd/ledger

RUN go build -o /app/cmd/ledger/main .

EXPOSE 8080
