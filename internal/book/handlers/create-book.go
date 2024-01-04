package book

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emanueltimlopez/books-motivation/internal/book"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	r.ParseForm()
	title := r.Form.Get("title")
	author := r.Form.Get("author")
	pages, _err := strconv.Atoi(r.Form.Get("pages"))
	if _err != nil {
		fmt.Println(_err)
	}

	image := r.Form.Get("image")

	booksService := book.NewBookService(dbRepository)
	err := booksService.CreateBook(ctx, book.Book{
		Title:  title,
		Author: author,
		ID:     uuid.New().String(),
		Pages:  pages,
		Image:  image,
	}, userSupa.ID)

	if err != nil {
		fmt.Println(err)
	}
}
