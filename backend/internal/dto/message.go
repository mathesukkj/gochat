package dto

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Message struct {
	Message string
	SentAt  time.Time
	SentBy  *WebsocketClient
}

func (m Message) GetSenderUsername() string {
	return strings.TrimSpace(m.SentBy.User)
}

func (m Message) ToString(messageType string) string {
	if messageType == "json" {
		return m.toJson()
	}
	return m.toText()
}

func (m Message) toText() string {
	fmtTime := m.SentAt.Format("15:04")
	return fmt.Sprintf(
		"\033[36m%s:\033[0m %s - \033[36m%s\033[0m",
		m.GetSenderUsername(),
		m.Message,
		fmtTime,
	)
}

func (m Message) toJson() string {
	messageStruct := struct {
		Message string `json:"message"`
		SentBy  string `json:"sentBy"`
		SentAt  string
	}{
		m.Message,
		m.GetSenderUsername(),
		m.SentAt.String(),
	}
	messagePayload, _ := json.Marshal(messageStruct)

	return string(messagePayload)
}
