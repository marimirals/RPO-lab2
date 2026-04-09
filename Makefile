.PHONY: build run migrate-up migrate-down swagger docker-build docker-run
export CGO_ENABLED=0 

# Сборка бинарника
build:
	go build -o bin/lab2 ./cmd/server

# Запуск локально (требуется .env)
run:
	@echo "Starting server on port $${SERVER_PORT:-8080}..."
	go run ./cmd/server

# Применение миграций
migrate-up:
	goose -dir internal/database/migrations sqlite $${DB_PATH:-./data/lab2.db} up

# Откат миграций
migrate-down:
	goose -dir internal/database/migrations sqlite $${DB_PATH:-./data/lab2.db} down

# Генерация Swagger документации
swagger:
	swag init -g cmd/server/main.go -o internal/swagger/docs

# Сборка Docker образа
docker-build:
	docker build -t lab2:latest -f deploy/Dockerfile .

# Запуск в Docker
docker-run:
	docker run -d \
		--name lab2 \
		-p 8888:8888 \
		-v $$(pwd)/data:/app/data \
		-e DB_PATH=/app/data/lab2.db \
		-e JWT_SECRET=$${JWT_SECRET:-dev-secret} \
		-e TERMINAL_TOKEN=$${TERMINAL_TOKEN:-terminal-token} \
		lab2:latest

# Остановка контейнера
docker-stop:
	docker stop lab2 && docker rm lab2

# Тесты
test:
	go test ./... -v

# Линтер
lint:
	golangci-lint run ./...