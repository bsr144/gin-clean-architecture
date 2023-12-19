package response

type HTTPResponseSuccess struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type HTTPResponseSuccessWithPagination struct {
	Code       int         `json:"code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination"`
}

func NewHTTPResponseSuccess(code int, message string, data interface{}) *HTTPResponseSuccess {
	return &HTTPResponseSuccess{
		Code:    code,
		Success: true,
		Message: message,
		Data:    data,
	}
}

func NewHTTPResponseSuccessWithPagination(code int, message string, data interface{}, pagination *PaginationResponse) *HTTPResponseSuccessWithPagination {
	return &HTTPResponseSuccessWithPagination{
		Code:       code,
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}
