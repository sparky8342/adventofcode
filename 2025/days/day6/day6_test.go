package day6

import "testing"

var data = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func Test1(t *testing.T) {
	nums, operators := parse_data_part1(data)

	got := calculate(nums, operators)
	want := 4277556

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	nums, operators := parse_data_part2(data)

	got := calculate(nums, operators)
	want := 3263827

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
