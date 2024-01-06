package bookshelf

import "github.com/emanueltimlopez/patmos/internal/book"

type UserBooksGet struct {
	UserID string    `json:"user_id"`
	BookID string    `json:"book_id"`
	Book   book.Book `json:"books,omitempty"`
	Folder string    `json:"folder"`
}
