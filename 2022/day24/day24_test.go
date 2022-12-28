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

func TestPath(t *testing.T) {
	grid, cycle, start, end := parse_data(strings.Split(string(data), "\n"))
	got_time := bfs(&grid, cycle, start, end, 0)
	want_time := 18

	if got_time != want_time {
		t.Errorf("got %d, wanted %d", got_time, want_time)
	}
}

func TestBackAgainPath(t *testing.T) {
	grid, cycle, start, end := parse_data(strings.Split(string(data), "\n"))
	time := bfs(&grid, cycle, start, end, 0)
	time2 := bfs(&grid, cycle, end, start, time)
	got_time := bfs(&grid, cycle, start, end, time2)
	want_time := 54

	if got_time != want_time {
		t.Errorf("got %d, wanted %d", got_time, want_time)
	}
}
