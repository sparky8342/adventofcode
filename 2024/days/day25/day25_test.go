package day25

import "testing"

var data = [][]string{
	[]string{
		"#####",
		".####",
		".####",
		".####",
		".#.#.",
		".#...",
		".....",
	},
	[]string{
		"#####",
		"##.##",
		".#.##",
		"...##",
		"...#.",
		"...#.",
		".....",
	},
	[]string{
		".....",
		"#....",
		"#....",
		"#...#",
		"#.#.#",
		"#.###",
		"#####",
	},
	[]string{
		".....",
		".....",
		"#.#..",
		"###..",
		"###.#",
		"###.#",
		"#####",
	},
	[]string{
		".....",
		".....",
		".....",
		"#....",
		"#.#..",
		"#.#.#",
		"#####",
	},
}

func Test1(t *testing.T) {
	keys, locks := parse_data(data)

	got := fit_combinations(keys, locks)
	want := 3

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
