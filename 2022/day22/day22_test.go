package main

import (
	"strings"
	"testing"
)

const data = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

func TestRoot(t *testing.T) {
	grid := parse_data(strings.Split(data, "\n"))
	got_password := grid.walk()
	want_password := 6032

	if got_password != want_password {
		t.Errorf("got %d, wanted %d", got_password, want_password)
	}
}
