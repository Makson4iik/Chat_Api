package models

import (
	"fmt"
	"net/http"
)

type Chat struct {
	Chatname string `json:"chatname"`
	Creator  string `json:"creator"`
}
type ChatList struct {
	Chats []Chat `json:"chats"`
}

func (i *Chat) Bind(r *http.Request) error {
	if i.Chatname == "" {
		return fmt.Errorf("Chatname is a required field")
	}
	return nil
}
func (*ChatList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Chat) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
