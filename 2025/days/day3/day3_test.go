package day3

import "testing"

var banks = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func Test1(t *testing.T) {
	got := total_joltage(banks)
	want := 357

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
