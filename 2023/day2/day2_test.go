package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	got_possible, got_power_total := possible(data)
	want_possible := 8
	want_power_total := 2286

	if got_possible != want_possible {
		t.Errorf("got %d, wanted %d", got_possible, want_possible)
	}

	if got_power_total != want_power_total {
		t.Errorf("got %d, wanted %d", got_power_total, want_power_total)
	}
}
