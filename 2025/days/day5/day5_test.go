package day5

import "testing"

var data = []string{
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

func Test1(t *testing.T) {
	ranges, ingredients := parse_data(data)

	got := count_fresh(ranges, ingredients)
	want := 3

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	ranges, _ := parse_data(data)

	got := fresh_ids(ranges)
	want := 14

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
