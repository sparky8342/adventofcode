package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"19, 13, 30 @ -2,  1, -2",
		"18, 19, 22 @ -1, -1, -2",
		"20, 25, 34 @ -2, -2, -4",
		"12, 31, 28 @ -1, -2, -1",
		"20, 19, 15 @  1, -5, -3",
	}

	hailstones := parse_data(data)
	got_intersections := find_intersections(hailstones, 7, 27)
	want_intersections := 2

	if got_intersections != want_intersections {
		t.Errorf("got %d, wanted %d", got_intersections, want_intersections)
	}
}
