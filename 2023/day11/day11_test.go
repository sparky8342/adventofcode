package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	universe := parse_data(data)
	universe.expand()

	got_sum := universe.path_sum()
	want_sum := 374

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}
}
