package utils

type ResponseMessage struct {
	Status  interface{} `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
