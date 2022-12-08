package main

import (
	"testing"
)

var data = []string{"30373", "25512", "65332", "33549", "35390"}

var grid Grid

func init() {
	grid = make_grid(data)
}

func TestGetVisible(t *testing.T) {
	got_trees := visible(grid)
	want_trees := 21

	if got_trees != want_trees {
		t.Errorf("got %d, wanted %d", got_trees, want_trees)
	}
}

func TestGetScore(t *testing.T) {
	got_score := highest_scenic_score(grid)
	want_score := 8

	if got_score != want_score {
		t.Errorf("got %d, wanted %d", got_score, want_score)
	}
}
