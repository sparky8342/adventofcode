package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}

	grid := parse_data(data)
	tilt_north(grid)
	got_load := get_load(grid)
	want_load := 136

	if got_load != want_load {
		t.Errorf("got %d, wanted %d", got_load, want_load)
	}

	spins(grid, 1000000000)

	got_load = get_load(grid)
	want_load = 64

	if got_load != want_load {
		t.Errorf("got %d, wanted %d", got_load, want_load)
	}
}
