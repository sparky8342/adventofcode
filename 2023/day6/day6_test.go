package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	races, combined_race := parse_data(data)

	got_race_wins_product := race_wins_product(races)
	want_race_wins_product := 288

	if got_race_wins_product != want_race_wins_product {
		t.Errorf("got %d, wanted %d", got_race_wins_product, want_race_wins_product)
	}

	got_combined_race_wins := race_wins(combined_race)
	want_combined_race_wins := 71503

	if got_combined_race_wins != want_combined_race_wins {
		t.Errorf("got %d, wanted %d", got_combined_race_wins, want_combined_race_wins)
	}
}
