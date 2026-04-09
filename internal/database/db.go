package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

func InitDB(dbPath string) error {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create db directory: %w", err)
	}

	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// 🔑 КРИТИЧНО для SQLite + Goose:
	// SQLite не поддерживает многопоточную запись, ограничиваем до 1 соединения
	DB.SetMaxOpenConns(1)
	DB.SetMaxIdleConns(1)

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	return nil
}

func RunMigrations(dbPath string) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database for migrations: %w", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)
	
	// 🔑 Явно указываем диалект, чтобы goose правильно парсил SQLite-синтаксис
	goose.SetDialect("sqlite3")

	migrationDir := "./internal/database/migrations"
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