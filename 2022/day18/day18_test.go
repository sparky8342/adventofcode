package main

import (
	"strings"
	"testing"
)

const data = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestSides(t *testing.T) {
	droplet := parse_data(strings.Split(data, "\n"))
	got_sides := count_sides(droplet)
	want_sides := 64

	if got_sides != want_sides {
		t.Errorf("got %d, wanted %d", got_sides, want_sides)
	}
}

func TestArea(t *testing.T) {
	droplet := parse_data(strings.Split(data, "\n"))
	sides := count_sides(droplet)
	inside_sides := count_holes(droplet)
	got_area := sides - inside_sides
	want_area := 58

	if got_area != want_area {
		t.Errorf("got %d, wanted %d", got_area, want_area)
	}
}
