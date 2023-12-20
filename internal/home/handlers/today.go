package home

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/book"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"

	supa "github.com/nedpals/supabase-go"
)

func TodayHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	tmpl := template.Must(template.ParseFiles("./web/templates/today.html"))
	ctx := context.Background()
	dbClient := supabase.InitClient()

	dbRepository := supabase.NewSupabaseRepository(dbClient)

	userService := user.NewUserService(dbRepository)
	booksService := book.NewBookService(dbRepository)

	_user, err := userService.GetUser(ctx, userSupa.ID)
	if err != nil {
		fmt.Println("[TodayHandler:getuser]", err)
	}

	if _user == nil {
		tmpl.Execute(w, nil)
		return
	}

	book, err := booksService.GetBook(ctx, _user.Book)
	if err != nil {
		fmt.Println("[TodayHandler:getbook]", err)
	}

	tmpl.Execute(w, book)
}
