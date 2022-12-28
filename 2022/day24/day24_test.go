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
	grids, cycle, start, end := parse_data(strings.Split(string(data), "\n"))
	got_time := bfs(grids, cycle, start, end, 0)
	want_time := 18

	if got_time != want_time {
		t.Errorf("got %d, wanted %d", got_time, want_time)
	}
}

func TestBackAgainPath(t *testing.T) {
	grids, cycle, start, end := parse_data(strings.Split(string(data), "\n"))
	time := bfs(grids, cycle, start, end, 0)
	time2 := bfs(grids, cycle, end, start, time)
	got_time := bfs(grids, cycle, start, end, time2)
	want_time := 54

	if got_time != want_time {
		t.Errorf("got %d, wanted %d", got_time, want_time)
	}
}
