package routes

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/emanueltimlopez/patmos/internal/auth"
	authHandlers "github.com/emanueltimlopez/patmos/internal/auth/handlers"
	book "github.com/emanueltimlopez/patmos/internal/book/handlers"
	bookshelf "github.com/emanueltimlopez/patmos/internal/bookshelf/handlers"
	home "github.com/emanueltimlopez/patmos/internal/home/handlers"
	plan "github.com/emanueltimlopez/patmos/internal/plan/handlers"
)

var Templates embed.FS
var TemplatesComponents embed.FS
var StaticFiles embed.FS

func NewRouter() http.Handler {

	var tmpl = template.Must(template.ParseFS(Templates, "web/templates/*.html"))
	var tmplComponents = template.Must(template.ParseFS(TemplatesComponents, "web/templates/components/*.html"))
	home.Tmpl = tmpl
	authHandlers.Tmpl = tmpl
	authHandlers.TmplComponents = tmplComponents
	bookshelf.Tmpl = tmpl
	plan.Tmpl = tmpl
	book.Tmpl = tmpl
	book.TmplComponents = tmplComponents

	mux := http.NewServeMux()

	sub, err := fs.Sub(StaticFiles, "web")
	if err != nil {
		panic(err)
	}
	fileServer := http.FileServer(http.FS(sub))
	mux.Handle("/static/", fileServer)

	mux.Handle("/", auth.NewAuth(home.IndexHandler))
	mux.HandleFunc("/login", authHandlers.LoginPageHandler)
	mux.HandleFunc("/signup", authHandlers.SignupPageHandler)
	mux.HandleFunc("/login-user", authHandlers.LoginHandler)
	mux.HandleFunc("/signup-user", authHandlers.SignupHandler)
	mux.HandleFunc("/logout-user", authHandlers.LogoutHandler)
	mux.Handle("/bookshelf", auth.NewAuth(bookshelf.BooksHandler))
	mux.Handle("/today", auth.NewAuth(home.TodayHandler))
	mux.Handle("/add-book", auth.NewAuth(book.AddBookHandler))
	mux.Handle("/set-book", auth.NewAuth(book.SetBookHandler))
	mux.Handle("/create-book", auth.NewAuth(book.CreateBookHandler))
	mux.Handle("/search-book", auth.NewAuth(book.SearchBookHandler))
	mux.Handle("/plan", auth.NewAuth(plan.PlanHandler))
	mux.Handle("/update-plan", auth.NewAuth(plan.UpdatePlanHandler))

	return mux
}
