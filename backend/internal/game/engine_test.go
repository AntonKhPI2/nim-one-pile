package game

import "testing"

func TestNewGame_PicksFirstTurnOptimally_Normal(t *testing.T) {
	cfg := Config{Variant: Normal, N: 21, K: 3}
	s, err := NewGame(cfg, "test")
	if err != nil {
		t.Fatal(err)
	}
	if s.PlayerTurn != "human" {
		t.Fatalf("computer should have moved first inside NewGame; expected PlayerTurn=human, got %s", s.PlayerTurn)
	}
	if s.Remaining%4 != 0 {
		t.Fatalf("computer should leave multiple of 4, remaining=%d", s.Remaining)
	}
}

func TestApplyHumanMove_And_ComputerMove_Normal(t *testing.T) {
	cfg := Config{Variant: Normal, N: 10, K: 3}
	s, err := NewGame(cfg, "g1")
	if err != nil {
		t.Fatal(err)
	}

	if s.PlayerTurn != "human" {
		ApplyComputerMove(s)
	}
	if err := ApplyHumanMove(s, 1); err != nil {
		t.Fatalf("human move: %v", err)
	}
	if s.PlayerTurn != "computer" {
		t.Fatalf("expected computer turn, got %s", s.PlayerTurn)
	}
	ApplyComputerMove(s)
	if s.PlayerTurn != "human" && s.Winner == "" {
		t.Fatalf("expected human turn (or someone won). turn=%s winner=%s", s.PlayerTurn, s.Winner)
	}
}

func TestWinConditions_Normal(t *testing.T) {
	cfg := Config{Variant: Normal, N: 1, K: 3}
	s, err := NewGame(cfg, "g2")
	if err != nil {
		t.Fatal(err)
	}

	if s.PlayerTurn != "human" {
		ApplyComputerMove(s)
	}

	if s.Winner == "" && s.PlayerTurn == "human" {
		if err := ApplyHumanMove(s, 1); err != nil {
			t.Fatal(err)
		}
	}

	if s.Winner != "human" && s.Winner != "computer" {
		t.Fatalf("expected a winner in normal with N=1, got none")
	}
}

func TestWinConditions_Misere(t *testing.T) {
	cfg := Config{Variant: Misere, N: 1, K: 3}
	s, err := NewGame(cfg, "g3")
	if err != nil {
		t.Fatal(err)
	}

	if s.PlayerTurn == "human" {
		if err := ApplyHumanMove(s, 1); err != nil {
			t.Fatal(err)
		}
		if s.Winner != "computer" {
			t.Fatalf("expected computer winner in misere when human takes last, got %s", s.Winner)
		}
	} else {
		ApplyComputerMove(s)
		if s.Winner != "human" {
			t.Fatalf("expected human winner in misere when computer takes last, got %s", s.Winner)
		}
	}
}

func TestInvalidMoves(t *testing.T) {
	cfg := Config{Variant: Normal, N: 5, K: 3}
	s, err := NewGame(cfg, "g4")
	if err != nil {
		t.Fatal(err)
	}
	if s.PlayerTurn != "human" {
		ApplyComputerMove(s)
	}
	if err := ApplyHumanMove(s, 0); err == nil {
		t.Fatalf("expected error for take=0")
	}
	if err := ApplyHumanMove(s, 4); err == nil {
		t.Fatalf("expected error for take>k")
	}
	if err := ApplyHumanMove(s, 3); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
