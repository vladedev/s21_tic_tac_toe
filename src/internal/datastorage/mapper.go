package datastorage

import "tic-tac-toe/internal/domain"

// преобразовывает игру из DTO в DTO domain из хранилища
func ToDomain(dto *GameDTO) *domain.Game {
	return &domain.Game{
		ID:    dto.ID,
		Board: dto.Board,
	}
}

// преобразовывает игру из domain в DTO для хранилища
func ToDTO(game *domain.Game) *GameDTO {
	return &GameDTO{
		ID:    game.ID,
		Board: game.Board,
	}
}
