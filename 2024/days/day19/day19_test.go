package day19

import "testing"

var data = []string{
	"r, wr, b, g, bwu, rb, gb, br",
	"",
	"brwrr",
	"bggr",
	"gbbr",
	"rrbgbr",
	"ubwu",
	"bwurrg",
	"brgr",
	"bbrgwb",
}

func Test1(t *testing.T) {
	got := possible_patterns(data)
	want := 6

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
