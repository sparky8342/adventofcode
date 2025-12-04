package day4

import "testing"

var grid = []string{
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
	got := rolls_reachable(grid)
	want := 13

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
