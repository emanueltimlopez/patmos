package shelf

import (
	"github.com/emanueltimlopez/books-motivation/internal/book"
	"github.com/emanueltimlopez/books-motivation/internal/user"
)

type BookShelf struct {
	user  *user.User
	books []*book.Book
}
