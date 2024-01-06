package book_usecases_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	book_usecases "github.com/emanueltimlopez/patmos/internal/book/use-cases"
)

// Should return the params from form book
func TestGetParamsSearchBook(t *testing.T) {
	want := book_usecases.SearchBookParams{
		Text: "sarasa",
	}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("search=sarasa"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	params := book_usecases.GetParamsSearchBook(req)

	if want.Text != params.Text {
		t.Fatalf(`%q != %v`, want.Text, params.Text)
	}
}
