package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "lab2/internal/models"
    "lab2/internal/repository"
)

type TransactionHandler struct {
    repo *repository.TransactionRepository
}

func NewTransactionHandler(repo *repository.TransactionRepository) *TransactionHandler {
    return &TransactionHandler{repo: repo}
}

// GetTransactionByID получение транзакции по ID
// @Summary Получить транзакцию по ID
// @Tags Transactions
// @Produce json
// @Security Bearer
// @Param id path int true "ID транзакции"
// @Success 200 {object} models.Transaction "Данные транзакции"
// @Failure 404 {object} models.ErrorResponse "Транзакция не найдена"
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    
    tx, err := h.repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    if tx == nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "transaction not found"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(tx))
}

// GetTransactionsByCard получение транзакций по карте
// @Summary Получить транзакции по карте
// @Tags Transactions
// @Produce json
// @Security Bearer
// @Param card_id path int true "ID карты"
// @Success 200 {array} models.Transaction "Список транзакций"
// @Router /transactions/card/{card_id} [get]
func (h *TransactionHandler) GetTransactionsByCard(c *gin.Context) {
    cardID, _ := strconv.ParseUint(c.Param("card_id"), 10, 32)
    
    txs, err := h.repo.GetByCardID(uint(cardID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(txs))
}

// UpdateTransactionStatus обновление статуса транзакции
// @Summary Обновить статус транзакции
// @Tags Transactions
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "ID транзакции"
// @Param status query string true "Новый статус"
// @Success 200 {object} models.APIResponse "Статус обновлен"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Router /transactions/{id}/status [put]
func (h *TransactionHandler) UpdateTransactionStatus(c *gin.Context) {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    status := c.Query("status")
    
    if status == "" {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "status parameter required"})
        return
    }

    if err := h.repo.UpdateStatus(uint(id), status); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to update status"})
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse(nil))
}