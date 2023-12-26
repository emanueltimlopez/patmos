package user

import (
	"context"

	"github.com/emanueltimlopez/books-motivation/internal/plan"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUserPlan(ctx context.Context, plan plan.Plan, user string) error
}
