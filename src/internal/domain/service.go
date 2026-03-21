package domain

// интерфейс который описывает методы следующего хода, валидации поля, проерки на окончание игры.(порядок действий игры)
type GameService interface {
	GetNextMove(board Board) (row, col int, err error)
	ValidateBoard(current, updated Board) error
	IsGameOver(board Board) bool
}
