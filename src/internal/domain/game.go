package domain

import "github.com/google/uuid"

type Game struct {
	ID    string // UUID игры
	Board Board  // состояние поля
}

// функция создания нового поля с новым UUID
func NewGame() *Game {
	return &Game{
		ID:    uuid.New().String(),
		Board: Board{},
	}
}
