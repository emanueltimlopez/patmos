package user

import (
	"context"
	"fmt"

	"github.com/emanueltimlopez/patmos/internal/plan"
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

func (us *Service) CreateUser(ctx context.Context, user User) error {
	err := us.repo.CreateUser(ctx, user)
	return err
}

func (us *Service) UpdateUserPlan(ctx context.Context, plan plan.Plan, id string) (*User, error) {
	_user, err := us.repo.GetUser(ctx, id)
	if err != nil {
		fmt.Println("[supabase:UpdateUserPlan:user]", err)
		return nil, err
	}
	_user.Plan = plan

	user, _err := us.repo.UpdateUserPlan(ctx, _user)
	return user, _err
}
