package main

import (
	"testing"
)

var data = []string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"}

func TestTailVisited(t *testing.T) {
	commands := parse_data(data)
	got_squares := run_commands_part1(commands)
	want_squares := 13

	if got_squares != want_squares {
		t.Errorf("got %d, wanted %d", got_squares, want_squares)
	}
}

func TestTailVisitedLongRope(t *testing.T) {
	commands := parse_data(data)
	got_squares := run_commands_part2(commands)
	want_squares := 1

	if got_squares != want_squares {
		t.Errorf("got %d, wanted %d", got_squares, want_squares)
	}
}
