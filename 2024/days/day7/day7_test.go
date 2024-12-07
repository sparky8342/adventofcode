package day7

import "testing"

var data = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

var equations = parse_data(data)

func Test1(t *testing.T) {
	got := total_calibration(equations, false)
	want := 3749

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := total_calibration(equations, true)
	want := 11387

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
