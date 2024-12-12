package day12

import "testing"

func Test1(t *testing.T) {
	data := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}

	got := price(data)
	want := 140

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"OOOOO",
		"OXOXO",
		"OOOOO",
		"OXOXO",
		"OOOOO",
	}

	got := price(data)
	want := 772

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	data := []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}

	got := price(data)
	want := 1930

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
