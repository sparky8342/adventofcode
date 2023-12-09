package main

import "testing"

func Test(t *testing.T) {
	nums := [][]int{
		[]int{0, 3, 6, 9, 12, 15},
		[]int{1, 3, 6, 10, 15, 21},
		[]int{10, 13, 16, 21, 30, 45},
	}

	got_next_sum, got_previous_sum := calculate_sum(nums)
	want_next_sum := 114
	want_previous_sum := 2

	if got_next_sum != want_next_sum {
		t.Errorf("got %d, wanted %d", got_next_sum, want_next_sum)
	}

	if got_previous_sum != want_previous_sum {
		t.Errorf("got %d, wanted %d", got_previous_sum, want_previous_sum)
	}
}
