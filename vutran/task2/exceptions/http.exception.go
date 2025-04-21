package exceptions

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *HttpError) Error() string {
	return err.Message
}