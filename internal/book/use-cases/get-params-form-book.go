package book_usecases

import (
	"fmt"
	"net/http"
	"strconv"
)

type ParamsFormBook struct {
	Title  string
	Author string
	Pages  int
	Image  string
	Isbn   string
}

func GetParamsFormBook(r *http.Request) ParamsFormBook {
	r.ParseForm()
	title := r.Form.Get("title")
	author := r.Form.Get("author")
	isbn := r.Form.Get("isbn")
	image := r.Form.Get("cover")
	pages, err := strconv.Atoi(r.Form.Get("pages"))
	if err != nil {
		fmt.Println("[GetParamsFormBook:pages]", err)
		pages = 0
	}

	return ParamsFormBook{
		Title:  title,
		Author: author,
		Pages:  pages,
		Image:  image,
		Isbn:   isbn,
	}
}
