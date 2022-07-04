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

/*func (db Database) GetItemById(itemId int) (models.Chat, error) {
	chat := models.Chat{}
	query := `SELECT * FROM chat WHERE chatname = $1;`
	row := db.Conn.QueryRow(query, itemId)
	switch err := row.Scan(&chat.Chatname, &chat.Creator,); err {
	case sql.ErrNoRows:
		return chat, ErrNoMatch
	default:
		return chat, err
	}
}

func (db Database) DeleteItem(itemId int) error {
	query := `DELETE FROM items WHERE id = $1;`
	_, err := db.Conn.Exec(query, itemId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
	item := models.Item{}
	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
	err := db.Conn.QueryRow(query, itemData.Name, itemData.Description, itemId).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, ErrNoMatch
		}
		return item, err
	}
	return item, nil
}*/
