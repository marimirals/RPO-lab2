#!/bin/sh
set -e

echo "🔄 Applying database migrations..."
# Запускаем миграции через сам бинарник (если есть команда migrate)
# Или через goose, если он в образе. Для простоты - полагаемся на инициализацию в main.go

echo "🚀 Starting Go server on port 8080 (background)..."
./lab2 server &
GO_PID=$!

echo "🌐 Starting Nginx on port 8888 (foreground)..."
# Nginx в foreground, чтобы контейнер не падал
exec nginx -g "daemon off;"