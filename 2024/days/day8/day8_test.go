package day8

import "testing"

var grid = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func Test1(t *testing.T) {
	got := antinodes(grid, false)
	want := 14

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"T.........",
		"...T......",
		".T........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
	}

	got := antinodes(grid, true)
	want := 9

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	got := antinodes(grid, true)
	want := 34

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
