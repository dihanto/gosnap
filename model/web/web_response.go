package web

type WebResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
