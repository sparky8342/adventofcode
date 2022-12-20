package main

import (
	"testing"
)

var data = []int{1, 2, -3, 3, -2, 0, 4}

func TestSum(t *testing.T) {
	list := create_list(data)
	list.move_nums()

	got_sum := list.find_sum()
	want_sum := 3

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}
}

func TestSumWithKey(t *testing.T) {
	list := create_list(data)
	list.apply_key()
	for i := 0; i < 10; i++ {
		list.move_nums()
	}

	got_sum := list.find_sum()
	want_sum := 1623178306

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}
}
