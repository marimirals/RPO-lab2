#!/bin/bash

# Переходим в директорию со скриптом
cd "$(dirname "$0")"

# Генерируем приватный ключ и самоподписанный сертификат
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout server.key \
  -out server.crt \
  -subj "//C=RU/ST=Moscow/L=Moscow/O=Lab2/OU=BMSTU/CN=localhost"

# Проверяем, что файлы созданы
if [ -f "server.crt" ] && [ -f "server.key" ]; then
    echo "SSL certificates generated successfully:"
    echo "  - server.crt"
    echo "  - server.key"
    chmod 600 server.key
    chmod 644 server.crt
else
    echo "Error: Failed to generate certificates"
    exit 1
fi