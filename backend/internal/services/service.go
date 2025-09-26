package services

import (
	"context"

	"github.com/AntonKhPI2/nim-one-pile/internal/game"
	"github.com/AntonKhPI2/nim-one-pile/internal/repositories"
)

type GameService struct{ repo repositories.GameRepo }

func NewGameService(r repositories.GameRepo) *GameService { return &GameService{repo: r} }

func (s *GameService) New(ctx context.Context, cfg game.Config) (*game.State, error) {
	if cfg.Variant != game.Normal && cfg.Variant != game.Misere {
		cfg.Variant = game.Normal
	}

	id := newID()

	st, err := game.NewGame(cfg, id)
	if err != nil {
		return nil, err
	}

	if st.PlayerTurn == "computer" && st.Winner == "" {
		game.ApplyComputerMove(st)
	}
	
	if err := s.repo.Save(ctx, st); err != nil {
		return nil, err
	}
	return st, nil
}

func (s *GameService) HumanMove(ctx context.Context, id string, take int) (*game.State, error) {
	st, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := game.ApplyHumanMove(st, take); err != nil {
		return nil, err
	}
	if st.Winner == "" {
		game.ApplyComputerMove(st)
	}
	if err := s.repo.Save(ctx, st); err != nil {
		return nil, err
	}
	return st, nil
}

func (s *GameService) Get(ctx context.Context, id string) (*game.State, error) {
	return s.repo.Get(ctx, id)
}
