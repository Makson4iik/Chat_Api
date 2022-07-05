package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/Makson4iik/Chat_Api/db"
	"github.com/Makson4iik/Chat_Api/models"
)

//var itemIDKey = "itemID"

func ChatApi(router chi.Router) {
	router.Post("/", createChat)
	router.Post("/", addMessage)
	router.Get("/", GetAllMessagesByChat)
	router.Get("/", GetMessagesById)

	/*router.Route("/{itemId}", func(router chi.Router) {
		router.Use(ItemContext)
		router.Get("/", getItem)
		router.Put("/", updateItem)
		router.Delete("/", deleteItem)
	})*/
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

}

func ItemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "itemId")
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
			return
		}
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
		}
		ctx := context.WithValue(r.Context(), itemIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

/*func createItem(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddItem(item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}*/

func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetAllItems()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, items); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.Context().Value(itemIDKey).(int)
	item, err := dbInstance.GetItemById(itemID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
