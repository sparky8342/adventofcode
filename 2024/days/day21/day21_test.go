package day21

import "testing"

var data = []string{
	"029A",
	"980A",
	"179A",
	"456A",
	"379A",
}

func Test1(t *testing.T) {
	got := find_sequences(data)
	want := 126384

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
