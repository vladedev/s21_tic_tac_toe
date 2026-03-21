package datastorage

import (
	"sync"
)

// InMemoryStorage — потокобезопасный класс-хранилище игр sync.Map
type GameStorage struct {
	games sync.Map
}

func NewGameStorage() *GameStorage {
	return &GameStorage{}
}

// моделт игрового поля
type GameDTO struct {
	ID    string
	Board [3][3]int
}
