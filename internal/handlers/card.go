package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "lab2/internal/models"
    "lab2/internal/repository"
)

type CardHandler struct {
    repo *repository.CardRepository
}

func NewCardHandler(repo *repository.CardRepository) *CardHandler {
    return &CardHandler{repo: repo}
}

// GetAllCards получение всех карт
// @Summary Получить все карты
// @Tags Cards
// @Produce json
// @Security Bearer
// @Success 200 {array} models.Card "Список карт"
// @Router /cards [get]
func (h *CardHandler) GetAllCards(c *gin.Context) {
    cards, err := h.repo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(cards))
}

// GetCardByID получение карты по ID
// @Summary Получить карту по ID
// @Tags Cards
// @Produce json
// @Security Bearer
// @Param id path int true "ID карты"
// @Success 200 {object} models.Card "Данные карты"
// @Failure 404 {object} models.ErrorResponse "Карта не найдена"
// @Router /cards/{id} [get]
func (h *CardHandler) GetCardByID(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    card, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if card == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "card not found"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(card))
}

// GetCardByNumber получение карты по номеру
// @Summary Получить карту по номеру
// @Tags Cards
// @Produce json
// @Security Bearer
// @Param number path string true "Номер карты"
// @Success 200 {object} models.Card "Данные карты"
// @Failure 404 {object} models.ErrorResponse "Карта не найдена"
// @Router /cards/number/{number} [get]
func (h *CardHandler) GetCardByNumber(c *gin.Context) {
    number := c.Param("number")
    
    card, err := h.repo.GetByNumber(number)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if card == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "card not found"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(card))
}

// CreateCard создание карты
// @Summary Создать карту
// @Tags Cards
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body models.CreateCardRequest true "Данные карты"
// @Success 201 {object} models.Card "Карта создана"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Router /cards [post]
func (h *CardHandler) CreateCard(c *gin.Context) {
    var req models.CreateCardRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    balance := int64(0)
    if req.Balance != nil {
        balance = *req.Balance
    }

    card := &models.Card{
        CardNumber: req.CardNumber,
        Balance:    balance,
        OwnerName:  req.OwnerName,
        KeyID:      req.KeyID,
        IsBlocked:  false,
    }

    if err := h.repo.Create(card); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to create card"})
        return
    }
    c.JSON(http.StatusCreated, models.SuccessResponse(card))
}

// UpdateCard обновление карты
// @Summary Обновить карту
// @Tags Cards
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "ID карты"
// @Param request body models.UpdateCardRequest true "Новые данные"
// @Success 200 {object} models.Card "Карта обновлена"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 404 {object} models.ErrorResponse "Карта не найдена"
// @Router /cards/{id} [put]
func (h *CardHandler) UpdateCard(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    var req models.UpdateCardRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    card, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if card == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "card not found"})
        return
    }

    if req.Balance != nil {
        card.Balance = *req.Balance
    }
    if req.IsBlocked != nil {
        card.IsBlocked = *req.IsBlocked
    }
    if req.OwnerName != "" {
        card.OwnerName = req.OwnerName
    }
    if req.KeyID != nil {
        card.KeyID = req.KeyID
    }

    if err := h.repo.Update(card); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to update card"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(card))
}

// DeleteCard удаление карты
// @Summary Удалить карту
// @Tags Cards
// @Produce json
// @Security Bearer
// @Param id path int true "ID карты"
// @Success 200 {object} models.APIResponse "Карта удалена"
// @Failure 404 {object} models.ErrorResponse "Карта не найдена"
// @Router /cards/{id} [delete]
func (h *CardHandler) DeleteCard(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    if err := h.repo.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to delete card"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(nil))
}