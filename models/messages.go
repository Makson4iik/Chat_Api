package models

import (
	"fmt"
	"net/http"
)

type Message struct {
	MessageID int    `json:"message_id"`
	Chatname  string `json:"chatname"`
	Creator   string `json:"creator"`
	MessText  string `json:"mess_text"`
}
type MessageList struct {
	Messages []Message `json:"messages"`
}

func (i *Message) Bind(r *http.Request) error {
	if i.Chatname == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}
func (*MessageList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Message) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
