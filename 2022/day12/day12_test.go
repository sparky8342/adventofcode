package main

import (
	"testing"
)

var data = []string{"Sabqponm", "abcryxxl", "accszExk", "acctuvwj", "abdefghi"}

func TestPath(t *testing.T) {
	nodes, start, end := parse_data(data)
	got_distance := bfs(nodes, start, end)
	want_distance := 31

	if got_distance != want_distance {
		t.Errorf("got %d, wanted %d", got_distance, want_distance)
	}
}
