package main

import (
	"net/http"

	"books.app/internal/routes"
)

func main() {
	router := routes.NewRouter()

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
