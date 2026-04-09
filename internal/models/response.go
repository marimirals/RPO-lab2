package models

// APIResponse - универсальный ответ API
type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

// ErrorResponse - структура для ответа с ошибкой
type ErrorResponse struct {
    Error string `json:"error"`
}

// SuccessResponse создаёт успешный ответ
func SuccessResponse(data interface{}) APIResponse {
    return APIResponse{
        Success: true,
        Data:    data,
    }
}

// NewErrorResponse создаёт ответ с ошибкой (переименовали функцию!)
func NewErrorResponse(message string) APIResponse {
    return APIResponse{
        Success: false,
        Error:   message,
    }
}