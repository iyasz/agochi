package helpers

type HttpError struct {
	Field      string
	Message    string
	StatusCode int
}

func (e *HttpError) Error() string {
	return e.Message
}

func New(field string, msg string, statusCode int) *HttpError {
	return &HttpError{
		Field: field,
		Message:    msg,
		StatusCode: statusCode,
	}
}
