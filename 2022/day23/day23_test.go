package main

import (
	"strings"
	"testing"
)

const data = `##
#.
..
##`

const data2 = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func TestExample1(t *testing.T) {
	grove := parse_data(strings.Split(data, "\n"))

	for i := 0; i < 4; i++ {
		grove.move_elves()
		grove.rotate_check_order()
	}

	got_empty := grove.print_grove()
	want_empty := 25

	if got_empty != want_empty {
		t.Errorf("got %d, wanted %d", got_empty, want_empty)
	}
}

func TestExample2(t *testing.T) {
	grove := parse_data(strings.Split(data2, "\n"))

	for i := 0; i < 10; i++ {
		grove.move_elves()
		grove.rotate_check_order()
	}

	got_empty := grove.print_grove()
	want_empty := 110

	if got_empty != want_empty {
		t.Errorf("got %d, wanted %d", got_empty, want_empty)
	}
}
