package main

import "testing"

func Test(t *testing.T) {
	data := load_data("input_test.txt")

	got_most, got_top_three := most_calories(data)
	want_most, want_top_three := 24000, 45000

	if got_most != want_most {
		t.Errorf("got %d, wanted %d", got_most, want_most)
	}

	if got_top_three != want_top_three {
		t.Errorf("got %d, wanted %d", got_top_three, want_top_three)
	}
}
