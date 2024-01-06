package book_usecases_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	book_usecases "github.com/emanueltimlopez/books-motivation/internal/book/use-cases"
)

// Should return the params from form book
func TestGetParamsFormBook(t *testing.T) {
	want := book_usecases.ParamsFormBook{
		Title:  "sarasa",
		Author: "sarasa2",
		Pages:  100,
		Image:  "sarasa3",
		Isbn:   "sarasa4",
	}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("title=sarasa&author=sarasa2&pages=100&cover=sarasa3&isbn=sarasa4"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	params := book_usecases.GetParamsFormBook(req)

	if want.Title != params.Title {
		t.Fatalf(`%q != %v`, want.Title, params.Title)
	}
	if want.Author != params.Author {
		t.Fatalf(`%q != %v`, want.Author, params.Author)
	}
	if want.Pages != params.Pages {
		t.Fatalf(`%q != %v`, want.Pages, params.Pages)
	}
	if want.Image != params.Image {
		t.Fatalf(`%q != %v`, want.Image, params.Image)
	}
	if want.Isbn != params.Isbn {
		t.Fatalf(`%q != %v`, want.Isbn, params.Isbn)
	}
}

func TestGetParamsFormBook_BadPages(t *testing.T) {
	want := book_usecases.ParamsFormBook{
		Title:  "sarasa",
		Author: "sarasa2",
		Pages:  0,
		Image:  "sarasa3",
		Isbn:   "sarasa4",
	}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("title=sarasa&author=sarasa2&pages=re&cover=sarasa3&isbn=sarasa4"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	params := book_usecases.GetParamsFormBook(req)

	if want.Title != params.Title {
		t.Fatalf(`%q != %v`, want.Title, params.Title)
	}
	if want.Author != params.Author {
		t.Fatalf(`%q != %v`, want.Author, params.Author)
	}
	if want.Pages != params.Pages {
		t.Fatalf(`%q != %v`, want.Pages, params.Pages)
	}
	if want.Image != params.Image {
		t.Fatalf(`%q != %v`, want.Image, params.Image)
	}
	if want.Isbn != params.Isbn {
		t.Fatalf(`%q != %v`, want.Isbn, params.Isbn)
	}
}
