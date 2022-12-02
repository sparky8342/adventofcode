package main

import "testing"

func Test(t *testing.T) {
	data := []string{"A Y", "B X", "C Z"}

	got_score := get_score(data)
	want_score := 15

	if got_score != want_score {
		t.Errorf("got %d, wanted %d", got_score, want_score)
	}

	got_score = get_score_with_strategy(data)
	want_score = 12

	if got_score != want_score {
		t.Errorf("got %d, wanted %d", got_score, want_score)
	}
}
