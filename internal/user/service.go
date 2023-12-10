package user

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewUserService(us Repository) *Service {
	return &Service{
		repo: us,
	}
}

func (us *Service) GetUser(ctx context.Context, id string) (*User, error) {
	user, err := us.repo.GetUser(ctx, id)
	return user, err
}
