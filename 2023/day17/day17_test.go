package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"2413432311323",
		"3215453535623",
		"3255245654254",
		"3446585845452",
		"4546657867536",
		"1438598798454",
		"4457876987766",
		"3637877979653",
		"4654967986887",
		"4564679986453",
		"1224686865563",
		"2546548887735",
		"4322674655533",
	}

	got_heat_loss := find_path(data, false)
	want_heat_loss := 102

	if got_heat_loss != want_heat_loss {
		t.Errorf("got %d, wanted %d", got_heat_loss, want_heat_loss)
	}

	got_heat_loss = find_path(data, true)
	want_heat_loss = 94

	if got_heat_loss != want_heat_loss {
		t.Errorf("got %d, wanted %d", got_heat_loss, want_heat_loss)
	}
}
