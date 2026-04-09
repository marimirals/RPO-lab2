package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "lab2/internal/models"
    "lab2/internal/repository"
    "lab2/internal/services"
)

type UserHandler struct {
    repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
    return &UserHandler{repo: repo}
}

// GetAllUsers получение всех пользователей
// @Summary Получить всех пользователей
// @Description Только для администраторов
// @Tags Users
// @Produce json
// @Security Bearer
// @Success 200 {array} models.User "Список пользователей"
// @Failure 401 {object} models.ErrorResponse "Не авторизован"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
    users, err := h.repo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    // Убираем хеши паролей
    for i := range users {
        users[i].PasswordHash = ""
    }
    c.JSON(http.StatusOK, models.SuccessResponse(users))
}

// GetUserByID получение пользователя по ID
// @Summary Получить пользователя по ID
// @Description Пользователь может получить только свои данные, админ - любые
// @Tags Users
// @Produce json
// @Security Bearer
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.User "Данные пользователя"
// @Failure 401 {object} models.ErrorResponse "Не авторизован"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Failure 404 {object} models.ErrorResponse "Пользователь не найден"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    userID := c.GetUint("user_id")
    isAdmin := c.GetBool("is_admin")

    // Проверка прав: пользователь может смотреть только себя
    if !isAdmin && uint(id) != userID {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "access denied"})
        return
    }

    user, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "user not found"})
        return
    }

    user.PasswordHash = ""
    c.JSON(http.StatusOK, models.SuccessResponse(user))
}

// CreateUser создание пользователя (только админ)
// @Summary Создать пользователя
// @Description Только для администраторов
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body models.CreateUserRequest true "Данные пользователя"
// @Success 201 {object} models.User "Пользователь создан"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
    // Только админ может создавать пользователей
    if !c.GetBool("is_admin") {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "admin access required"})
        return
    }

    var req models.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Проверка на дубликат
    existing, _ := h.repo.GetByLogin(req.Login)
    if existing != nil {
        c.JSON(http.StatusConflict, models.ErrorResponse{Error: "user already exists"})
        return
    }

    passwordHash, err := services.HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "password hashing failed"})
        return
    }

    user := &models.User{
        Login:        req.Login,
        Name:         req.Name,
        PasswordHash: passwordHash,
        IsAdmin:      req.IsAdmin,
    }

    if err := h.repo.Create(user); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to create user"})
        return
    }

    user.PasswordHash = ""
    c.JSON(http.StatusCreated, models.SuccessResponse(user))
}

// UpdateUser обновление пользователя
// @Summary Обновить пользователя
// @Description Пользователь может редактировать только себя, админ - всех
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "ID пользователя"
// @Param request body models.UpdateUserRequest true "Новые данные"
// @Success 200 {object} models.User "Пользователь обновлен"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Failure 404 {object} models.ErrorResponse "Пользователь не найден"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    userID := c.GetUint("user_id")
    isAdmin := c.GetBool("is_admin")

    // Проверка прав
    if !isAdmin && uint(id) != userID {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "cannot edit other users"})
        return
    }

    var req models.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    user, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "user not found"})
        return
    }

    // Обновляем поля
    if req.Name != "" {
        user.Name = req.Name
    }
    if req.IsAdmin != nil {
        // Только админ может менять роль
        if isAdmin {
            user.IsAdmin = *req.IsAdmin
        }
    }

    if err := h.repo.Update(user); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to update user"})
        return
    }

    user.PasswordHash = ""
    c.JSON(http.StatusOK, models.SuccessResponse(user))
}

// DeleteUser удаление пользователя
// @Summary Удалить пользователя
// @Description Только для администраторов. Нельзя удалить самого себя
// @Tags Users
// @Produce json
// @Security Bearer
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.APIResponse "Пользователь удален"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Failure 404 {object} models.ErrorResponse "Пользователь не найден"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
    if !c.GetBool("is_admin") {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "admin access required"})
        return
    }

    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    userID := c.GetUint("user_id")

    // Нельзя удалить самого себя
    if uint(id) == userID {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "cannot delete yourself"})
        return
    }

    if err := h.repo.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, models.SuccessResponse(nil))
}