package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "lab2/internal/services"
)

// AuthMiddleware проверяет JWT токен
// @Summary JWT Auth Middleware
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
            c.Abort()
            return
        }

        token := strings.TrimSpace(authHeader)
        if strings.HasPrefix(token, "Bearer ") {
            token = strings.TrimPrefix(token, "Bearer ")
        }
        
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
            c.Abort()
            return
        }

        claims, err := authService.ValidateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        // Сохраняем данные пользователя в контексте
        c.Set("user_id", claims["user_id"])
        c.Set("login", claims["login"])
        c.Set("is_admin", claims["is_admin"])
        
        c.Next()
    }
}

// AdminMiddleware проверяет роль администратора
// @Summary Admin role check
func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        isAdmin, ok := c.Get("is_admin")
        if !ok || isAdmin != true {
            c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// TerminalAuthMiddleware для аутентификации терминалов
func TerminalAuthMiddleware(expectedToken string) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("X-Terminal-Token")
        if token != expectedToken {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid terminal token"})
            c.Abort()
            return
        }
        c.Next()
    }
}