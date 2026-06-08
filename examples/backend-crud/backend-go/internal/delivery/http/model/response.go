package model

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}
