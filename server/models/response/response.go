package response

type BaseResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	RequestID  string      `json:"requestID"`
	Payload    interface{} `json:"payload"`
}
