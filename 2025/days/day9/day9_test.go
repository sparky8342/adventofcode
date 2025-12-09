package day9

import "testing"

var data = []string{
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
}

func Test1(t *testing.T) {
	tiles := parse_data(data)

	got := largest_rectangle(tiles)
	want := 50

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
