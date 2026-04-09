// @title Payment Authorization API
// @version 1.0
// @description REST API для авторизации платежей транспортными картами
// @host localhost:8888
// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @schemes http https

package main

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "lab2/internal/config"
    "lab2/internal/database"
    "lab2/internal/handlers"
    "lab2/internal/middleware"
    "lab2/internal/repository"
    "lab2/internal/services"
    "lab2/internal/swagger"
)

// HealthResponse ответ health check
type HealthResponse struct {
    Status string `json:"status" example:"ok"`
}

// healthCheck проверка работоспособности API
// @Summary Health check
// @Description Проверка работоспособности API
// @Tags System
// @Produce json
// @Success 200 {object} HealthResponse "Статус сервера"
// @Router /health [get]
func healthCheck(c *gin.Context) {
    c.JSON(200, HealthResponse{Status: "ok"})
}

func main() {
    // Загружаем переменные окружения из .env
    _ = godotenv.Load()

    // Загружаем конфигурацию
    cfg := config.Load()

    // Инициализируем БД
    if err := database.InitDB(cfg.DatabasePath); err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer database.CloseDB()

    // Применяем миграции
    if err := database.RunMigrations(cfg.DatabasePath); err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }

    // Создаем репозитории
    userRepo := repository.NewUserRepository(database.DB)
    cardRepo := repository.NewCardRepository(database.DB)
    terminalRepo := repository.NewTerminalRepository(database.DB)
    keyRepo := repository.NewKeyRepository(database.DB)
    transactionRepo := repository.NewTransactionRepository(database.DB)

    // Создаем сервисы
    authService := services.NewAuthService(services.JWTConfig{
        SecretKey:     cfg.JWTSecret,
        AccessExpire:  cfg.JWTAccessExpire,
        RefreshExpire: cfg.JWTRefreshExpire,
    })

    // Создаем хендлеры
    authHandler := handlers.NewAuthHandler(authService, userRepo)
    userHandler := handlers.NewUserHandler(userRepo)
    cardHandler := handlers.NewCardHandler(cardRepo)
    terminalHandler := handlers.NewTerminalHandler(terminalRepo)
    keyHandler := handlers.NewKeyHandler(keyRepo)
    transactionHandler := handlers.NewTransactionHandler(transactionRepo)
    terminalAPIHandler := handlers.NewTerminalAPIHandler(
        cardRepo,
        terminalRepo,
        transactionRepo,
        keyRepo,
        cfg.TerminalToken,
    )

    // Настраиваем роутер
    router := setupRouter(
        authService,
        authHandler,
        userHandler,
        cardHandler,
        terminalHandler,
        keyHandler,
        transactionHandler,
        terminalAPIHandler,
    )

    // Запускаем сервер (без TLS, т.к. перед ним будет Nginx)
    addr := fmt.Sprintf(":%d", cfg.ServerPort)
    log.Printf("Starting server on %s", addr)
    
    if err := router.Run(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

func setupRouter(
    authService *services.AuthService,
    authHandler *handlers.AuthHandler,
    userHandler *handlers.UserHandler,
    cardHandler *handlers.CardHandler,
    terminalHandler *handlers.TerminalHandler,
    keyHandler *handlers.KeyHandler,
    transactionHandler *handlers.TransactionHandler,
    terminalAPIHandler *handlers.TerminalAPIHandler,
) *gin.Engine {
    gin.SetMode(gin.ReleaseMode)
    router := gin.New()
    router.Use(gin.Recovery())

    // Swagger UI
    swagger.Init()
    router.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Health check
    router.GET("/api/v1/health", healthCheck)

    // === Public routes ===
    auth := router.Group("/api/v1/auth")
    {
        auth.POST("/login", authHandler.Login)
        auth.POST("/register", authHandler.Register)
    }

    // === Terminal API (отдельная аутентификация) ===
    terminal := router.Group("/api/v1/terminal")
    {
        terminal.POST("/auth", terminalAPIHandler.AuthorizeTransaction)
        terminal.GET("/keys", terminalAPIHandler.DownloadKeys)
    }

    // === Protected routes (требуют JWT) ===
    api := router.Group("/api/v1")
    api.Use(middleware.AuthMiddleware(authService))
    {
        // Users CRUD
        users := api.Group("/users")
        {
            users.GET("", userHandler.GetAllUsers)
            users.GET("/:id", userHandler.GetUserByID)
            users.POST("", middleware.AdminMiddleware(), userHandler.CreateUser)
            users.PUT("/:id", userHandler.UpdateUser)
            users.DELETE("/:id", middleware.AdminMiddleware(), userHandler.DeleteUser)
        }

        // Cards CRUD
        cards := api.Group("/cards")
        {
            cards.GET("", cardHandler.GetAllCards)
            cards.GET("/:id", cardHandler.GetCardByID)
            cards.GET("/number/:number", cardHandler.GetCardByNumber)
            cards.POST("", cardHandler.CreateCard)
            cards.PUT("/:id", cardHandler.UpdateCard)
            cards.DELETE("/:id", cardHandler.DeleteCard)
        }

        // Terminals CRUD
        terminals := api.Group("/terminals")
        {
            terminals.GET("", terminalHandler.GetAllTerminals)
            terminals.GET("/:id", terminalHandler.GetTerminalByID)
            terminals.POST("", terminalHandler.CreateTerminal)
            terminals.PUT("/:id", terminalHandler.UpdateTerminal)
            terminals.DELETE("/:id", terminalHandler.DeleteTerminal)
        }

        // Keys CRUD (admin only)
        keys := api.Group("/keys")
        keys.Use(middleware.AdminMiddleware())
        {
            keys.GET("", keyHandler.GetAllKeys)
            keys.GET("/:id", keyHandler.GetKeyByID)
            keys.POST("", keyHandler.CreateKey)
            keys.PUT("/:id", keyHandler.UpdateKey)
            keys.DELETE("/:id", keyHandler.DeleteKey)
        }

        // Transactions (read only)
        transactions := api.Group("/transactions")
        {
            transactions.GET("/:id", transactionHandler.GetTransactionByID)
            transactions.GET("/card/:card_id", transactionHandler.GetTransactionsByCard)
            transactions.PUT("/:id/status", transactionHandler.UpdateTransactionStatus)
        }
    }

    return router
}