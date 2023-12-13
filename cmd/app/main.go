package main

import (
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/routes"
)

func main() {
	router := routes.NewRouter()

	err_ := http.ListenAndServe(":8000", router)
	if err_ != nil {
		panic(err_)
	}
}
