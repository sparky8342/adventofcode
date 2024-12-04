package day1

import "testing"

var nums = [][]int{
	[]int{3, 4, 2, 1, 3, 3},
	[]int{4, 3, 5, 3, 9, 3},
}

func Test1(t *testing.T) {
	got := distance(nums[0], nums[1])
	want := 11

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := similarity(nums[0], nums[1])
	want := 31

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
