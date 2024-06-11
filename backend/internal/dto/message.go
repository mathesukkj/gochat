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

func (m Message) ToString() string {
	fmtTime := m.SentAt.Format("15:04")
	return fmt.Sprintf(
		"\033[36m%s:\033[0m %s - \033[36m%s\033[0m",
		strings.TrimSpace(m.SentBy.User),
		m.Message,
		fmtTime,
	)
}
