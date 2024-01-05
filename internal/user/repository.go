package user

import (
	"context"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUserPlan(ctx context.Context, user *User) (*User, error)
}
