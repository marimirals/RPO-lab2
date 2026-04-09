-- +goose Up
-- +goose StatementBegin

-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    is_admin INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Таблица ключей
CREATE TABLE IF NOT EXISTS keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key_value TEXT NOT NULL,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Таблица терминалов
CREATE TABLE IF NOT EXISTS terminals (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    serial_number TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    address TEXT,
    location TEXT,
    is_active INTEGER DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Таблица транспортных карт
CREATE TABLE IF NOT EXISTS cards (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    card_number TEXT UNIQUE NOT NULL,
    balance INTEGER DEFAULT 0,
    is_blocked INTEGER DEFAULT 0,
    owner_name TEXT,
    key_id INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (key_id) REFERENCES keys(id)
);

-- Таблица транзакций
CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    card_id INTEGER NOT NULL,
    terminal_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    transaction_type TEXT DEFAULT 'payment',
    status TEXT DEFAULT 'pending',
    transaction_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (card_id) REFERENCES cards(id),
    FOREIGN KEY (terminal_id) REFERENCES terminals(id)
);

-- Индексы
CREATE INDEX IF NOT EXISTS idx_cards_number ON cards(card_number);
CREATE INDEX IF NOT EXISTS idx_transactions_card ON transactions(card_id);
CREATE INDEX IF NOT EXISTS idx_transactions_terminal ON transactions(terminal_id);
CREATE INDEX IF NOT EXISTS idx_users_login ON users(login);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS idx_users_login;
DROP INDEX IF EXISTS idx_transactions_terminal;
DROP INDEX IF EXISTS idx_transactions_card;
DROP INDEX IF EXISTS idx_cards_number;

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS terminals;
DROP TABLE IF EXISTS keys;
DROP TABLE IF EXISTS users;

-- +goose StatementEnd