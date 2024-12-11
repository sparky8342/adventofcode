package day11

import "testing"

var stones = []int{125, 17}

func Test1(t *testing.T) {
	got := blink(stones, 25)
	want := 55312

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
