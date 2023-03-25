package response

type BaseResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	RequestID  string      `json:"requestID"`
	Payload    interface{} `json:"payload,omitempty"`
}

type APIFormError struct {
	Field        string `json:"field"`
	ErrorMessage string `json:"error_message"`
}

func (e APIFormError) Error() string { return e.ErrorMessage }
