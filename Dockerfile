FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env .env

RUN apt-get update && apt-get install -y gettext-base bash

RUN go build -o app ./cmd/main.go

ENTRYPOINT ["/bin/bash", "-c", "source .env && env && ./app"]