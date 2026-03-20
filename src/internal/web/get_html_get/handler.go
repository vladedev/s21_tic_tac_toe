package get_html_get

import (
	_ "embed"
	"net/http"
)

//go:embed assets/index.html
var indexHtml string

type Handler struct{}

func New() *Handler { return &Handler{} }

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", h.returnHtml)
}

func (h *Handler) returnHtml(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexHtml))
}
