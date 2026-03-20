package domain

import "errors"

const (
	TurnX = 10
	TurnO = 11
)

const (
	NoneWinner = 20
	Draw       = 21
	WinnerX    = 22
	WinnerO    = 23
)

type Game struct {
	id     string
	board  Board
	turn   int
	winner int
}

func NewGame(id string) Game {
	return Game{
		id:     id,
		board:  Board{},
		turn:   TurnX,
		winner: NoneWinner,
	}
}

func SetGame(id string, b Board, turn, winner int) Game {
	return Game{
		id:     id,
		board:  b,
		turn:   turn,
		winner: winner,
	}
}

func (g *Game) GetID() string {
	return g.id
}

func (g *Game) Board() *Board {
	return &g.board
}

func (g *Game) GetTurn() int {
	return g.turn
}

func (g *Game) GetWinner() int {
	return g.winner
}

func (g *Game) NextTurn() {
	if g.turn == TurnX {
		g.turn = TurnO
	} else if g.turn == TurnO {
		g.turn = TurnX
	}
}

func (g *Game) SetGameWinner() (int, bool) {
	if won, _ := g.board.CheckWinner(PointValueO); won {
		g.winner = WinnerO
		return WinnerO, true
	}
	if won, _ := g.board.CheckWinner(PointValueX); won {
		g.winner = WinnerX
		return WinnerX, true
	}
	if g.board.IsFull() {
		g.winner = Draw
		return Draw, true
	}
	return -1, false
}

func (g *Game) ValidateNextState(next Game) error {
	if g.GetWinner() != NoneWinner || next.GetWinner() != NoneWinner {
		return errors.New("game already finished")
	}

	prevXCount, prevOCount := 0, 0
	nextXCount, nextOCount := 0, 0
	changes := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			old := g.board.GetPoint(i, j)
			cur := next.board.GetPoint(i, j)

			if old == PointValueX {
				prevXCount++
			}
			if old == PointValueO {
				prevOCount++
			}
			if cur == PointValueX {
				nextXCount++
			}
			if cur == PointValueO {
				nextOCount++
			}

			if old != PointValueEmpty && old != cur {
				return errors.New("attempt to overwrite existing cell")
			}
			if old != cur {
				changes++
			}
		}
	}

	if changes != 1 {
		return errors.New("exactly one cell must change")
	}

	switch g.GetTurn() {
	case TurnX:
		if !(nextXCount == prevXCount+1 && nextOCount == prevOCount) {
			return errors.New("X must add exactly one mark")
		}
	case TurnO:
		if !(nextOCount == prevOCount+1 && nextXCount == prevXCount) {
			return errors.New("O must add exactly one mark")
		}
	default:
		return errors.New("invalid turn value")
	}

	return nil
}
