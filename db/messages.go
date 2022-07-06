package db

import (
	"database/sql"

	"github.com/Makson4iik/Chat_Api/models"
)

func (db Database) GetAllMessagesByChat(chatName string, depth int) (*models.MessageList, error) {
	query := "SELECT * FROM messages WHERE chatname = $1 LIMIT $2"
	list := &models.MessageList{}
	rows, err := db.Conn.Query(query, chatName, depth)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var mess models.Message
		err := rows.Scan(&mess.MessageID, &mess.Chatname, &mess.Creator, &mess.MessText)
		if err != nil {
			return list, err
		}
		list.Messages = append(list.Messages, mess)
	}
	return list, nil
}

func (db Database) GetMessageById(messId int) (models.Message, error) {
	mess := models.Message{}
	query := "SELECT * FROM messages WHERE chatname = $1 LIMIT $2"
	row := db.Conn.QueryRow(query, messId)
	switch err := row.Scan(&mess.MessageID, &mess.Chatname, &mess.Creator, &mess.MessText); err {
	case sql.ErrNoRows:
		return mess, ErrNoMatch
	default:
		return mess, err
	}
}

func (db Database) AddMessage(mess *models.Message) error {
	var MessageID int
	query := `INSERT INTO messages (chatname, creator, messtext) VALUES ($1, $2, $3) RETURNING message_id`
	err := db.Conn.QueryRow(query, mess.Chatname, mess.Creator, mess.MessText).Scan(&MessageID)
	if err != nil {
		return err
	}
	mess.MessageID = MessageID
	return nil
}
