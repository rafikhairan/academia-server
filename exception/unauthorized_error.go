package exception

type UnauthorizedError struct {
	Code    int
	Status  string
	Message string
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Code:    401,
		Status:  "UNAUTHORIZED",
		Message: message,
	}
}

func (err UnauthorizedError) Error() string {
	return err.Message
}
