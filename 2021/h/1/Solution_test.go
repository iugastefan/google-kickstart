package main

import "testing"

func TestSolve(t *testing.T) {
	type Input struct {
		S string
		F string
	}
	type Case struct {
		In   Input
		Want int
	}
	cases := []Case{
		{Input{S: "abcd", F: "a"}, 6},
		{Input{S: "pppp", F: "p"}, 0},
		{Input{S: "pqrst", F: "ou"}, 9},
		{Input{S: "abd", F: "abd"}, 0},
		{Input{S: "aaaaaaaaaaaaaaab", F: "aceg"}, 1},
	}
	for i, c := range cases {
		got := solve(c.In.S, c.In.F)
		if got != c.Want {
			t.Errorf("%d: got %d, want %d", i, got, c.Want)
		}
	}
}

func TestDistance(t *testing.T) {
	type Input struct {
		A rune
		B rune
	}
	type Case struct {
		In   Input
		Want int
	}
	cases := []Case{
		{Input{A: 'a', B: 'a'}, 0},
		{Input{A: 'a', B: 'b'}, 1},
		{Input{A: 'a', B: 'c'}, 2},
		{Input{A: 'a', B: 'd'}, 3},
		{Input{A: 'a', B: 'e'}, 4},
		{Input{A: 'a', B: 'f'}, 5},
		{Input{A: 'a', B: 'g'}, 6},
		{Input{A: 'a', B: 'h'}, 7},
		{Input{A: 'a', B: 'i'}, 8},
		{Input{A: 'a', B: 'j'}, 9},
		{Input{A: 'a', B: 'k'}, 10},
		{Input{A: 'a', B: 'l'}, 11},
		{Input{A: 'a', B: 'm'}, 12},
		{Input{A: 'a', B: 'n'}, 13},
		{Input{A: 'a', B: 'o'}, 12},
		{Input{A: 'a', B: 'p'}, 11},
		{Input{A: 'a', B: 'q'}, 10},
		{Input{A: 'a', B: 'r'}, 9},
		{Input{A: 'a', B: 's'}, 8},
		{Input{A: 'a', B: 't'}, 7},
		{Input{A: 'a', B: 'u'}, 6},
		{Input{A: 'a', B: 'v'}, 5},
		{Input{A: 'a', B: 'w'}, 4},
		{Input{A: 'a', B: 'x'}, 3},
		{Input{A: 'a', B: 'y'}, 2},
		{Input{A: 'a', B: 'z'}, 1},
	}
	for i, c := range cases {
		got := distance(c.In.A, c.In.B)
		if got != c.Want {
			t.Errorf("%d: got %d, want %d", i, got, c.Want)
		}
	}
}
