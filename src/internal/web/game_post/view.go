package game_post

type GameWebRequest struct {
	Board [3][3]int `json:"board"`
}

type GameWebResponse struct {
	Board  [3][3]int `json:"board"`
	Winner int       `json:"winner"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
