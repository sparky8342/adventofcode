package day22

import "testing"

var nums = []int{1, 10, 100, 2024}

func Test1(t *testing.T) {
	got := sequences(nums, 2000)
	want := 37327623

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
