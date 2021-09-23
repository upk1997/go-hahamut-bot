package hahamut

import "time"

type WebhookEvent struct {
	BotID     string `json:"botid"`
	TimeStamp int64  `json:"time"`
	Time      time.Time
	Messaging []struct {
		SenderID string `json:"sender_id"`
		Message  struct {
			BotCommand string `json:"bot_command"`
			EventID    string `json:"event_id"`
			Text       string `json:"text"`
			Sticker    struct {
				Group string `json:"group"`
				ID    string `json:"id"`
			} `json:"sticker"`
		} `json:"message"`
	} `json:"messaging"`
}
