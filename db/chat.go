package db

import (
	"github.com/Makson4iik/Chat_Api/models"
)

func (db Database) GetAllChats() (*models.ChatList, error) {
	list := &models.ChatList{}
	rows, err := db.Conn.Query("SELECT * FROM chat")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(&chat.Chatname, &chat.Creator)
		if err != nil {
			return list, err
		}
		list.Chats = append(list.Chats, chat)
	}
	return list, nil
}

func (db Database) AddChat(addChat *models.Chat) error {

	chat := models.Chat{}
	query := `INSERT INTO chat (chatname, creator) VALUES ($1, $2)`
	err := db.Conn.QueryRow(query, addChat.Chatname, addChat.Creator).Scan(&chat.Chatname, &chat.Creator)
	if err != nil {
		return err
	}
	return nil
}
