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
	got_energized := energize(grid)
	want_energized := 46

	if got_energized != want_energized {
		t.Errorf("got %d, wanted %d", got_energized, want_energized)
	}
}
