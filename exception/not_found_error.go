package exception

type NotFoundError struct {
	Code    int
	Status  string
	Message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Code:    404,
		Status:  "NOT_FOUND",
		Message: message,
	}
}

func (err NotFoundError) Error() string {
	return err.Message
}
