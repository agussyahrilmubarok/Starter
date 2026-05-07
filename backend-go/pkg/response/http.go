package response

type HttpSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type HttpError struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func HttpErrMap(key, err string) map[string]string {
	return map[string]string{
		key: err,
	}
}
