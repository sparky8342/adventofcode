package day1

import "testing"

var data = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

func Test1(t *testing.T) {
	got, _ := turn_dial(data)
	want := 3

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	_, got := turn_dial(data)
	want := 6

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
