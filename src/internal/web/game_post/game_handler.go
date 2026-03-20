package game_post

import (
	"Project03-Go_Bootcamp/internal/domain"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type Service interface {
	WriteNewState(ctx context.Context, prevId string, next domain.Game) (domain.Game, error)
	NextMove(ctx context.Context, id string) (domain.Game, error)
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/game/", h.postGame)
}

func (h *Handler) postGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/game/")
	if id == "" || strings.Contains(id, "/") {
		writeErr(w, http.StatusBadRequest, "invalid game id")
		return
	}

	var req GameWebRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		return
	}

	ctx := r.Context()

	updatedGame, err := h.svc.WriteNewState(ctx, id, FromWeb(id, req))
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}

	if updatedGame.GetWinner() != domain.NoneWinner {
		json.NewEncoder(w).Encode(ToWeb(updatedGame))
		return
	}

	updatedGame, err = h.svc.NextMove(ctx, id)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(ToWeb(updatedGame))
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}
