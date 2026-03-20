package minimax

import (
	"Project03-Go_Bootcamp/internal/domain"
	"context"
	"errors"
	"math"
)

type Repository interface {
	Get(ctx context.Context, id string) (domain.Game, error)
	Save(ctx context.Context, g domain.Game, id string) error
}

type Service struct {
	repo Repository
}

func New(repository Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) WriteNewState(ctx context.Context, prevId string, next domain.Game) (domain.Game, error) {
	game, err := s.repo.Get(ctx, prevId)
	if err != nil {
		return domain.Game{}, err
	}

	if err = game.ValidateNextState(next); err != nil {
		return domain.Game{}, errors.New("ValidateNextState: " + err.Error())
	}

	game = next
	game.SetGameWinner()

	if err = s.repo.Save(ctx, game, prevId); err != nil {
		return domain.Game{}, err
	}
	return game, nil
}

func (s *Service) NextMove(ctx context.Context, id string) (domain.Game, error) {
	g, err := s.repo.Get(ctx, id)
	if err != nil {
		return domain.Game{}, err
	}

	if g.GetWinner() != domain.NoneWinner {
		return domain.Game{}, errors.New("game is already over")
	}

	if g.GetTurn() != domain.TurnO {
		return domain.Game{}, errors.New("not computer's turn")
	}

	bestScore := math.MinInt
	var bestMove [2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if err := g.Board().SetPoint(i, j, domain.PointValueO); err == nil {
				score := minimax(g.Board(), 0, false)
				g.Board().SetPoint(i, j, domain.PointValueEmpty)

				if score > bestScore {
					bestScore = score
					bestMove = [2]int{i, j}
				}
			}
		}
	}

	g.Board().SetPoint(bestMove[0], bestMove[1], domain.PointValueO)
	g.NextTurn()
	g.SetGameWinner()

	if err = s.repo.Save(ctx, g, id); err != nil {
		return domain.Game{}, err
	}
	return g, nil
}

func minimax(b *domain.Board, depth int, isMaximizing bool) int {
	if won, _ := b.CheckWinner(domain.PointValueO); won {
		return 10 - depth
	}
	if won, _ := b.CheckWinner(domain.PointValueX); won {
		return depth - 10
	}
	if b.IsFull() {
		return 0
	}

	if isMaximizing {
		best := math.MinInt
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if b.GetPoint(i, j) == domain.PointValueEmpty {
					b.SetPoint(i, j, domain.PointValueO)
					score := minimax(b, depth+1, false)
					b.SetPoint(i, j, domain.PointValueEmpty)
					if score > best {
						best = score
					}
				}
			}
		}
		return best
	}

	best := math.MaxInt
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.GetPoint(i, j) == domain.PointValueEmpty {
				b.SetPoint(i, j, domain.PointValueX)
				score := minimax(b, depth+1, true)
				b.SetPoint(i, j, domain.PointValueEmpty)
				if score < best {
					best = score
				}
			}
		}
	}
	return best
}
