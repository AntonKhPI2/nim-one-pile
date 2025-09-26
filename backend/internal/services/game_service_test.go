package services

import (
	"context"
	"errors"
	"testing"

	"github.com/AntonKhPI2/nim-one-pile/internal/game"
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
	return nil, errors.New("not found")
}

func TestGameService_New_And_Move(t *testing.T) {
	svc := NewGameService(newFakeRepo())

	cfg := game.Config{Variant: game.Normal, N: 10, K: 3}
	st, err := svc.New(context.Background(), cfg)
	if err != nil {
		t.Fatalf("svc.New: %v", err)
	}
	if st.ID == "" {
		t.Fatalf("expected non-empty id")
	}

	prevRemaining := st.Remaining

	out, err := svc.HumanMove(context.Background(), st.ID, 1)
	if err != nil {
		t.Fatalf("svc.HumanMove: %v", err)
	}
	if out.Remaining >= prevRemaining {
		t.Fatalf("remaining should decrease, was %d now %d", prevRemaining, out.Remaining)
	}
}

func TestGameService_Get_NotFound(t *testing.T) {
	svc := NewGameService(newFakeRepo())
	if _, err := svc.Get(context.Background(), "nope"); err == nil {
		t.Fatalf("expected error for missing game")
	}
}
