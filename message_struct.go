package hahamut

type message struct {
	Recipient recipient   `json:"recipient"`
	Message   interface{} `json:"message"`
}
type recipient struct {
	ID string `json:"id"`
}

type textMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type stickerMessage struct {
	Type         string `json:"type"`
	StickerGroup string `json:"sticker_group"`
	StickerID    string `json:"sticker_id"`
}

type imageMessage struct {
	Type      string `json:"type"`
	ID        string `json:"id"`
	Extension string `json:"ext"`
	Width     string `json:"width"`
	Height    string `json:"height"`
}

type Image struct {
	ID        string `json:"id"`
	Extension string `json:"ext"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}
