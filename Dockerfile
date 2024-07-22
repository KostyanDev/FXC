FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Скопировать файл .env
COPY .env .env

# Установить утилиту envsubst для подстановки переменных окружения и bash
RUN apt-get update && apt-get install -y gettext-base bash

# Сборка приложения
RUN go build -o app ./cmd/main.go

# Передача переменных окружения в команду запуска
ENTRYPOINT ["/bin/bash", "-c", "source .env && env && ./app"]