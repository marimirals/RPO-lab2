package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "lab2/internal/models"
    "lab2/internal/repository"
)

type TerminalAPIHandler struct {
    cardRepo       *repository.CardRepository
    terminalRepo   *repository.TerminalRepository
    transactionRepo *repository.TransactionRepository
    keyRepo        *repository.KeyRepository
    terminalToken  string // токен для аутентификации терминалов
}

func NewTerminalAPIHandler(
    cardRepo *repository.CardRepository,
    terminalRepo *repository.TerminalRepository,
    transactionRepo *repository.TransactionRepository,
    keyRepo *repository.KeyRepository,
    terminalToken string,
) *TerminalAPIHandler {
    return &TerminalAPIHandler{
        cardRepo:        cardRepo,
        terminalRepo:    terminalRepo,
        transactionRepo: transactionRepo,
        keyRepo:         keyRepo,
        terminalToken:   terminalToken,
    }
}

// AuthorizeTransaction авторизация платежа от терминала
// @Summary Авторизация платежной транзакции
// @Description Терминал отправляет данные карты и суммы, сервер проверяет и возвращает результат
// @Tags TerminalAPI
// @Accept json
// @Produce json
// @Param X-Terminal-Token header string true "Токен терминала"
// @Param request body models.TerminalAuthRequest true "Данные транзакции"
// @Success 200 {object} models.TerminalAuthResponse "Результат авторизации"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 401 {object} models.ErrorResponse "Неверный токен терминала"
// @Router /terminal/auth [post]
func (h *TerminalAPIHandler) AuthorizeTransaction(c *gin.Context) {
    // Проверка токена терминала
    token := c.GetHeader("X-Terminal-Token")
    if token != h.terminalToken {
        c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "invalid terminal token"})
        return
    }

    var req models.TerminalAuthRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Проверяем терминал
    terminal, err := h.terminalRepo.GetByID(req.TerminalID)
    if err != nil || terminal == nil || !terminal.IsActive {
        c.JSON(http.StatusBadRequest, models.TerminalAuthResponse{
            Authorized: false,
            Message:    "invalid terminal",
        })
        return
    }

    // Находим карту по номеру
    card, err := h.cardRepo.GetByNumber(req.CardNumber)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.TerminalAuthResponse{
            Authorized: false,
            Message:    "database error",
        })
        return
    }
    if card == nil {
        c.JSON(http.StatusOK, models.TerminalAuthResponse{
            Authorized: false,
            Message:    "card not found",
        })
        return
    }

    // Проверяем блокировку
    if card.IsBlocked {
        c.JSON(http.StatusOK, models.TerminalAuthResponse{
            Authorized: false,
            Message:    "card is blocked",
            Balance:    card.Balance,
        })
        return
    }

    // Проверяем баланс (в копейках)
    if card.Balance < req.Amount {
        c.JSON(http.StatusOK, models.TerminalAuthResponse{
            Authorized: false,
            Message:    "insufficient funds",
            Balance:    card.Balance,
        })
        return
    }

    // ✅ Авторизация успешна - списываем сумму
    if err := h.cardRepo.UpdateBalance(card.ID, -req.Amount); err != nil {
        c.JSON(http.StatusInternalServerError, models.TerminalAuthResponse{
            Authorized: false,
            Message:    "failed to process transaction",
        })
        return
    }

    // Создаем запись о транзакции
    transaction := &models.Transaction{
        CardID:          card.ID,
        TerminalID:      req.TerminalID,
        Amount:          req.Amount,
        TransactionType: "payment",
        Status:          "completed",
        TransactionTime: time.Now(),
    }
    if err := h.transactionRepo.Create(transaction); err != nil {
        // Не прерываем, т.к. платеж уже прошел
        // Но логируем ошибку (в реальном проекте)
    }

    // Получаем актуальный баланс
    updatedCard, _ := h.cardRepo.GetByID(card.ID)
    newBalance := updatedCard.Balance

    c.JSON(http.StatusOK, models.TerminalAuthResponse{
        Authorized: true,
        Message:    "transaction approved",
        Balance:    newBalance,
    })
}

// DownloadKeys загрузка ключей для терминала
// @Summary Загрузка ключей шифрования
// @Description Возвращает все ключи для расшифровки данных карт
// @Tags TerminalAPI
// @Produce json
// @Param X-Terminal-Token header string true "Токен терминала"
// @Success 200 {array} models.Key "Список ключей"
// @Failure 401 {object} models.ErrorResponse "Неверный токен терминала"
// @Router /terminal/keys [get]
func (h *TerminalAPIHandler) DownloadKeys(c *gin.Context) {
    // Проверка токена терминала
    token := c.GetHeader("X-Terminal-Token")
    if token != h.terminalToken {
        c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "invalid terminal token"})
        return
    }

    keys, err := h.keyRepo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "database error"})
        return
    }

    c.JSON(http.StatusOK, models.SuccessResponse(keys))
}