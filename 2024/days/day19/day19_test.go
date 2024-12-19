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
	got_possible, got_ways := possible_patterns(data)
	want_possible := 6

	if want_possible != got_possible {
		t.Errorf("got %d, wanted %d", got_possible, want_possible)
	}

	want_ways := 16

	if want_ways!= got_ways {
		t.Errorf("got %d, wanted %d", got_ways, want_ways)
	}
}
