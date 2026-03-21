package web

import "tic-tac-toe/internal/domain"

func ToDomain(gameID string, req GameRequest) *domain.Game {
	return &domain.Game{
		ID:    gameID,
		Board: req.Board,
	}
}

func FromDomain(game *domain.Game, over bool, message string) GameResponse {
	return GameResponse{
		ID:      game.ID,
		Board:   game.Board,
		Over:    over,
		Message: message,
	}
}
