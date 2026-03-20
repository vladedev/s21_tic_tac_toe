package domain

import "context"

type GameService interface {
	WriteNewState(ctx context.Context, prevId string, next Game) (Game, error)
	NextMove(ctx context.Context, id string) (Game, error)
}

type GameRepository interface {
	Get(ctx context.Context, id string) (Game, error)
	Save(ctx context.Context, g Game, id string) error
}
