package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "lab2/internal/models"
    "lab2/internal/repository"
)

type KeyHandler struct {
    repo *repository.KeyRepository
}

func NewKeyHandler(repo *repository.KeyRepository) *KeyHandler {
    return &KeyHandler{repo: repo}
}

// GetAllKeys получение всех ключей
// @Summary Получить все ключи
// @Description Только для администраторов
// @Tags Keys
// @Produce json
// @Security Bearer
// @Success 200 {array} models.Key "Список ключей"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Router /keys [get]
func (h *KeyHandler) GetAllKeys(c *gin.Context) {
    if !c.GetBool("is_admin") {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "admin access required"})
        return
    }
    
    keys, err := h.repo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(keys))
}

// GetKeyByID получение ключа по ID
// @Summary Получить ключ по ID
// @Description Только для администраторов
// @Tags Keys
// @Produce json
// @Security Bearer
// @Param id path int true "ID ключа"
// @Success 200 {object} models.Key "Данные ключа"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Failure 404 {object} models.ErrorResponse "Ключ не найден"
// @Router /keys/{id} [get]
func (h *KeyHandler) GetKeyByID(c *gin.Context) {
    if !c.GetBool("is_admin") {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "admin access required"})
        return
    }
    
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    key, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if key == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "key not found"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(key))
}

// CreateKey создание ключа
// @Summary Создать ключ
// @Description Только для администраторов
// @Tags Keys
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body models.CreateKeyRequest true "Данные ключа"
// @Success 201 {object} models.Key "Ключ создан"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Router /keys [post]
func (h *KeyHandler) CreateKey(c *gin.Context) {
    if !c.GetBool("is_admin") {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "admin access required"})
        return
    }
    
    var req models.CreateKeyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    key := &models.Key{
        KeyValue:    req.KeyValue,
        Description: req.Description,
    }

    if err := h.repo.Create(key); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to create key"})
        return
    }
    c.JSON(http.StatusCreated, models.SuccessResponse(key))
}

// UpdateKey обновление ключа
// @Summary Обновить ключ
// @Description Только для администраторов
// @Tags Keys
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "ID ключа"
// @Param request body models.UpdateKeyRequest true "Новые данные"
// @Success 200 {object} models.Key "Ключ обновлен"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Failure 404 {object} models.ErrorResponse "Ключ не найден"
// @Router /keys/{id} [put]
func (h *KeyHandler) UpdateKey(c *gin.Context) {
    if !c.GetBool("is_admin") {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "admin access required"})
        return
    }
    
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    var req models.UpdateKeyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    key, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if key == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "key not found"})
        return
    }

    if req.KeyValue != "" {
        key.KeyValue = req.KeyValue
    }
    if req.Description != "" {
        key.Description = req.Description
    }

    if err := h.repo.Update(key); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to update key"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(key))
}

// DeleteKey удаление ключа
// @Summary Удалить ключ
// @Description Только для администраторов
// @Tags Keys
// @Produce json
// @Security Bearer
// @Param id path int true "ID ключа"
// @Success 200 {object} models.APIResponse "Ключ удален"
// @Failure 403 {object} models.ErrorResponse "Нет прав"
// @Failure 404 {object} models.ErrorResponse "Ключ не найден"
// @Router /keys/{id} [delete]
func (h *KeyHandler) DeleteKey(c *gin.Context) {
    if !c.GetBool("is_admin") {
        c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "admin access required"})
        return
    }
    
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    if err := h.repo.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to delete key"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(nil))
}