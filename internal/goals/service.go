package goals

import "context"

type Service struct {
	repo Repository
}

func NewGoalService(gs Repository) *Service {
	return &Service{
		repo: gs,
	}
}

func (gs *Service) GetGoals(ctx context.Context, id string) ([]*Goal, error) {
	goals, err := gs.repo.GetGoals(ctx, id)
	return goals, err
}

func (gs *Service) CreateGoal(ctx context.Context, goal Goal) error {
	err := gs.repo.CreateGoal(ctx, goal)
	return err
}
