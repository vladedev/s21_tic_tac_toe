package web

// для отправки запроса от клиента в domain
type GameRequest struct {
	Board [3][3]int `json:"board"` // модель игрового поля
}

// для получения запроса к клиенту из domain
type GameResponse struct {
	ID      string    `json:"id"`
	Board   [3][3]int `json:"board"`
	Over    bool      `json:"over"`
	Message string    `json:"message,omitempty"`
}
