package http

type APIResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type MetaInstagramMediaResponse struct {
	Id string `json:"id"`
}
