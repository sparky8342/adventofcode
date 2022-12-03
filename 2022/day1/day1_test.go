package main

import "testing"

func Test(t *testing.T) {
	data := "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000\n"

	got_most, got_top_three := most_calories(data)
	want_most, want_top_three := 24000, 45000

	if got_most != want_most {
		t.Errorf("got %d, wanted %d", got_most, want_most)
	}

	if got_top_three != want_top_three {
		t.Errorf("got %d, wanted %d", got_top_three, want_top_three)
	}
}
