package main

import (
	"embed"
	"net/http"
	"os"

	"github.com/emanueltimlopez/books-motivation/internal/routes"
)

//go:embed web/templates/*.html
var templates embed.FS

//go:embed web/templates/components/*.html
var templatesComponents embed.FS

//go:embed web/static
var staticFiles embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.Templates = templates
	routes.TemplatesComponents = templatesComponents
	routes.StaticFiles = staticFiles
	router := routes.NewRouter()

	err_ := http.ListenAndServe(":"+port, router)
	if err_ != nil {
		panic(err_)
	}
}
