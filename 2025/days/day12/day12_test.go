package day12

import "testing"

var data = []string{
	"0:",
	"###",
	"##.",
	"##.",
	"",
	"1:",
	"###",
	"##.",
	".##",
	"",
	"2:",
	".##",
	"###",
	"##.",
	"",
	"3:",
	"##.",
	"###",
	"##.",
	"",
	"4:",
	"###",
	"#..",
	"###",
	"",
	"5:",
	"###",
	".#.",
	"###",
	"",
	"4x4: 0 0 0 0 2 0",
	"12x5: 1 0 1 0 2 2",
	"12x5: 1 0 1 0 3 2",
}

func Test1(t *testing.T) {
	shapes, regions := parse_data(data)

	got := place_shapes(shapes, regions)
	want := 2

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
