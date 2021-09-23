package hahamut

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Bot struct {
	botID       string
	accessToken string
	appSecret   string
}

const (
	sendMessageAPI = "https://us-central1-hahamut-8888.cloudfunctions.net/messagePush?bot_id=%s&access_token=%s"
	uploadImageAPI = "https://us-central1-hahamut-8888.cloudfunctions.net/imagePush?bot_id=%s&access_token=%s"
)

func NewBot(botID, accessToken, appSecret string) *Bot {
	return &Bot{
		botID:       botID,
		accessToken: accessToken,
		appSecret:   appSecret,
	}
}

func (b *Bot) SendText(receiver, content string) (string, error) {
	msg := &message{
		Recipient: recipient{
			ID: receiver,
		},
		Message: &textMessage{
			Type: MessageTypeText,
			Text: content,
		},
	}
	resp, err := b.sendMessage(msg)
	if resp != ResponseGetData {
		return "", fmt.Errorf(err.Error())
	}
	return resp, nil
}

func (b *Bot) SendImage(receiver string, image *Image) (string, error) {
	msg := &message{
		Recipient: recipient{
			ID: receiver,
		},
		Message: &imageMessage{
			Type:      MessageTypeImage,
			ID:        image.ID,
			Extension: image.Extension,
			Width:     strconv.Itoa(image.Width),
			Height:    strconv.Itoa(image.Height),
		},
	}
	resp, err := b.sendMessage(msg)
	if resp != ResponseGetData {
		return "", fmt.Errorf(err.Error())
	}
	return resp, nil
}

func (b *Bot) SendSticker(receiver, stickerGroup, stickerID string) (string, error) {
	msg := &message{
		Recipient: recipient{
			ID: receiver,
		},
		Message: &stickerMessage{
			Type:         MessageTypeSticker,
			StickerGroup: stickerGroup,
			StickerID:    stickerID,
		},
	}
	resp, err := b.sendMessage(msg)
	if resp != ResponseGetData {
		return "", fmt.Errorf(err.Error())
	}
	return resp, nil
}

func (b *Bot) sendMessage(body interface{}) (string, error) {
	url := fmt.Sprintf(sendMessageAPI, b.botID, b.accessToken)
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	return doPostRequest(url, jsonBytes)
}

func (b *Bot) StartEvent(receiver, startImg string, content *EventContent) (string, error) {
	msg := &message{
		Recipient: recipient{
			ID: receiver,
		},
		Message: startEventMessage{
			Type:     MessageTypeStartEvent,
			StartImg: startImg,
			Init: eventInit{
				Image:  content.Image,
				HP:     content.HP,
				Text:   content.Text,
				Button: content.Button,
			},
		},
	}
	return b.sendMessage(msg)
}

func (b *Bot) AddEvent(receiver, eventID string, content *EventContent) (string, error) {
	msg := &message{
		Recipient: recipient{
			ID: receiver,
		},
		Message: addEventMessage{
			Type:    MessageTypeAddEvent,
			EventID: eventID,
			Image:   content.Image,
			HP:      content.HP,
			Text:    content.Text,
			Button:  content.Button,
		},
	}
	return b.sendMessage(msg)
}

func (b *Bot) UploadImageFromLocal(path string) (*Image, error) {

	// open file
	file, err := os.Open(path)
	if err != nil {
		return &Image{}, err
	}
	defer file.Close()

	// do request
	url := fmt.Sprintf(uploadImageAPI, b.botID, b.accessToken)
	respBytes, err := doFileUploadRequest(url, path, file)
	if err != nil {
		return &Image{}, err
	}

	// check if ID exists
	var resp Image
	json.Unmarshal(respBytes, &resp)
	if resp.ID == "" {
		return &Image{}, fmt.Errorf(string(respBytes))
	}
	return &Image{
		ID:        resp.ID,
		Extension: resp.Extension,
		Width:     resp.Width,
		Height:    resp.Height,
	}, nil
}

func (b *Bot) UploadImageFromURL(imageURL string) (*Image, error) {
	// get image
	res, err := http.Get(imageURL)
	if err != nil {
		return &Image{}, err
	}
	defer res.Body.Close()

	// do request
	url := fmt.Sprintf(uploadImageAPI, b.botID, b.accessToken)
	respBytes, err := doFileUploadRequest(url, imageURL, res.Body)
	if err != nil {
		return &Image{}, err
	}

	// check if ID exists
	var resp Image
	json.Unmarshal(respBytes, &resp)
	if resp.ID == "" {
		return &Image{}, fmt.Errorf(string(respBytes))
	}
	return &Image{
		ID:        resp.ID,
		Extension: resp.Extension,
		Width:     resp.Width,
		Height:    resp.Height,
	}, nil
}
