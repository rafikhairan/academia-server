package exception

type BadRequestError struct {
	Code    int
	Status  string
	Message string
}

func (err BadRequestError) Error() string {
	return err.Message
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Code:    400,
		Status:  "BAD_REQUEST",
		Message: message,
	}
}
