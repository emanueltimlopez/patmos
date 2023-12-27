package main

import (
	"net/http"
	"os"

	"github.com/emanueltimlopez/books-motivation/internal/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routes.NewRouter()

	err_ := http.ListenAndServe(port, router)
	if err_ != nil {
		panic(err_)
	}
}
