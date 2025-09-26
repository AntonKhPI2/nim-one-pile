package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AntonKhPI2/nim-one-pile/internal/game"
	"github.com/AntonKhPI2/nim-one-pile/internal/repositories"
	"github.com/AntonKhPI2/nim-one-pile/internal/services"
	"golang.org/x/net/context"
)

type fakeRepo struct {
	store map[string]*game.State
}

func newFakeRepo() *fakeRepo { return &fakeRepo{store: map[string]*game.State{}} }

func (r *fakeRepo) Save(ctx context.Context, s *game.State) error {
	cp := *s
	r.store[s.ID] = &cp
	return nil
}

func (r *fakeRepo) Get(ctx context.Context, id string) (*game.State, error) {
	if s, ok := r.store[id]; ok {
		cp := *s
		return &cp, nil
	}
	return nil, repositories.ErrNotFound
}

func TestNewGameHandler(t *testing.T) {
	repo := newFakeRepo()
	svc := services.NewGameService(repo)

	mux := http.NewServeMux()
	RegisterGameRoutes(mux, svc)

	body, _ := json.Marshal(map[string]any{
		"variant": "normal",
		"n":       21,
		"k":       3,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/new-game", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status=%d, body=%s", w.Code, w.Body.String())
	}
	var st game.State
	if err := json.Unmarshal(w.Body.Bytes(), &st); err != nil {
		t.Fatalf("bad json: %v", err)
	}
	if st.ID == "" || st.K != 3 || st.Remaining == 0 {
		t.Fatalf("unexpected state: %+v", st)
	}
}

func TestMoveAndStateHandlers(t *testing.T) {
	repo := newFakeRepo()
	svc := services.NewGameService(repo)
	mux := http.NewServeMux()
	RegisterGameRoutes(mux, svc)

	newBody, _ := json.Marshal(map[string]any{"variant": "misere", "n": 7, "k": 3})
	req := httptest.NewRequest(http.MethodPost, "/api/new-game", bytes.NewReader(newBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("new-game status=%d body=%s", w.Code, w.Body.String())
	}
	var st game.State
	_ = json.Unmarshal(w.Body.Bytes(), &st)
	if st.ID == "" {
		t.Fatal("empty id")
	}

	moveBody, _ := json.Marshal(map[string]any{"id": st.ID, "take": 2})
	req2 := httptest.NewRequest(http.MethodPost, "/api/take", bytes.NewReader(moveBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("move status=%d body=%s", w2.Code, w2.Body.String())
	}
	var st2 game.State
	_ = json.Unmarshal(w2.Body.Bytes(), &st2)
	if st2.Remaining >= st.Remaining {
		t.Fatalf("remaining should decrease: was %d now %d", st.Remaining, st2.Remaining)
	}

	req3 := httptest.NewRequest(http.MethodGet, "/api/game/"+st.ID, nil)
	w3 := httptest.NewRecorder()
	mux.ServeHTTP(w3, req3)

	if w3.Code != http.StatusOK {
		t.Fatalf("state status=%d body=%s", w3.Code, w3.Body.String())
	}
	var st3 game.State
	_ = json.Unmarshal(w3.Body.Bytes(), &st3)
	if st3.ID != st.ID {
		t.Fatalf("wrong game id: %s", st3.ID)
	}
}
