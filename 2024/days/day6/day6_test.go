package day6

import "testing"

var grid = []string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

func Test1(t *testing.T) {
	got := walk(grid)
	want := 41

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
