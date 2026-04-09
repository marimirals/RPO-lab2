package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "lab2/internal/models"
    "lab2/internal/repository"
    "lab2/internal/services"
)

type AuthHandler struct {
    authService *services.AuthService
    userRepo    *repository.UserRepository
}

func NewAuthHandler(authService *services.AuthService, userRepo *repository.UserRepository) *AuthHandler {
    return &AuthHandler{authService: authService, userRepo: userRepo}
}

// Login авторизация пользователя
// @Summary Авторизация пользователя
// @Description Вход по логину и паролю, возврат JWT токенов
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Данные для входа"
// @Success 200 {object} models.LoginResponse "Успешный вход"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 401 {object} models.ErrorResponse "Неверный логин или пароль"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    var req models.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Находим пользователя
    user, err := h.userRepo.GetByLogin(req.Login)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if user == nil {
        c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "invalid credentials"})
        return
    }

    // Проверяем пароль
    if err := services.CheckPassword(req.Password, user.PasswordHash); err != nil {
        c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "invalid credentials"})
        return
    }

    // Генерируем токены
    accessToken, refreshToken, err := h.authService.GenerateTokens(user.ID, user.Login, user.IsAdmin)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "token generation failed"})
        return
    }

    // Возвращаем ответ (без пароля)
    response := models.LoginResponse{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        User: models.User{
            ID:        user.ID,
            Login:     user.Login,
            Name:      user.Name,
            IsAdmin:   user.IsAdmin,
            CreatedAt: user.CreatedAt,
            UpdatedAt: user.UpdatedAt,
        },
    }

    c.JSON(http.StatusOK, response)
}

// Register регистрация нового пользователя
// @Summary Регистрация пользователя
// @Description Создание нового аккаунта
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "Данные пользователя"
// @Success 201 {object} models.User "Пользователь создан"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 409 {object} models.ErrorResponse "Пользователь уже существует"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
    var req models.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Проверяем, не занят ли логин
    existing, _ := h.userRepo.GetByLogin(req.Login)
    if existing != nil {
        c.JSON(http.StatusConflict, models.ErrorResponse{Error: "user already exists"})
        return
    }

    // Хешируем пароль
    passwordHash, err := services.HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "password hashing failed"})
        return
    }

    // Создаем пользователя
    user := &models.User{
        Login:        req.Login,
        Name:         req.Name,
        PasswordHash: passwordHash,
        IsAdmin:      req.IsAdmin,
    }

    if err := h.userRepo.Create(user); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to create user"})
        return
    }

    // Возвращаем пользователя без пароля
    user.PasswordHash = ""
    c.JSON(http.StatusCreated, user)
}