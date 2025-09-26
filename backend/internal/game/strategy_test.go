package game

import "testing"

func TestIsLosingStart_Normal(t *testing.T) {
	for k := 1; k <= 5; k++ {
		m := k + 1
		for N := 1; N <= 60; N++ {
			want := (N % m) == 0
			if got := isLosingStart(Normal, N, k); got != want {
				t.Fatalf("normal: k=%d N=%d got=%v want=%v", k, N, got, want)
			}
		}
	}
}

func TestIsLosingStart_Misere(t *testing.T) {
	for k := 1; k <= 5; k++ {
		m := k + 1
		for N := 1; N <= 60; N++ {
			want := (N % m) == 1
			if got := isLosingStart(Misere, N, k); got != want {
				t.Fatalf("misere: k=%d N=%d got=%v want=%v", k, N, got, want)
			}
		}
	}
}

func TestBestResponse_Normal(t *testing.T) {

	k := 3
	cases := []struct {
		remaining int
		want      int
	}{
		{remaining: 21, want: 1},
		{remaining: 22, want: 2},
		{remaining: 23, want: 3},
		{remaining: 24, want: 0},
		{remaining: 1, want: 1},
	}
	for _, c := range cases {
		if got := BestResponse(Normal, c.remaining, k); got != c.want {
			t.Errorf("normal: remaining=%d got=%d want=%d", c.remaining, got, c.want)
		}
	}
}

func TestBestResponse_Misere(t *testing.T) {

	k := 3
	type tc struct {
		remaining int
		want      int
	}
	tests := []tc{
		{remaining: 21, want: 0},
		{remaining: 22, want: 1},
		{remaining: 23, want: 2},
		{remaining: 24, want: 3},
		{remaining: 1, want: 0},
	}
	for _, c := range tests {
		if got := BestResponse(Misere, c.remaining, k); got != c.want {
			t.Errorf("misere: remaining=%d got=%d want=%d", c.remaining, got, c.want)
		}
	}
}
