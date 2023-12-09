package main

import "testing"

func Test(t *testing.T) {
	nums := [][]int{
		[]int{0, 3, 6, 9, 12, 15},
		[]int{1, 3, 6, 10, 15, 21},
		[]int{10, 13, 16, 21, 30, 45},
	}

	got_sum := calculate_sum(nums)
	want_sum := 114

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}
}
