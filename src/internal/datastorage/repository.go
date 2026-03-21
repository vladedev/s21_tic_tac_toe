package datastorage

import (
	"tic-tac-toe/internal/domain"
)

type gameRepository struct {
	storage *GameStorage
}

func NewGameRepository(storage *GameStorage) domain.GameRepository {
	return &gameRepository{storage: storage}
}

// реализация метода save, для сохранения текущей игры с id и состоянием поля
func (r *gameRepository) Save(game *domain.Game) error {
	r.storage.games.Store(game.ID, ToDTO(game))
	return nil
}

// реализация метода get, находит игру с id и состоянием поля
func (r *gameRepository) Get(id string) (*domain.Game, error) {
	val, ok := r.storage.games.Load(id)
	if !ok {
		return nil, domain.ErrGameNotFound
	}
	return ToDomain(val.(*GameDTO)), nil
}
