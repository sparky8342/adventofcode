package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	}

	got_distance := distance(data)
	want_distance := 4

	if got_distance != want_distance {
		t.Errorf("got %d, wanted %d", got_distance, want_distance)
	}
}
