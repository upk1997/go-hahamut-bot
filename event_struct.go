package hahamut

type startEventMessage struct {
	Type     string    `json:"type"`
	StartImg string    `json:"start_img"`
	Init     eventInit `json:"init"`
}

type addEventMessage struct {
	Type    string      `json:"type"`
	EventID string      `json:"event_id"`
	Image   string      `json:"image"`
	HP      EventHP     `json:"hp"`
	Text    EventText   `json:"text"`
	Button  EventButton `json:"button"`
}

type eventInit struct {
	Image  string      `json:"image"`
	HP     EventHP     `json:"hp"`
	Text   EventText   `json:"text"`
	Button EventButton `json:"button"`
}

type EventContent struct {
	Image       string
	HP          EventHP
	Text        EventText
	ButtonStyle int
	Button      EventButton
}

type EventHP struct {
	Hidden  bool   `json:"hidden"`
	Max     int    `json:"max"`
	Current int    `json:"current"`
	Color   string `json:"color"`
}

type EventText struct {
	Hidden  bool   `json:"hidden"`
	Message string `json:"message"`
	Color   string `json:"color"`
}

type EventButton struct {
	Style   int                  `json:"style"`
	Setting []EventButtonSetting `json:"setting"`
}

type EventButtonSetting struct {
	Hidden   bool   `json:"hidden"`
	Disabled bool   `json:"disabled"`
	Order    int    `json:"order"`
	Text     string `json:"text"`
	Command  string `json:"command"`
}
