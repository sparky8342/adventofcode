package main

import (
	"math/big"
	"testing"
)

func Test(t *testing.T) {
	data := []string{
		"R 6 (#70c710)",
		"D 5 (#0dc571)",
		"L 2 (#5713f0)",
		"D 2 (#d2c081)",
		"R 2 (#59c680)",
		"D 2 (#411b91)",
		"L 5 (#8ceee2)",
		"U 2 (#caa173)",
		"L 1 (#1b58a2)",
		"U 2 (#caa171)",
		"R 2 (#7807d2)",
		"U 3 (#a77fa3)",
		"L 2 (#015232)",
		"U 2 (#7a21e3)",
	}

	got_size := get_size(data, false)
	want_size := big.NewInt(int64(62))

	if got_size.Cmp(want_size) != 0 {
		t.Errorf("got %f, wanted %f", got_size, want_size)
	}

	got_size = get_size(data, true)
	want_size = big.NewInt(int64(952408144115))

	if got_size.Cmp(want_size) != 0 {
		t.Errorf("got %f, wanted %f", got_size, want_size)
	}

}
