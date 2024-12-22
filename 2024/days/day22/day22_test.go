package day22

import "testing"

func Test1(t *testing.T) {
	nums := []int{1, 10, 100, 2024}

	got := sequences(nums)
	want := 37327623

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	nums := []int{1, 2, 3, 2024}

	got := best_sequence(nums)
	want := 23

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
