package plan_usecases_test

import (
	"testing"

	plan_usecases "github.com/emanueltimlopez/books-motivation/internal/plan/use-cases"
	"github.com/emanueltimlopez/books-motivation/internal/user"
)

// Should return the params from form book
func TestBooksLeft(t *testing.T) {
	want := 737

	user := user.User{}
	user.Plan.Words = 300
	user.Plan.Minutes = 15
	user.Birth = 1989

	left := plan_usecases.BooksLeft(&user)

	if want != left {
		t.Fatalf(`%q != %v`, want, left)
	}
}
