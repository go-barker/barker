package client

type ErrorResponse struct {
	Message string `json:"error,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
