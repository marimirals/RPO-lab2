package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "lab2/internal/models"
    "lab2/internal/repository"
)

type TerminalHandler struct {
    repo *repository.TerminalRepository
}

func NewTerminalHandler(repo *repository.TerminalRepository) *TerminalHandler {
    return &TerminalHandler{repo: repo}
}

// GetAllTerminals получение всех терминалов
// @Summary Получить все терминалы
// @Tags Terminals
// @Produce json
// @Security Bearer
// @Success 200 {array} models.Terminal "Список терминалов"
// @Router /terminals [get]
func (h *TerminalHandler) GetAllTerminals(c *gin.Context) {
    terminals, err := h.repo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(terminals))
}

// GetTerminalByID получение терминала по ID
// @Summary Получить терминал по ID
// @Tags Terminals
// @Produce json
// @Security Bearer
// @Param id path int true "ID терминала"
// @Success 200 {object} models.Terminal "Данные терминала"
// @Failure 404 {object} models.ErrorResponse "Терминал не найден"
// @Router /terminals/{id} [get]
func (h *TerminalHandler) GetTerminalByID(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    terminal, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if terminal == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "terminal not found"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(terminal))
}

// CreateTerminal создание терминала
// @Summary Создать терминал
// @Tags Terminals
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body models.CreateTerminalRequest true "Данные терминала"
// @Success 201 {object} models.Terminal "Терминал создан"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Router /terminals [post]
func (h *TerminalHandler) CreateTerminal(c *gin.Context) {
    var req models.CreateTerminalRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    terminal := &models.Terminal{
        SerialNumber: req.SerialNumber,
        Name:         req.Name,
        Address:      req.Address,
        Location:     req.Location,
        IsActive:     true,
    }

    if err := h.repo.Create(terminal); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to create terminal"})
        return
    }
    c.JSON(http.StatusCreated, models.SuccessResponse(terminal))
}

// UpdateTerminal обновление терминала
// @Summary Обновить терминал
// @Tags Terminals
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "ID терминала"
// @Param request body models.UpdateTerminalRequest true "Новые данные"
// @Success 200 {object} models.Terminal "Терминал обновлен"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 404 {object} models.ErrorResponse "Терминал не найден"
// @Router /terminals/{id} [put]
func (h *TerminalHandler) UpdateTerminal(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    var req models.UpdateTerminalRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.NewErrorResponse(err.Error()))
        return
    }

    terminal, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.NewErrorResponse("database error"))
        return
    }
    if terminal == nil {
        c.JSON(http.StatusNotFound, models.NewErrorResponse("terminal not found"))
        return
    }

    if req.Name != "" {
        terminal.Name = req.Name
    }
    if req.Address != "" {
        terminal.Address = req.Address
    }
    if req.Location != "" {
        terminal.Location = req.Location
    }
    if req.IsActive != nil {
        terminal.IsActive = *req.IsActive
    }

    if err := h.repo.Update(terminal); err != nil {
        c.JSON(http.StatusInternalServerError, models.NewErrorResponse("failed to update terminal"))
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(terminal))
}

// DeleteTerminal удаление терминала
// @Summary Удалить терминал
// @Tags Terminals
// @Produce json
// @Security Bearer
// @Param id path int true "ID терминала"
// @Success 200 {object} models.APIResponse "Терминал удален"
// @Failure 404 {object} models.ErrorResponse "Терминал не найден"
// @Router /terminals/{id} [delete]
func (h *TerminalHandler) DeleteTerminal(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    if err := h.repo.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, models.NewErrorResponse("failed to delete terminal"))
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(nil))
}