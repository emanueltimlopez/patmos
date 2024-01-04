package bookshelf

import "github.com/emanueltimlopez/books-motivation/internal/book"

type UserBooksSave struct {
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
}

type UserBooksGet struct {
	UserID string    `json:"user_id"`
	BookID string    `json:"book_id"`
	Book   book.Book `json:"books,omitempty"`
}
