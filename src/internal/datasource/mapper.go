package datasource

import "Project03-Go_Bootcamp/internal/domain"

func ToDS(g domain.Game) GameDS {
	return GameDS{
		ID:     g.GetID(),
		Matrix: g.Board().Matrix(),
		Turn:   g.GetTurn(),
		Winner: g.GetWinner(),
	}
}

func ToDomain(g GameDS) domain.Game {
	return domain.SetGame(g.ID, domain.NewBoard(g.Matrix), g.Turn, g.Winner)
}
