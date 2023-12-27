package user

import "github.com/emanueltimlopez/books-motivation/internal/plan"

type User struct {
	ID    string    `json:"id"`
	Book  string    `json:"book,omitempty"`
	Plan  plan.Plan `json:"plan,omitempty"`
	Birth int       `json:"birth"`
}
