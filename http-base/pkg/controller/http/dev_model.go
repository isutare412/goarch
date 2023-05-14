package http

type handlePostRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int32  `json:"age"`
}
