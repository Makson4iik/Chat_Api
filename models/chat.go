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
	if i.Chatname == "" && i.Creator == "" {
		return fmt.Errorf("chatname and creator is a required field")
	} else if i.Chatname == "" {
		return fmt.Errorf("chatname is a required field")
	} else if i.Creator == "" {
		return fmt.Errorf("creator is a required field")
	}
	return nil
}
func (*ChatList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Chat) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
