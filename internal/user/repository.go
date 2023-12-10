package user

import "context"

type Repository interface {
	GetUser(ctx context.Context, id string) (*User, error)
}
