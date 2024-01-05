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
	pages, _err := strconv.Atoi(r.Form.Get("pages"))
	if _err != nil {
		fmt.Println("[GetParamsFormBook:pages]", _err)
	}
	image := r.Form.Get("cover")

	return ParamsFormBook{
		Title:  title,
		Author: author,
		Pages:  pages,
		Image:  image,
		Isbn:   isbn,
	}
}
