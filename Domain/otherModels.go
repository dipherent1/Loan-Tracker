package domain

type Respose struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data ,omitempty"`
}