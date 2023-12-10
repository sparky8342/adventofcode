package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	}

	got_distance, _ := path_info(data)
	want_distance := 4

	if got_distance != want_distance {
		t.Errorf("got %d, wanted %d", got_distance, want_distance)
	}

	data = []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}

	_, got_inside := path_info(data)
	want_inside := 4

	if got_inside != want_inside {
		t.Errorf("got %d, wanted %d", got_inside, want_inside)
	}

	data = []string{
		"..........",
		".S------7.",
		".|F----7|.",
		".||....||.",
		".||....||.",
		".|L-7F-J|.",
		".|..||..|.",
		".L--JL--J.",
		"..........",
	}

	_, got_inside = path_info(data)
	want_inside = 4

	if got_inside != want_inside {
		t.Errorf("got %d, wanted %d", got_inside, want_inside)
	}

	data = []string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}

	_, got_inside = path_info(data)
	want_inside = 8

	if got_inside != want_inside {
		t.Errorf("got %d, wanted %d", got_inside, want_inside)
	}
}
