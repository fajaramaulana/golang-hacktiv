package views

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status" example:"200"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
