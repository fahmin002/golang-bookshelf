package repository

import (
	"database/sql"

	"github.com/fahminbl/bookshelf/internal/model"
)

type BookRepository struct {
	DB *sql.DB
}

func (r *BookRepository) GetAll() ([]model.Book, error) {
	rows, err := r.DB.Query("SELECT id, title, author, description, cover_image FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []model.Book{}
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.CoverImage)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
