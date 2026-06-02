package payload

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}
