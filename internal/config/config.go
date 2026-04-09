package config

import (
    "os"
    "strconv"
    "time"
)

type Config struct {
    ServerPort      int
    DatabasePath    string
    JWTSecret       string
    JWTAccessExpire time.Duration
    JWTRefreshExpire time.Duration
    TerminalToken   string
    Env             string
}

func Load() *Config {
    port, _ := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
    
    accessExpire, _ := time.ParseDuration(getEnv("JWT_ACCESS_EXPIRE", "1h"))
    refreshExpire, _ := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRE", "168h"))

    return &Config{
        ServerPort:      port,
        DatabasePath:    getEnv("DB_PATH", "./data/lab2.db"),
        JWTSecret:       getEnv("JWT_SECRET", "change-me-in-production"),
        JWTAccessExpire: accessExpire,
        JWTRefreshExpire: refreshExpire,
        TerminalToken:   getEnv("TERMINAL_TOKEN", "terminal-secret-token"),
        Env:             getEnv("ENV", "development"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}