package main

import "testing"

func Test(t *testing.T) {
	data := [][]byte{
		[]byte("-L|F7"),
		[]byte("7S-7|"),
		[]byte("L|7||"),
		[]byte("-L-J|"),
		[]byte("L|-JF"),
	}

	got_distance, _ := path_info(data)
	want_distance := 4

	if got_distance != want_distance {
		t.Errorf("got %d, wanted %d", got_distance, want_distance)
	}

	data = [][]byte{
		[]byte("..........."),
		[]byte(".S-------7."),
		[]byte(".|F-----7|."),
		[]byte(".||.....||."),
		[]byte(".||.....||."),
		[]byte(".|L-7.F-J|."),
		[]byte(".|..|.|..|."),
		[]byte(".L--J.L--J."),
		[]byte("..........."),
	}

	_, got_inside := path_info(data)
	want_inside := 4

	if got_inside != want_inside {
		t.Errorf("got %d, wanted %d", got_inside, want_inside)
	}

	data = [][]byte{
		[]byte(".........."),
		[]byte(".S------7."),
		[]byte(".|F----7|."),
		[]byte(".||....||."),
		[]byte(".||....||."),
		[]byte(".|L-7F-J|."),
		[]byte(".|..||..|."),
		[]byte(".L--JL--J."),
		[]byte(".........."),
	}

	_, got_inside = path_info(data)
	want_inside = 4

	if got_inside != want_inside {
		t.Errorf("got %d, wanted %d", got_inside, want_inside)
	}

	data = [][]byte{
		[]byte(".F----7F7F7F7F-7...."),
		[]byte(".|F--7||||||||FJ...."),
		[]byte(".||.FJ||||||||L7...."),
		[]byte("FJL7L7LJLJ||LJ.L-7.."),
		[]byte("L--J.L7...LJS7F-7L7."),
		[]byte("....F-J..F7FJ|L7L7L7"),
		[]byte("....L7.F7||L7|.L7L7|"),
		[]byte(".....|FJLJ|FJ|F7|.LJ"),
		[]byte("....FJL-7.||.||||..."),
		[]byte("....L---J.LJ.LJLJ..."),
	}

	_, got_inside = path_info(data)
	want_inside = 8

	if got_inside != want_inside {
		t.Errorf("got %d, wanted %d", got_inside, want_inside)
	}
}
