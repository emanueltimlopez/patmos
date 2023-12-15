package goals

import "context"

type Repository interface {
	GetGoals(ctx context.Context, id string) ([]*Goal, error)
	CreateGoal(ctx context.Context, goal Goal) error
}
