package day4

import "testing"

var grid = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func Test1(t *testing.T) {
	got := count_xmas(grid)
	want := 18

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := count_x_mas(grid)
	want := 9

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
