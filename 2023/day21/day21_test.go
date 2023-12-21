package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"...........",
		".....###.#.",
		".###.##..#.",
		"..#.#...#..",
		"....#.#....",
		".##..S####.",
		".##..#...#.",
		".......##..",
		".##.#.####.",
		".##..##.##.",
		"...........",
	}

	got_plots := plots(data, 6)
	want_plots := 16

	if got_plots != want_plots {
		t.Errorf("got %d, wanted %d", got_plots, want_plots)
	}
}
