# go-hahamut-bot
An unofficial SDK of hahamut bot.

## Before Using go-hahamut-bot
1. Apply a Hahamut bot on [Bahamut](https://haha.gamer.com.tw/bot_list.php).
2. Get your access token & secret key from the bot detail page.
3. Keep them safe. Don't let anyone know access token or secret key of Hahamut bot.

## Start To Use
1. Get this package from terminal
```shell
go get -u "github.com/upk1997/go-hahamut-bot"
```
2. Create a new bot
```go
myBot := hahamut.NewBot(botID, botAccessToken, botSecretKey)
```
3. Import the package
```go
import "github.com/upk1997/go-hahamut-bot"
```


## Sending Messages To Someone

### Text
```go
receiverID := "sega"
message := "Test message"
_, err := myBot.SendText(receiverID, message)
if err != nil {
    log.Fatalln(err)
}
```

### Sticker

```go
receiverID := "sega"
stickerGroup := "1"
stickerID := "08"
_, err := myBot.SendSticker(receiverID, stickerGroup, stickerID)
if err != nil {
    log.Fatalln(err)
}
```
More stickers information about group & ID: [Stickers list](https://haha.gamer.com.tw/bot_sticker_list.php)

### Image

#### Uploading local image file
```go
// upload image from local
image, err := myBot.UploadImageFromLocal("D:/test.png")
if err != nil {
    log.Fatalln(err)
}
```
#### Uploading & image file from URL
```go
// upload image from URL
imageURL := "https://avatar2.bahamut.com.tw/avataruserpic/s/e/sega/sega.png"
image, err := myBot.UploadImageFromURL(imageURL)
if err != nil {
    log.Fatalln(err)
}
```
#### Sending image
```go
// send image
receiverID := "sega"
_, err = myBot.SendImage(receiverID, image)
if err != nil {
    log.Fatalln(err)
}
```

You must upload image to Bahamut's server before sending it. After uploaded, it returns an *Image containing properties of ID, Extension, Width, Height.

The images uploaded to Bahamut server are reusable. You can save these properties to reuse them if you need.

### Start A New Event

```go
mainImageFilename := fmt.Sprintf("%s.%s", image.ID, image.Extension)
eventContent := &hahamut.EventContent{
    Image: mainImageFilename,
    HP: hahamut.EventHP{
        Max:     200,
        Current: 200,
        Color:   "#FFFF00",
    },
    Text: hahamut.EventText{
        Message: "Main message",
        Color:   "#FF00FF",
    },
    Button: hahamut.EventButton{
        Style: 2,
        Setting: []hahamut.EventButtonSetting{
            {
                Disabled: false,
                Hidden:   false,
                Order:    1,
                Text:     "Option 1 message",
                Command:  "/bot option 1",
            },
            {
                Disabled: false,
                Hidden:   false,
                Order:    2,
                Text:     "Option 2 message",
                Command:  "/bot option 2",
            },
            {
                Disabled: false,
                Hidden:   false,
                Order:    3,
                Text:     "Option 3 message",
                Command:  "/bot option 3",
            },
            {
                Disabled: false,
                Hidden:   false,
                Order:    4,
                Text:     "Option 4 message",
                Command:  "/bot option 4",
            },
        },
    },
}
receiverID := "sega"
entryImageFilename := fmt.Sprintf("%s.%s", image.ID, image.Extension)
eventID, err := myBot.StartEvent(receiverID, entryImageFilename, eventContent)
if err != nil {
    fmt.Println("Error:", err.Error())
}
```
* Entry image: displayed outside to click
* Main image: displayed in your event interface

Same as the function of sending images, images must be uploaded to server before using. The propeties of Disabled and Hidden in the button setting can be omitted if they are false.

If successfully started a new event, you will get an event ID as response. The event ID is used for adding new event.

Beware of the content of eventID variable, it will be error message instead of event ID if an error occurs.

### Adding (Refreshing) An Exising Event

```go
mainImageFilename := fmt.Sprintf("%s.%s", image.ID, image.Extension)
eventContent := &hahamut.EventContent{
    Image: mainImageFilename,
    HP: hahamut.EventHP{
        Max:     200,
        Current: 150,
        Color:   "#999900",
    },
    Text: hahamut.EventText{
        Message: "Main message (new!)",
        Color:   "#990099",
    },
    Button: hahamut.EventButton{
        Style: 2,
        Setting: []hahamut.EventButtonSetting{
            {
                Disabled: true,
                Order:    1,
                Text:     "Option 1 message ++",
                Command:  "/bot option 1 2",
            },
            {
                Order:    2,
                Text:     "Option 2 message ++",
                Command:  "/bot option 2 2",
            },
            {
                Disabled: true,
                Order:    3,
                Text:     "Option 3 message ++",
                Command:  "/bot option 3 2",
            },
            {
                Order:    4,
                Text:     "Option 4 message ++",
                Command:  "/bot option 4 2",
            },
        },
    },
}

```
The definition of event content is same as starting an event. Get an **"event adding"** (hahamut.ResponseEventAdding) as response if successfully refreshed the event.


### Catching Errors

When sending messages, it also returns an error if the response content isn't **"get data\~\~"** (hahamut.ResponseGetData) defined by Bahamut developers. If you want to handle different response types, you can find them from response_type.go. Like this:

```go
if err != nil {
    switch err.Error() {
    case hahamut.ResponseInvalidAccessToken:
        fmt.Println("Invalid access token.")
    case hahamut.ResponseCannotBeInitiative:
        fmt.Printf("Cannot send messages: No conversation with %s before.\n", receiverID)
    default:
        // other error types
        // ...
    }
}
```

## Receiving Messages from Bahamut Server

After your server got a request, parse the content of body in your handler:

```go
func myWebhook(w http.ResponseWriter, r *http.Request) {
    // get body content
    bodyBytes, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatalln(err)
    }

    // get the signature from request header
    signature := r.Header.Get("x-baha-data-signature")

    // parse event body & check signature
    event, err := myBot.ParseWebhookEventBody(bodyBytes, signature)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
}
```

The ParseWebhookEventBody function will check the signature forcibly for security. It returns an error if the signature header of request is different from the signature calculated by the function.

## HTTP Server & Webhook Example

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/upk1997/go-hahamut-bot"
)

var myBot *hahamut.Bot

func init() {
	botID := os.Getenv("BOT_ID")
	botAccessToken := os.Getenv("BOT_ACCESS_TOKEN")
	botSecretKey := os.Getenv("BOT_SECRET_KEY")
	myBot = hahamut.NewBot(botID, botAccessToken, botSecretKey)
}

func main() {
	// set router & serve
	http.HandleFunc("/hahamutBot", myWebhook)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func myWebhook(w http.ResponseWriter, r *http.Request) {
	// get body content
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// get the signature from request header
	signature := r.Header.Get("x-baha-data-signature")

	// parse event body & check signature
	event, err := myBot.ParseWebhookEventBody(bodyBytes, signature)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// get messages
	for _, m := range event.Messaging {
		if m.Message.Text != "" {
			// received a text message
			handleTextMessage(event.Time, m.SenderID, m.Message.Text)
		} else if m.Message.Sticker.ID != "" {
			// received a sticker message
			handleStickerMessage(event.Time, m.SenderID, m.Message.Sticker.Group, m.Message.Sticker.ID)
		} else if m.Message.BotCommand != "" {
			// received a bot command message
			handleCommandMessage(event.Time, m.SenderID, m.Message.BotCommand, m.Message.EventID)
		} else {
			// unknown type
			fmt.Printf("[%v] Received an unknown type message\n", event.Time)
			fmt.Println(string(bodyBytes))
		}
	}
}

func handleTextMessage(t time.Time, id, message string) {
	fmt.Printf("[%v] Received a text message:\nSenderID: %s\nMessage: %s\n", t, id, message)
	myBot.SendText(id, "Nice!")
}

func handleStickerMessage(t time.Time, id, stickerGroup, stickerID string) {
	fmt.Printf("[%v] Received a sticker:\nSenderID: %s\nSticker group: %s\nSticker ID: %s\n", t, id, stickerGroup, stickerID)
	myBot.SendSticker(id, "11", "03")
}

func handleCommandMessage(t time.Time, id, command, eventID string) {
	fmt.Printf("[%v] Received a command:\nSenderID: %s\nCommand: %s\nEvent ID: %s\n", t, id, command, eventID)
	// then you can do something to continue events after clicking a button
	// ...
}
```

After setting BOT_ID, BOT_ACCESS_TOKEN, BOT_SECRET_KEY of config vars manually, you can deploy it on Heroku and set webhook in your Hahamut bot config page.

As the sample above, the webhook url will be like this:
```
https://your-app-name.herokuapp.com/hahamutBot
```

## Tips

* The function of receiving messages (webhook) should run in HTTPS environment. In other words, you need a cloud server like Heroku to test it.
