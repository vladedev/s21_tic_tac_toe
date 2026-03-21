package application

import (
	"tic-tac-toe/internal/domain"
)

// структура реализации интерфейса GameRepository
type gameService struct {
	repo domain.GameRepository
}

func NewGameService(repo domain.GameRepository) domain.GameService {
	return &gameService{repo: repo}
}

// Полноценный алгоритм Minimax
func (s *gameService) GetNextMove(board domain.Board) (int, int, error) {
	// Проверяем, не закончена ли игра
	if s.IsGameOver(board) {
		return -1, -1, domain.ErrGameAlreadyOver
	}

	// Проверяем, есть ли доступные ходы
	availableMoves := s.getAvailableMoves(board)
	if len(availableMoves) == 0 {
		return -1, -1, domain.ErrGameAlreadyOver
	}

	// Запускаем Minimax для компьютера (игрок 2)
	bestScore := -1000
	bestRow, bestCol := -1, -1

	for _, move := range availableMoves {
		// Делаем пробный ход
		newBoard := s.copyBoard(board)
		newBoard[move.row][move.col] = 2 // компьютер

		// Вычисляем оценку хода
		score := s.minimax(newBoard, 0, false)

		// Выбираем лучший ход
		if score > bestScore {
			bestScore = score
			bestRow, bestCol = move.row, move.col
		}
	}

	return bestRow, bestCol, nil
}

// Minimax алгоритм с альфа-бета отсечением
func (s *gameService) minimax(board domain.Board, depth int, isMaximizing bool) int {
	// Проверяем победу компьютера (игрок 2)
	if s.checkWinner(board, 2) {
		return 10 - depth
	}
	// Проверяем победу игрока (игрок 1)
	if s.checkWinner(board, 1) {
		return -10 + depth
	}
	// Проверяем ничью
	if s.isBoardFull(board) {
		return 0
	}

	if isMaximizing {
		// Ход компьютера (максимизируем)
		bestScore := -1000
		moves := s.getAvailableMoves(board)

		for _, move := range moves {
			newBoard := s.copyBoard(board)
			newBoard[move.row][move.col] = 2 // компьютер
			score := s.minimax(newBoard, depth+1, false)
			if score > bestScore {
				bestScore = score
			}
		}
		return bestScore
	} else {
		// Ход игрока (минимизируем)
		bestScore := 1000
		moves := s.getAvailableMoves(board)

		for _, move := range moves {
			newBoard := s.copyBoard(board)
			newBoard[move.row][move.col] = 1 // игрок
			score := s.minimax(newBoard, depth+1, true)
			if score < bestScore {
				bestScore = score
			}
		}
		return bestScore
	}
}

func (s *gameService) ValidateBoard(current, updated domain.Board) error {
	// Проверяем, что изменился только один слот и он стал 1 (ход игрока)
	changes := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if current[i][j] != updated[i][j] {
				changes++
				// Проверяем, что ход корректен: клетка была пустой и стала 1
				if updated[i][j] != 1 || current[i][j] != 0 {
					return domain.ErrInvalidMove
				}
			}
		}
	}
	if changes != 1 {
		return domain.ErrInvalidMove
	}
	return nil
}

func (s *gameService) IsGameOver(board domain.Board) bool {
	// Проверка на победу любого игрока
	if s.checkWinner(board, 1) || s.checkWinner(board, 2) {
		return true
	}
	// Проверка на ничью (все клетки заполнены)
	return s.isBoardFull(board)
}

// Проверка победителя для конкретного игрока
func (s *gameService) checkWinner(board domain.Board, player int) bool {
	// Проверка горизонталей
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}
	// Проверка вертикалей
	for j := 0; j < 3; j++ {
		if board[0][j] == player && board[1][j] == player && board[2][j] == player {
			return true
		}
	}
	// Проверка диагоналей
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

// Проверка, заполнена ли доска
func (s *gameService) isBoardFull(board domain.Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// Получение всех доступных ходов
func (s *gameService) getAvailableMoves(board domain.Board) []struct{ row, col int } {
	var moves []struct{ row, col int }
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				moves = append(moves, struct{ row, col int }{i, j})
			}
		}
	}
	return moves
}

// Копирование доски
func (s *gameService) copyBoard(board domain.Board) domain.Board {
	newBoard := domain.Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newBoard[i][j] = board[i][j]
		}
	}
	return newBoard
}
