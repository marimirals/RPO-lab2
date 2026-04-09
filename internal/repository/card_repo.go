package repository

import (
    "database/sql"
    "fmt"
    "lab2/internal/models"
)

type CardRepository struct {
    db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
    return &CardRepository{db: db}
}

func (r *CardRepository) Create(card *models.Card) error {
    query := `INSERT INTO cards (card_number, balance, is_blocked, owner_name, key_id) VALUES (?, ?, ?, ?, ?)`
    result, err := r.db.Exec(query, card.CardNumber, card.Balance, card.IsBlocked, card.OwnerName, card.KeyID)
    if err != nil {
        return fmt.Errorf("failed to create card: %w", err)
    }
    id, _ := result.LastInsertId()
    card.ID = uint(id)
    return nil
}

func (r *CardRepository) GetByNumber(number string) (*models.Card, error) {
    query := `SELECT id, card_number, balance, is_blocked, owner_name, key_id, created_at, updated_at FROM cards WHERE card_number = ?`
    row := r.db.QueryRow(query, number)
    
    var card models.Card
    err := row.Scan(&card.ID, &card.CardNumber, &card.Balance, &card.IsBlocked, &card.OwnerName, &card.KeyID, &card.CreatedAt, &card.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get card: %w", err)
    }
    return &card, nil
}

func (r *CardRepository) GetByID(id uint) (*models.Card, error) {
    query := `SELECT id, card_number, balance, is_blocked, owner_name, key_id, created_at, updated_at FROM cards WHERE id = ?`
    row := r.db.QueryRow(query, id)
    
    var card models.Card
    err := row.Scan(&card.ID, &card.CardNumber, &card.Balance, &card.IsBlocked, &card.OwnerName, &card.KeyID, &card.CreatedAt, &card.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get card: %w", err)
    }
    return &card, nil
}

func (r *CardRepository) Update(card *models.Card) error {
    query := `UPDATE cards SET balance = ?, is_blocked = ?, owner_name = ?, key_id = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
    _, err := r.db.Exec(query, card.Balance, card.IsBlocked, card.OwnerName, card.KeyID, card.ID)
    if err != nil {
        return fmt.Errorf("failed to update card: %w", err)
    }
    return nil
}

func (r *CardRepository) Delete(id uint) error {
    query := `DELETE FROM cards WHERE id = ?`
    _, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("failed to delete card: %w", err)
    }
    return nil
}

func (r *CardRepository) GetAll() ([]models.Card, error) {
    query := `SELECT id, card_number, balance, is_blocked, owner_name, key_id, created_at, updated_at FROM cards`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to get cards: %w", err)
    }
    defer rows.Close()

    var cards []models.Card
    for rows.Next() {
        var c models.Card
        err := rows.Scan(&c.ID, &c.CardNumber, &c.Balance, &c.IsBlocked, &c.OwnerName, &c.KeyID, &c.CreatedAt, &c.UpdatedAt)
        if err != nil {
            return nil, err
        }
        cards = append(cards, c)
    }
    return cards, nil
}

// UpdateBalance атомарно обновляет баланс (для транзакций)
func (r *CardRepository) UpdateBalance(cardID uint, amount int64) error {
    query := `UPDATE cards SET balance = balance + ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
    result, err := r.db.Exec(query, amount, cardID)
    if err != nil {
        return fmt.Errorf("failed to update balance: %w", err)
    }
    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("card not found")
    }
    return nil
}