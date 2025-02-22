package handlers

import (
	"gotth/internal/middleware"
	"gotth/internal/store"
	"gotth/internal/templates"
	"net/http"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.UserKey).(*store.User)
	if !ok {
		c := templates.GuestIndex()
		err := templates.Layout(c, "My website").Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	c := templates.Index(user.Email)
	err := templates.Layout(c, "My website").Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
