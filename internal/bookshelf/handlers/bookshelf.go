package bookshelf

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/emanueltimlopez/patmos/internal/book"
	"github.com/emanueltimlopez/patmos/internal/bookshelf"
	"github.com/emanueltimlopez/patmos/internal/platform/supabase"
	"github.com/emanueltimlopez/patmos/internal/user"
	supa "github.com/nedpals/supabase-go"
)

var Tmpl *template.Template

type BookshelfFromView struct {
	Books []book.Book
}

func BooksHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	userService := user.NewUserService(dbRepository)
	bookshelfService := bookshelf.NewBookShelfService(dbRepository)

	user, err := userService.GetUser(ctx, userSupa.ID)
	if err != nil {
		fmt.Println(err)
	}

	books, err := bookshelfService.GetBookShelf(ctx, user.ID)
	if err != nil {
		fmt.Println(err)
	}

	Tmpl.ExecuteTemplate(w, "books.html", BookshelfFromView{Books: books})
}
