package plan_usecases_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	plan_usecases "github.com/emanueltimlopez/patmos/internal/plan/use-cases"
)

// Should return the params from form book
func TestGetParamsPlan(t *testing.T) {
	want := plan_usecases.PlanParams{
		Minutes:  100,
		Sessions: 2,
		Words:    300,
	}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("minutes=100&sessions=2&words=300"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	params := plan_usecases.GetParamsPlan(req)

	if want.Minutes != params.Minutes {
		t.Fatalf(`%q != %v`, want.Minutes, params.Minutes)
	}
	if want.Sessions != params.Sessions {
		t.Fatalf(`%q != %v`, want.Sessions, params.Sessions)
	}
	if want.Words != params.Words {
		t.Fatalf(`%q != %v`, want.Words, params.Words)
	}
}

func TestGetParamsFormBook_BadParams(t *testing.T) {
	want := plan_usecases.PlanParams{
		Minutes:  0,
		Sessions: 0,
		Words:    0,
	}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("minutes=ew&sessions=we&words=rt"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	params := plan_usecases.GetParamsPlan(req)

	if want.Minutes != params.Minutes {
		t.Fatalf(`%q != %v`, want.Minutes, params.Minutes)
	}
	if want.Sessions != params.Sessions {
		t.Fatalf(`%q != %v`, want.Sessions, params.Sessions)
	}
	if want.Words != params.Words {
		t.Fatalf(`%q != %v`, want.Words, params.Words)
	}
}
