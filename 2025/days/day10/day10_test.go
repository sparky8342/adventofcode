package day10

import "testing"

var data = []string{
	"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
	"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
}

func Test1(t *testing.T) {
	machines := parse_data(data)

	got := presses_needed(machines)
	want := 7

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
