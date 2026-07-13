package http

type MetaInstagramMediaRequest struct {
	ImageUrl string `json:"image_url"`
	Caption  string `json:"caption"`
}

type MetaInstagramSendMessageModel struct {
	Recipient struct {
		Id string `json:"id"`
	} `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}
