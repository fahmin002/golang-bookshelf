package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fahminbl/bookshelf/internal/handler"
	"github.com/fahminbl/bookshelf/internal/repository"
	"github.com/fahminbl/bookshelf/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "fahmi:1234@tcp(127.0.0.1:3306)/bookshelf")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookRepo := &repository.BookRepository{DB: db}
	bookService := &service.BookService{Repo: bookRepo}
	bookHandler := &handler.BookHandler{Service: bookService}

	r := mux.NewRouter()
	r.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
