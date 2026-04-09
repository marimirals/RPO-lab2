package services

import (
    "errors"
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

// Config для JWT
type JWTConfig struct {
    SecretKey     string
    AccessExpire  time.Duration
    RefreshExpire time.Duration
}

type AuthService struct {
    config JWTConfig
}

func NewAuthService(config JWTConfig) *AuthService {
    return &AuthService{config: config}
}

// HashPassword хеширует пароль через bcrypt
// @Summary Hash password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    return string(bytes), nil
}

// CheckPassword проверяет пароль против хеша
func CheckPassword(password, hash string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GenerateTokens создает access и refresh токены
// @Summary Generate JWT tokens
// @Produce json
// @Param user_id path int true "User ID"
// @Param is_admin path bool true "Is admin"
// @Success 200 {object} models.LoginResponse
func (s *AuthService) GenerateTokens(userID uint, login string, isAdmin bool) (string, string, error) {
    accessToken, err := s.createToken(userID, login, isAdmin, s.config.AccessExpire)
    if err != nil {
        return "", "", err
    }
    
    refreshToken, err := s.createToken(userID, login, isAdmin, s.config.RefreshExpire)
    if err != nil {
        return "", "", err
    }
    
    return accessToken, refreshToken, nil
}

// createToken создает JWT токен
func (s *AuthService) createToken(userID uint, login string, isAdmin bool, expire time.Duration) (string, error) {
    claims := jwt.MapClaims{
        "user_id":  userID,
        "login":    login,
        "is_admin": isAdmin,
        "exp":      time.Now().Add(expire).Unix(),
        "iat":      time.Now().Unix(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.config.SecretKey))
}

// ValidateToken проверяет и парсит JWT токен
// @Summary Validate JWT token
// @Param token header string true "Bearer token"
func (s *AuthService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(s.config.SecretKey), nil
    })
    
    if err != nil {
        return nil, fmt.Errorf("invalid token: %w", err)
    }
    
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, errors.New("invalid token claims")
}

// ValidateTerminalToken проверяет токен терминала (упрощенная версия)
func (s *AuthService) ValidateTerminalToken(token, expectedToken string) bool {
    return token == expectedToken
}