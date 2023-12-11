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

	got_sum := universe.path_sum(2)
	want_sum := 374

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}

	got_sum = universe.path_sum(10)
	want_sum = 1030

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}

	got_sum = universe.path_sum(100)
	want_sum = 8410

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}
}
