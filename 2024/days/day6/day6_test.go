package day6

import "testing"

var data = []string{
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

var grid, start_x, start_y = parse_data(data)

func Test1(t *testing.T) {
	_, got := walk(grid, start_x, start_y)
	want := 41

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := obstructions(grid, start_x, start_y)
	want := 6

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
