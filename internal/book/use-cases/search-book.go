package book_usecases

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type BookFromAPI struct {
	Cover       string   `json:"cover_edition_key"`
	FullText    bool     `json:"has_fulltext"`
	Edition     int      `json:"edition_count"`
	Title       string   `json:"title"`
	Author      []string `json:"author_name"`
	PublishYear int      `json:"first_publish_year"`
	Key         string   `json:"key"`
	Ia          []string `json:"ia"`
	AuthorKey   []string `json:"author_key"`
	PublicScan  bool     `json:"public_scan_b"`
	Isbn        []string `json:"isbn"`
	Pages       int      `json:"number_of_pages_median"`
}

type BooksFromAPI struct {
	Start    int           `json:"start"`
	NumFound int           `json:"num_found"`
	Docs     []BookFromAPI `json:"docs"`
}

func SearchBook(searchQuery string) ([]BookFromAPI, error) {
	response, err := http.Get(fmt.Sprintf("https://openlibrary.org/search.json?q=%s", url.QueryEscape(searchQuery)))
	if err != nil {
		fmt.Print("[SearchBook:get]", err.Error())
		return nil, err
	}

	booksData, _err := io.ReadAll(response.Body)
	if _err != nil {
		fmt.Println("[SearchBook:read]", _err)
		return nil, _err
	}

	var books BooksFromAPI
	__err := json.Unmarshal(booksData, &books)
	if __err != nil {
		fmt.Println("[SearchBook:unmarshal]", __err)
		return nil, __err
	}

	return books.Docs, nil
}
