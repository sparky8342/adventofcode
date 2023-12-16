package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		".|...\\....",
		"|.-.\\.....",
		".....|-...",
		"........|.",
		"..........",
		".........\\",
		"..../.\\\\..",
		".-.-/..|..",
		".|....-|.\\",
		"..//.|....",
	}

	grid := parse_data(data)
	got_energized := energize_top_left(grid)
	want_energized := 46

	if got_energized != want_energized {
		t.Errorf("got %d, wanted %d", got_energized, want_energized)
	}

	got_max_energy := max_energy(grid)
	want_max_energy := 51

	if got_max_energy != want_max_energy {
		t.Errorf("got %d, wanted %d", got_max_energy, want_max_energy)
	}
}
