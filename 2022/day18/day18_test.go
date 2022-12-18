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
	cubes := parse_data(strings.Split(data, "\n"))
	got_sides := count_sides(cubes)
	want_sides := 64

	if got_sides != want_sides {
		t.Errorf("got %d, wanted %d", got_sides, want_sides)
	}
}
