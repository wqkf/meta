package errors

type ErrorResponse struct {
	Code    int
	Message string
}

func SendError(code int, message string) ErrorResponse {
	return ErrorResponse{code, message}
}
