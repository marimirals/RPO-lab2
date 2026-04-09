package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "path/filepath"

    _ "github.com/mattn/go-sqlite3"
    "github.com/pressly/goose/v3"
)

var DB *sql.DB

func InitDB(dbPath string) error {
    // Создаем директорию для БД если не существует
    dir := filepath.Dir(dbPath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return fmt.Errorf("failed to create db directory: %w", err)
    }

    var err error
    DB, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        return fmt.Errorf("failed to open database: %w", err)
    }

    // Проверка подключения
    if err = DB.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %w", err)
    }

    return nil
}

func RunMigrations(dbPath string) error {
    // Открываем соединение для миграций
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return fmt.Errorf("failed to open database for migrations: %w", err)
    }
    defer db.Close()

    // Путь к папке с миграциями
    migrationDir := "./internal/database/migrations"
    
    // Запускаем миграции вверх
    if err := goose.Up(db, migrationDir); err != nil {
        return fmt.Errorf("failed to run migrations: %w", err)
    }

    log.Println("Migrations completed successfully")
    return nil
}

func CloseDB() error {
    if DB != nil {
        return DB.Close()
    }
    return nil
}