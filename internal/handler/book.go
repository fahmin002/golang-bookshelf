package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fahminbl/bookshelf/internal/service"
)

type BookHandler struct {
	Service *service.BookService
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
