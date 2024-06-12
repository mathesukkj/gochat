package dto

import (
	"fmt"
	"strings"
	"time"
)

type Message struct {
	Message string
	SentAt  time.Time
	SentBy  WebsocketClient
}

func (m Message) GetSenderUsername() string {
	return strings.TrimSpace(m.SentBy.User)
}

func (m Message) ToString() string {
	fmtTime := m.SentAt.Format("15:04")
	return fmt.Sprintf(
		"\033[36m%s:\033[0m %s - \033[36m%s\033[0m",
		m.GetSenderUsername(),
		m.Message,
		fmtTime,
	)
}

func (m Message) ToJson() string {
	return fmt.Sprintf(
		`{
      message: %s,
      sentBy: %s,
      sentAt: %v,
    }`,
		m.Message,
		m.GetSenderUsername(),
		m.SentAt,
	)
}
