# ===== Stage 1: Сборка Go-бинарника =====
FROM golang:1.25-alpine AS builder

RUN apk add git

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOPROXY=direct
ENV GOSUMDB=off

COPY go.mod go.sum ./

COPY . .

# Собираем бинарник с именем "lab2" в корне /app
RUN go build -mod=vendor -o lab2 ./cmd/server

# ===== Stage 2: Runtime с Nginx =====
FROM alpine:3.18

RUN apk add nginx openssl curl ca-certificates bash

WORKDIR /lab2

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
# COPY internal/swagger/docs/ /usr/share/nginx/swagger/

#RUN ln -sf /usr/share/nginx/swagger/swagger.json /usr/share/nginx/swagger/doc.json

COPY internal/database/migrations/ /lab2/internal/database/migrations/

# Делаем скрипты исполняемыми
RUN chmod +x /entrypoint.sh ./lab2

EXPOSE 8888 8080

ENTRYPOINT ["/entrypoint.sh"]