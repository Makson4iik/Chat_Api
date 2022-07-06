package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/Makson4iik/Chat_Api/db"
	"github.com/Makson4iik/Chat_Api/models"
)

//var itemIDKey = "itemID"

func ChatApi(router chi.Router) {
	router.Post("/createchat", createChat)
	router.Post("/addmessage", addMessage)
	router.Get("/getallmessagesbychat", GetAllMessagesByChat)
	router.Get("/getmessagesnyid", GetMessagesById)
}

func createChat(w http.ResponseWriter, r *http.Request) {
	chat := &models.Chat{}
	if err := render.Bind(r, chat); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddChat(chat); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, chat); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	message := &models.Message{}
	if err := render.Bind(r, message); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddMessage(message); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, message); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func GetAllMessagesByChat(w http.ResponseWriter, r *http.Request) {
	chatNameStr := r.URL.Query().Get("chatname")
	chatDepthStr := r.URL.Query().Get("chatdepth")
	chatDepth, err := strconv.Atoi(chatDepthStr)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	messages, err := dbInstance.GetAllMessagesByChat(chatNameStr, chatDepth)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, messages); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func GetMessagesById(w http.ResponseWriter, r *http.Request) {
	messIdStr := r.URL.Query().Get("messageid")
	messId, err := strconv.Atoi(messIdStr)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	message, err := dbInstance.GetMessageById(messId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &message); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
