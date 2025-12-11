package day11

import "testing"

var data = []string{
	"aaa: you hhh",
	"you: bbb ccc",
	"bbb: ddd eee",
	"ccc: ddd eee fff",
	"ddd: ggg",
	"eee: out",
	"fff: out",
	"ggg: out",
	"hhh: ccc fff iii",
	"iii: out",
}

func Test1(t *testing.T) {
	you := parse_data(data)

	got := count_paths(you)
	want := 5

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
