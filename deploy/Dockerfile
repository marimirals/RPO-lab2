# ===== Stage 1: Сборка Go-бинарника =====
FROM golang:1.22-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOPROXY=direct
ENV GOSUMDB=off

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарник с именем "lab2" в корне /app
RUN go build -o lab2 ./cmd/server

# ===== Stage 2: Runtime с Nginx =====
FROM alpine:latest

RUN apk --no-cache add nginx openssl curl ca-certificates

WORKDIR /app

# Создаем директории
RUN mkdir -p /app/data /etc/nginx/ssl /usr/share/nginx/swagger

# Копируем бинарник (из builder, из /app/lab2)
COPY --from=builder /app/lab2 ./lab2

# Копируем скрипты и конфиги
COPY deploy/entrypoint.sh /entrypoint.sh
COPY deploy/nginx.conf /etc/nginx/nginx.conf
COPY deploy/ssl/server.crt /etc/nginx/ssl/server.crt
COPY deploy/ssl/server.key /etc/nginx/ssl/server.key

# Копируем Swagger UI (ОБЯЗАТЕЛЬНО раскомментируй!)
COPY internal/swagger/docs /usr/share/nginx/swagger/

# Делаем скрипты исполняемыми
RUN chmod +x /entrypoint.sh ./lab2

EXPOSE 8888

ENTRYPOINT ["/entrypoint.sh"]