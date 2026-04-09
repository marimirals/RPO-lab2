# Stage 1: Build Go binary
FROM golang:1.22-alpine AS builder
WORKDIR /app
ENV GOPROXY=direct
ENV GOSUMDB=off
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o lab2 ./cmd/server

# Stage 2: Runtime with Nginx
FROM alpine:latest
RUN apk --no-cache add nginx openssl sqlite-libs supervisor

# Create directories
RUN mkdir -p /app/bin /app/data /etc/nginx/ssl /usr/share/nginx/swagger

# Copy binary and entrypoint
COPY --from=builder /app/lab2 /app/bin/
COPY deploy/entrypoint.sh /entrypoint.sh
COPY deploy/nginx.conf /etc/nginx/nginx.conf
COPY deploy/ssl/server.crt /etc/nginx/ssl/
COPY deploy/ssl/server.key /etc/nginx/ssl/

# Copy swagger UI (после генерации)
# COPY internal/swagger/docs /usr/share/nginx/swagger/

# Expose port
EXPOSE 8888

# Use entrypoint
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
ENV GOSUMDB=off
ENV GOSUMDB=off