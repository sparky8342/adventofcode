package day16

import "testing"

func Test1(t *testing.T) {
	grid := []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}

	got_time, got_squares := score(grid)
	want_time := 7036

	if want_time != got_time {
		t.Errorf("got %d, wanted %d", got_time, want_time)
	}

	want_squares := 45
	if want_squares != got_squares {
		t.Errorf("got %d, wanted %d", got_squares, want_squares)
	}

}

func Test2(t *testing.T) {
	grid := []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}

	got_time, got_squares := score(grid)
	want_time := 11048

	if want_time != got_time {
		t.Errorf("got %d, wanted %d", got_time, want_time)
	}

	want_squares := 64
	if want_squares != got_squares {
		t.Errorf("got %d, wanted %d", got_squares, want_squares)
	}
}
