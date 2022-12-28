package main

import (
	"strings"
	"testing"
)

const data = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

func Test(t *testing.T) {
	grid := parse_data(strings.Split(string(data), "\n"))
	got_time := bfs(&grid)
	want_time := 18

	if got_time != want_time {
		t.Errorf("got %d, wanted %d", got_time, want_time)
	}
}
