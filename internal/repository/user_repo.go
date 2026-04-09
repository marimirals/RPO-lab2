package repository

import (
    "database/sql"
    "fmt"

    "lab2/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

// Create создает нового пользователя
// @Summary Create user
func (r *UserRepository) Create(user *models.User) error {
    query := `INSERT INTO users (login, name, password_hash, is_admin) VALUES (?, ?, ?, ?)`
    result, err := r.db.Exec(query, user.Login, user.Name, user.PasswordHash, user.IsAdmin)
    if err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }
    id, _ := result.LastInsertId()
    user.ID = uint(id)
    return nil
}

// GetByLogin находит пользователя по логину
func (r *UserRepository) GetByLogin(login string) (*models.User, error) {
    query := `SELECT id, login, name, password_hash, is_admin, created_at, updated_at FROM users WHERE login = ?`
    row := r.db.QueryRow(query, login)
    
    var user models.User
    err := row.Scan(&user.ID, &user.Login, &user.Name, &user.PasswordHash, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    return &user, nil
}

// GetByID находит пользователя по ID
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
    query := `SELECT id, login, name, password_hash, is_admin, created_at, updated_at FROM users WHERE id = ?`
    row := r.db.QueryRow(query, id)
    
    var user models.User
    err := row.Scan(&user.ID, &user.Login, &user.Name, &user.PasswordHash, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    return &user, nil
}

// Update обновляет пользователя
func (r *UserRepository) Update(user *models.User) error {
    query := `UPDATE users SET name = ?, is_admin = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
    _, err := r.db.Exec(query, user.Name, user.IsAdmin, user.ID)
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }
    return nil
}

// Delete удаляет пользователя
func (r *UserRepository) Delete(id uint) error {
    query := `DELETE FROM users WHERE id = ?`
    _, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("failed to delete user: %w", err)
    }
    return nil
}

// GetAll возвращает всех пользователей (только для админов)
func (r *UserRepository) GetAll() ([]models.User, error) {
    query := `SELECT id, login, name, is_admin, created_at, updated_at FROM users`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to get users: %w", err)
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var u models.User
        err := rows.Scan(&u.ID, &u.Login, &u.Name, &u.IsAdmin, &u.CreatedAt, &u.UpdatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, nil
}