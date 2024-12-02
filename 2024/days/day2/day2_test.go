package day2

import "testing"

func Test1(t *testing.T) {
	nums := [][]int{
		[]int{7, 6, 4, 2, 1},
		[]int{1, 2, 7, 8, 9},
		[]int{9, 7, 6, 2, 1},
		[]int{1, 3, 2, 4, 5},
		[]int{8, 6, 4, 4, 1},
		[]int{1, 3, 6, 7, 9},
	}

	got := count_safe(nums)
	want := 2

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	nums := [][]int{
		[]int{7, 6, 4, 2, 1},
		[]int{1, 2, 7, 8, 9},
		[]int{9, 7, 6, 2, 1},
		[]int{1, 3, 2, 4, 5},
		[]int{8, 6, 4, 4, 1},
		[]int{1, 3, 6, 7, 9},
	}

	got := count_safe_with_tolerance(nums)
	want := 4

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
