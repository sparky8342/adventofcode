package day7

import "testing"

var grid = []string{
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}

func Test1(t *testing.T) {
	got, _ := follow_beam(grid)
	want := 21

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	_, got := follow_beam(grid)
	want := 40

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
