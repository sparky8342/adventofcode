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

func TestMinPath(t *testing.T) {
	nodes, _, end := parse_data(data)

	got_min_distance := 99999
	for pos, node := range nodes {
		if node.elevation == 'a' {
			distance := bfs(nodes, pos, end)
			if distance > 0 && distance < got_min_distance {
				got_min_distance = distance
			}
		}
	}

	want_min_distance := 29

	if got_min_distance != want_min_distance {
		t.Errorf("got %d, wanted %d", got_min_distance, want_min_distance)
	}
}
