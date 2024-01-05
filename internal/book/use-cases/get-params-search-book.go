package book_usecases

import (
	"net/http"
)

type SearchBookParams struct {
	Text string
}

func GetParamsSearchBook(r *http.Request) SearchBookParams {
	r.ParseForm()
	search := r.Form.Get("search")

	return SearchBookParams{
		Text: search,
	}
}
