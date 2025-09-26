package game

import "errors"

func NewGame(cfg Config, id string) (*State, error) {
	if cfg.N <= 0 || cfg.K <= 0 {
		return nil, errors.New("N and K must be > 0")
	}
	s := &State{
		ID:         id,
		Variant:    cfg.Variant,
		K:          cfg.K,
		Remaining:  cfg.N,
		PlayerTurn: "human",
	}
	if !isLosingStart(cfg.Variant, cfg.N, cfg.K) {
		s.PlayerTurn = "computer"
		ApplyComputerMove(s)
	}
	return s, nil
}

func ApplyHumanMove(s *State, take int) error {
	if s.Winner != "" {
		return errors.New("game already finished")
	}
	if s.PlayerTurn != "human" {
		return errors.New("not human turn")
	}
	if take < 1 || take > s.K {
		return errors.New("invalid take")
	}
	if take > s.Remaining {
		return errors.New("take exceeds remaining")
	}

	s.Remaining -= take
	if s.Remaining == 0 {
		if s.Variant == Normal {
			s.Winner = "human"
		} else {
			s.Winner = "computer"
		}
		return nil
	}
	s.PlayerTurn = "computer"
	return nil
}

func ApplyComputerMove(s *State) {
	if s.Winner != "" || s.PlayerTurn != "computer" {
		return
	}
	take := BestResponse(s.Variant, s.Remaining, s.K)
	if take <= 0 {
		take = 1
	}
	if take > s.Remaining {
		take = s.Remaining
	}
	s.Remaining -= take

	if s.Remaining == 0 {
		if s.Variant == Normal {
			s.Winner = "computer"
		} else {
			s.Winner = "human"
		}
		return
	}
	s.PlayerTurn = "human"
}
