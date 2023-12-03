package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	got_sum, got_ratio := sum_and_ratio(data)
	want_sum := 4361
	want_ratio := 467835

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}

	if got_ratio != want_ratio {
		t.Errorf("got %d, wanted %d", got_ratio, want_ratio)
	}
}
