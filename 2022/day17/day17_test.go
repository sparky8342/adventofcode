package main

import (
	"testing"
)

const data = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func TestPieces(t *testing.T) {
	chamber := NewChamber(data)

	for i := 0; i < 2022; i++ {
		chamber.drop_piece()
	}

	got_height := chamber.top_y + 1
	want_height := 3068

	if got_height != want_height {
		t.Errorf("got %d, wanted %d", got_height, want_height)
	}
}

func TestLargePieces(t *testing.T) {
	chamber := NewChamber(data)
	got_height := part2(chamber)
	want_height := 1514285714288

	if got_height != want_height {
		t.Errorf("got %d, wanted %d", got_height, want_height)
	}
}
