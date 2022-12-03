package main

import "testing"

var data = []string{"A Y", "B X", "C Z"}

func TestGetScore(t *testing.T) {
	got_score := get_score(data)
	want_score := 15

	if got_score != want_score {
		t.Errorf("got %d, wanted %d", got_score, want_score)
	}
}

func TestGetScoreWithStrategy(t *testing.T) {
	got_score := get_score_with_strategy(data)
	want_score := 12

	if got_score != want_score {
		t.Errorf("got %d, wanted %d", got_score, want_score)
	}
}
