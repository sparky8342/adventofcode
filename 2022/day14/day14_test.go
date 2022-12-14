package main

import (
	"testing"
)

var data = []string{"498,4 -> 498,6 -> 496,6", "503,4 -> 502,4 -> 502,9 -> 494,9"}

func TestSand(t *testing.T) {
	cave := NewCave()
	cave.draw_points(data)
	got_sand := cave.drop_sand()
	want_sand := 24

	if got_sand != want_sand {
		t.Errorf("got %d, wanted %d", got_sand, want_sand)
	}
}
