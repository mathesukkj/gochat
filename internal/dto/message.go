package dto

import "time"

type Message struct {
	Message string
	SentAt  time.Time
	SentBy  WebsocketClient
}

func (m Message) ToString() string {
	fmtTime := m.SentAt.Format("15:04")
	return m.Message + "- \033[36m" + fmtTime + "\033[0m"
}
