.PHONY: build run migrate-up swagger docker-build

build:
	go build -o bin/lab2 ./cmd/server

run:
	godotenv -f .env ./bin/lab2 server

migrate-up:
	goose -dir internal/database/migrations sqlite3 $$DB_PATH up

swagger:
	swag init -g cmd/server/main.go -o internal/swagger/docs

docker-build:
	docker build -t lab2:latest -f deploy/Dockerfile .