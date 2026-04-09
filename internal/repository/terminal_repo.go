package repository

import (
    "database/sql"
    "fmt"
    "lab2/internal/models"
)

type TerminalRepository struct {
    db *sql.DB
}

func NewTerminalRepository(db *sql.DB) *TerminalRepository {
    return &TerminalRepository{db: db}
}

func (r *TerminalRepository) Create(t *models.Terminal) error {
    query := `INSERT INTO terminals (serial_number, name, address, location, is_active) VALUES (?, ?, ?, ?, ?)`
    result, err := r.db.Exec(query, t.SerialNumber, t.Name, t.Address, t.Location, t.IsActive)
    if err != nil {
        return fmt.Errorf("failed to create terminal: %w", err)
    }
    id, _ := result.LastInsertId()
    t.ID = uint(id)
    return nil
}

func (r *TerminalRepository) GetByID(id uint) (*models.Terminal, error) {
    query := `SELECT id, serial_number, name, address, location, is_active, created_at, updated_at FROM terminals WHERE id = ?`
    row := r.db.QueryRow(query, id)
    
    var terminal models.Terminal
    err := row.Scan(&terminal.ID, &terminal.SerialNumber, &terminal.Name, &terminal.Address, &terminal.Location, &terminal.IsActive, &terminal.CreatedAt, &terminal.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get terminal: %w", err)
    }
    return &terminal, nil
}

func (r *TerminalRepository) GetBySerial(serial string) (*models.Terminal, error) {
    query := `SELECT id, serial_number, name, address, location, is_active, created_at, updated_at FROM terminals WHERE serial_number = ?`
    row := r.db.QueryRow(query, serial)
    
    var terminal models.Terminal
    err := row.Scan(&terminal.ID, &terminal.SerialNumber, &terminal.Name, &terminal.Address, &terminal.Location, &terminal.IsActive, &terminal.CreatedAt, &terminal.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get terminal: %w", err)
    }
    return &terminal, nil
}

func (r *TerminalRepository) Update(t *models.Terminal) error {
    query := `UPDATE terminals SET name = ?, address = ?, location = ?, is_active = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
    _, err := r.db.Exec(query, t.Name, t.Address, t.Location, t.IsActive, t.ID)
    if err != nil {
        return fmt.Errorf("failed to update terminal: %w", err)
    }
    return nil
}

func (r *TerminalRepository) Delete(id uint) error {
    query := `DELETE FROM terminals WHERE id = ?`
    _, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("failed to delete terminal: %w", err)
    }
    return nil
}

func (r *TerminalRepository) GetAll() ([]models.Terminal, error) {
    query := `SELECT id, serial_number, name, address, location, is_active, created_at, updated_at FROM terminals`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to get terminals: %w", err)
    }
    defer rows.Close()

    var terminals []models.Terminal
    for rows.Next() {
        var t models.Terminal
        err := rows.Scan(&t.ID, &t.SerialNumber, &t.Name, &t.Address, &t.Location, &t.IsActive, &t.CreatedAt, &t.UpdatedAt)
        if err != nil {
            return nil, err
        }
        terminals = append(terminals, t)
    }
    return terminals, nil
}