package game_post

import "Project03-Go_Bootcamp/internal/domain"

func ToWeb(g domain.Game) GameWebResponse {
	return GameWebResponse{
		Board:  g.Board().Matrix(),
		Winner: g.GetWinner(),
	}
}

func FromWeb(id string, req GameWebRequest) domain.Game {
	return domain.SetGame(id, domain.NewBoard(req.Board), domain.TurnO, domain.NoneWinner)
}
