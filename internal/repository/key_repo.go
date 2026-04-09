package repository

import (
    "database/sql"
    "fmt"
    "lab2/internal/models"
)

type KeyRepository struct {
    db *sql.DB
}

func NewKeyRepository(db *sql.DB) *KeyRepository {
    return &KeyRepository{db: db}
}

func (r *KeyRepository) Create(key *models.Key) error {
    query := `INSERT INTO keys (key_value, description) VALUES (?, ?)`
    result, err := r.db.Exec(query, key.KeyValue, key.Description)
    if err != nil {
        return fmt.Errorf("failed to create key: %w", err)
    }
    id, _ := result.LastInsertId()
    key.ID = uint(id)
    return nil
}

func (r *KeyRepository) GetByID(id uint) (*models.Key, error) {
    query := `SELECT id, key_value, description, created_at FROM keys WHERE id = ?`
    row := r.db.QueryRow(query, id)
    
    var key models.Key
    err := row.Scan(&key.ID, &key.KeyValue, &key.Description, &key.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get key: %w", err)
    }
    return &key, nil
}

func (r *KeyRepository) Update(key *models.Key) error {
    query := `UPDATE keys SET key_value = ?, description = ? WHERE id = ?`
    _, err := r.db.Exec(query, key.KeyValue, key.Description, key.ID)
    if err != nil {
        return fmt.Errorf("failed to update key: %w", err)
    }
    return nil
}

func (r *KeyRepository) Delete(id uint) error {
    query := `DELETE FROM keys WHERE id = ?`
    _, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("failed to delete key: %w", err)
    }
    return nil
}

func (r *KeyRepository) GetAll() ([]models.Key, error) {
    query := `SELECT id, key_value, description, created_at FROM keys`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to get keys: %w", err)
    }
    defer rows.Close()

    var keys []models.Key
    for rows.Next() {
        var k models.Key
        err := rows.Scan(&k.ID, &k.KeyValue, &k.Description, &k.CreatedAt)
        if err != nil {
            return nil, err
        }
        keys = append(keys, k)
    }
    return keys, nil
}