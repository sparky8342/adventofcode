package day4

import "testing"

var data = []string{
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

func Test1(t *testing.T) {
	grid := parse_data(data)

	got := rolls_reachable(grid)
	want := 13

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := parse_data(data)

	got := remove_rolls(grid)
	want := 43

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
