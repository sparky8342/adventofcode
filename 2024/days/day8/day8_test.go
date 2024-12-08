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
	got := antinodes(grid)
	want := 14

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
