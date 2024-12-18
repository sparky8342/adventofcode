package day18

import "testing"

var data = []string{
	"5,4",
	"4,2",
	"4,5",
	"3,0",
	"2,1",
	"6,3",
	"2,4",
	"1,5",
	"0,6",
	"3,3",
	"2,6",
	"5,1",
	"1,2",
	"5,5",
	"2,5",
	"6,5",
	"1,4",
	"0,4",
	"6,4",
	"1,1",
	"6,1",
	"1,0",
	"0,5",
	"1,6",
	"2,0",
}

func Test1(t *testing.T) {
	blocks := parse_data(data, 12)

	got := find_path(7, blocks)
	want := 22

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	blocks := parse_data(data, 12)

	got := find_blocked_path(data, 7, blocks, 12)
	want := "6,1"

	if want != got {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
