package models

// APIResponse - универсальный ответ API
type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

// ErrorResponse - ответ с ошибкой
type ErrorResponse struct {
    Error string `json:"error"`
}

func SuccessResponse(data interface{}) APIResponse {
    return APIResponse{
        Success: true,
        Data:    data,
    }
}

func ErrorResponse(message string) APIResponse {
    return APIResponse{
        Success: false,
        Error:   message,
    }
}