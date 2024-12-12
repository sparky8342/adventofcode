package day12

import "testing"

func Test1(t *testing.T) {
	data := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}

	got, _ := price(data)
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

	got, _ := price(data)
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

	got, _ := price(data)
	want := 1930

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	data := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}

	_, got := price(data)
	want := 80

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	data := []string{
		"OOOOO",
		"OXOXO",
		"OOOOO",
		"OXOXO",
		"OOOOO",
	}

	_, got := price(data)
	want := 436

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test6(t *testing.T) {
	data := []string{
		"EEEEE",
		"EXXXX",
		"EEEEE",
		"EXXXX",
		"EEEEE",
	}

	_, got := price(data)
	want := 236

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test7(t *testing.T) {
	data := []string{
		"AAAAAA",
		"AAABBA",
		"AAABBA",
		"ABBAAA",
		"ABBAAA",
		"AAAAAA",
	}

	_, got := price(data)
	want := 368

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test8(t *testing.T) {
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

	_, got := price(data)
	want := 1206

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
