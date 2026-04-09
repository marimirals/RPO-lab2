package repository

import (
    "database/sql"
    "fmt"
    "lab2/internal/models"
)

type TransactionRepository struct {
    db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
    return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(tx *models.Transaction) error {
    query := `INSERT INTO transactions (card_id, terminal_id, amount, transaction_type, status) VALUES (?, ?, ?, ?, ?)`
    result, err := r.db.Exec(query, tx.CardID, tx.TerminalID, tx.Amount, tx.TransactionType, tx.Status)
    if err != nil {
        return fmt.Errorf("failed to create transaction: %w", err)
    }
    id, _ := result.LastInsertId()
    tx.ID = uint(id)
    return nil
}

func (r *TransactionRepository) GetByID(id uint) (*models.Transaction, error) {
    query := `SELECT id, card_id, terminal_id, amount, transaction_type, status, transaction_time FROM transactions WHERE id = ?`
    row := r.db.QueryRow(query, id)
    
    var tx models.Transaction
    err := row.Scan(&tx.ID, &tx.CardID, &tx.TerminalID, &tx.Amount, &tx.TransactionType, &tx.Status, &tx.TransactionTime)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get transaction: %w", err)
    }
    return &tx, nil
}

func (r *TransactionRepository) UpdateStatus(id uint, status string) error {
    query := `UPDATE transactions SET status = ? WHERE id = ?`
    _, err := r.db.Exec(query, status, id)
    if err != nil {
        return fmt.Errorf("failed to update transaction status: %w", err)
    }
    return nil
}

func (r *TransactionRepository) GetByCardID(cardID uint) ([]models.Transaction, error) {
    query := `SELECT id, card_id, terminal_id, amount, transaction_type, status, transaction_time FROM transactions WHERE card_id = ? ORDER BY transaction_time DESC`
    rows, err := r.db.Query(query, cardID)
    if err != nil {
        return nil, fmt.Errorf("failed to get transactions: %w", err)
    }
    defer rows.Close()

    var txs []models.Transaction
    for rows.Next() {
        var t models.Transaction
        err := rows.Scan(&t.ID, &t.CardID, &t.TerminalID, &t.Amount, &t.TransactionType, &t.Status, &t.TransactionTime)
        if err != nil {
            return nil, err
        }
        txs = append(txs, t)
    }
    return txs, nil
}