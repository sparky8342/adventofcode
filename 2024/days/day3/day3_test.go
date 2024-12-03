package day3

import "testing"

func Test1(t *testing.T) {
	data := []string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}

	got := find_valid(data, false)
	want := 161

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}

	got := find_valid(data, true)
	want := 48

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
