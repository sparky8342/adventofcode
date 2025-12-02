package day2

import "testing"

var data = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func Test1(t *testing.T) {
	ranges := parse_data(data)

	got := total_invalid(ranges, 1)
	want := 1227775554

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	ranges := parse_data(data)

	got := total_invalid(ranges, 2)
	want := 4174379265

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
