package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AntonKhPI2/nim-one-pile/internal/game"
	"github.com/AntonKhPI2/nim-one-pile/internal/repositories"
	"github.com/AntonKhPI2/nim-one-pile/internal/services"
)

type gameHandlers struct{ svc *services.GameService }

func RegisterGameRoutes(mux *http.ServeMux, svc *services.GameService) {
	h := &gameHandlers{svc: svc}
	mux.HandleFunc("/api/new-game", h.newGame)
	mux.HandleFunc("/api/take", h.take)
	mux.HandleFunc("/api/game/", h.get)
}
func (h *gameHandlers) take(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		ID   string `json:"id"`
		Take int    `json:"take"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	st, err := h.svc.HumanMove(r.Context(), req.ID, req.Take)
	if err != nil {
		status := http.StatusBadRequest
		if err == repositories.ErrNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	respondJSON(w, st)
}

func (h *gameHandlers) get(w http.ResponseWriter, r *http.Request) {
	const prefix = "/api/game/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.NotFound(w, r)
		return
	}
	id := r.URL.Path[len(prefix):]
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	st, err := h.svc.Get(r.Context(), id)
	if err != nil {
		status := http.StatusBadRequest
		if err == repositories.ErrNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	respondJSON(w, st)
}

func (h *gameHandlers) newGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var cfg game.Config
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	st, err := h.svc.New(r.Context(), cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondJSON(w, st)
}

func respondJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(v)
}
