package helpers

type HttpError struct {
	Message    string
	StatusCode int
}


func (e *HttpError) Error() string {
	return e.Message
}

func New(msg string, statusCode int) *HttpError{
	return &HttpError{
		Message:    msg,
		StatusCode: statusCode,
	}
}