# go-hahamut-bot
非官方的哈哈姆特不 EY 機器人套件

## 不同語言版本
* [English](https://github.com/upk1997/go-hahamut-bot/blob/master/README.md)
* [繁體中文](https://github.com/upk1997/go-hahamut-bot/blob/master/README_zh-tw.md)

## 事前準備
1. 到巴哈姆特的[創作者後台](https://haha.gamer.com.tw/bot_list.php)申請一個哈哈姆特的 BOT
2. 從該 BOT 的設定頁面取得它的 access token 和 secret key
3. 妥善保存它們，不要讓任何人取得這組 access token 或 secret key 的內容

## 開始使用
1. 在 cmd 或 powershell 用以下指令安裝本套件：
```shell
go get -u "github.com/upk1997/go-hahamut-bot"
```
2. 把這個套件 import 到你的專案裡：
```go
import "github.com/upk1997/go-hahamut-bot"
```
3. 產生一個 Bot 的物件：
```go
myBot := hahamut.NewBot(botID, botAccessToken, botSecretKey)
```

## 傳送訊息給巴友

### 文字訊息
```go
receiverID := "sega"
message := "Test message"
_, err := myBot.SendText(receiverID, message)
if err != nil {
    log.Fatalln(err)
}
```

### 貼圖訊息

```go
receiverID := "sega"
stickerGroup := "1"
stickerID := "08"
_, err := myBot.SendSticker(receiverID, stickerGroup, stickerID)
if err != nil {
    log.Fatalln(err)
}
```
你可以從這裡取得貼圖的 group ID 和貼圖本身的 ID：[可使用貼圖列表](https://haha.gamer.com.tw/bot_sticker_list.php)

### 圖片訊息

#### 從你的電腦直接上傳圖片
```go
// 從你的電腦上傳圖片
image, err := myBot.UploadImageFromLocal("D:/test.png")
if err != nil {
    log.Fatalln(err)
}
```
#### 從網址上傳圖片
```go
// 從網址上傳圖片
imageURL := "https://avatar2.bahamut.com.tw/avataruserpic/s/e/sega/sega.png"
image, err := myBot.UploadImageFromURL(imageURL)
if err != nil {
    log.Fatalln(err)
}
```
#### 傳送圖片
```go
// 傳送圖片
receiverID := "sega"
_, err = myBot.SendImage(receiverID, image)
if err != nil {
    log.Fatalln(err)
}
```

在傳送圖片之前，你必須先把圖片上傳到巴哈姆特的伺服器。上傳完畢之後，你會獲得一個有 ID（編號）、Extension（副檔名）、Width（寬度）、Height（高度）等屬性的 \*Image 物件。上傳好的圖片是可以拿來重複使用的，你可以在上傳好之後把整組的圖片資訊存下來，以利於下次重複使用。

### 啟動新的事件

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
* Entry image: 顯示在聊天訊息介面裡，讓使用者點入的圖片
* Main image: 從 entry image 點進去之後顯示的主要圖片

如上面提到的傳送圖片流程，這裡也是要先用 Upload 方法把圖片上傳到巴哈姆特之後才能拿來用，EventButtonSetting 裡面的 Disabled（凍結）和 Hidden（隱藏）屬性如果是 false 的話可以直接省略。

對特定巴友啟動了新的事件之後，你會得到一組 event ID（事件編號），這個事件編號可以用來讓你接續之後觸發的事件。這裡要注意，如果沒有成功啟動新事件的話，event ID 這個變數的內容會是錯誤訊息而不是真正的事件編號。

### 新增（更新）已啟動的現有事件

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
receiverID := "sega"
resp, err := myBot.AddEvent(receiverID, eventID, eventContent)
if err != nil {
	log.Fatalln(err)
}
fmt.Println(resp)
```
事件的內容定義方法跟啟動事件的時候一樣，如果成功更新了現有的事件，你會得到一個 **"event adding"**（hahamut.ResponseEventAdding）的字串。


### 捕捉錯誤

傳送訊息的時候，如果從巴哈姆特伺服器得到的回應不是站方定義的 **"get data\~\~"**（hahamut.ResponseGetData），傳送訊息的函數就會回傳一個 error。如果你想處理不同的錯誤內容，你可以在 response_type.go 檔案裡面看到一些常見的錯誤內容，然後像這樣做：

```go
if err != nil {
    switch err.Error() {
    case hahamut.ResponseInvalidAccessToken:
        fmt.Println("Invalid access token.")
    case hahamut.ResponseCannotBeInitiative:
        fmt.Printf("Cannot send messages: No conversation with %s before.\n", receiverID)
    default:
        // 其他錯誤類型
        // ...
    }
}
```

## 從巴哈姆特的伺服器接收訊息

在你的伺服器接收到巴哈姆特傳來的 request 之後，把請求的 body 內容拿去解析：

```go
func myWebhook(w http.ResponseWriter, r *http.Request) {
    // 取得連線請求的 body 內容
    bodyBytes, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatalln(err)
    }

    // 取得連線請求附帶的簽章
    signature := r.Header.Get("x-baha-data-signature")

    // 解析 body 內容，同時檢查簽章是否有效
    event, err := myBot.ParseWebhookEventBody(bodyBytes, signature)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
}
```

為了確保安全性，ParseWebhookEventBody 函數會強制檢查連線請求裡的簽章是不是正確的，如果連線請求附帶的簽章和實際計算出來的簽章不同的話，函數就會回傳一個 error。

## HTTP 伺服器 & Webhook 範例

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
	// 設定 webhook，建置 HTTP 伺服器
	http.HandleFunc("/hahamutBot", myWebhook)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func myWebhook(w http.ResponseWriter, r *http.Request) {
	// 取得連線請求的 body 內容
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// 取得連線請求附帶的簽章
	signature := r.Header.Get("x-baha-data-signature")

	// 解析 body 內容，同時檢查簽章是否有效
	event, err := myBot.ParseWebhookEventBody(bodyBytes, signature)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 取得訊息內容
	for _, m := range event.Messaging {
		if m.Message.Text != "" {
			// 得到文字訊息
			handleTextMessage(event.Time, m.SenderID, m.Message.Text)
		} else if m.Message.Sticker.ID != "" {
			// 得到貼圖訊息
			handleStickerMessage(event.Time, m.SenderID, m.Message.Sticker.Group, m.Message.Sticker.ID)
		} else if m.Message.BotCommand != "" {
			// 得到機器人指令訊息
			handleCommandMessage(event.Time, m.SenderID, m.Message.BotCommand, m.Message.EventID)
		} else {
			// 未知的類型
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
	// 這裡可以用來處理按了 event 的按鈕之後接續的事件
	// ...
}
```

把 BOT_ID, BOT_ACCESS_TOKEN, BOT_SECRET_KEY 設定到 Heroku APP 裡的 config vars 之後，你就可以把你的專案部署上去，接著再把 webhook 放到你的哈哈姆特 BOT 設定上面就可以了。

照著上面的範例來做，你的 webhook 網址會像這樣:
```
https://your-app-name.herokuapp.com/hahamutBot
```

## 提醒

* 接收訊息的功能需要在 HTTPS 環境下運作，也就是說你必須要有一個像 Heroku 這樣的雲端伺服器才能測試它。
