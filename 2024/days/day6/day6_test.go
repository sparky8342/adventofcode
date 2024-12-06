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

var grid, guard = parse_data(data)

func Test1(t *testing.T) {
	_, visited := walk(grid, guard)
	got := len(visited)
	want := 41

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	_, visited := walk(grid, guard)
	got := obstructions(grid, guard, visited)
	want := 6

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
