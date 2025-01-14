package day20

import "testing"

var grid = []string{
	"###############",
	"#...#...#.....#",
	"#.#.#.#.#.###.#",
	"#S#...#.#.#...#",
	"#######.#.#.###",
	"#######.#.#...#",
	"#######.#.###.#",
	"###..E#...#...#",
	"###.#######.###",
	"#...###...#...#",
	"#.#####.#.###.#",
	"#.#...#.#.#...#",
	"#.#.#.#.#.#.###",
	"#...#...#...###",
	"###############",
}

func Test1(t *testing.T) {
	got := find_cheats(grid, 1, 2)
	want := 44

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := find_cheats(grid, 50, 20)
	want := 285

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
