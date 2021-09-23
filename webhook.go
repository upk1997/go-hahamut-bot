package hahamut

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"
)

func (b *Bot) ParseWebhookEventBody(body []byte, signature string) (*WebhookEvent, error) {
	webhookEvent := &WebhookEvent{}

	// check signature
	if !b.isValidSignature(body, signature) {
		// For security, it will return &WebhookEvent{} and a ErrorInvalidSignature if signature is invalid
		return webhookEvent, fmt.Errorf(ErrorInvalidSignature)
	}

	// parse the body
	err := json.Unmarshal(body, &webhookEvent)
	if err != nil {
		return webhookEvent, err
	}

	// parse timestamp
	webhookEvent.Time = time.Unix(webhookEvent.TimeStamp/1000, 0)

	return webhookEvent, err
}

func (b *Bot) isValidSignature(body []byte, signature string) bool {
	mac := hmac.New(sha1.New, []byte(b.appSecret))
	mac.Write(body)
	return fmt.Sprintf("sha1=%x", mac.Sum(nil)) == signature
}
