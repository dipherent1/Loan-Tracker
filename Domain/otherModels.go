package domain

type Response struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data ,omitempty"`
	AccessToken string      `json:"access_token ,omitempty"`
}

type Filter struct {
	StatusOrder []string `json:"status_order"`
	OrderBy	 string   `json:"order_by"`
}
