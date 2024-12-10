package day10

import "testing"

func Test1(t *testing.T) {
	grid := []string{
		"0123",
		"1234",
		"8765",
		"9876",
	}

	got, _ := score(grid)
	want := 1

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	grid := []string{
		"...0...",
		"...1...",
		"...2...",
		"6543456",
		"7.....7",
		"8.....8",
		"9.....9",
	}

	got, _ := score(grid)
	want := 2

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	grid := []string{
		"..90..9",
		"...1.98",
		"...2..7",
		"6543456",
		"765.987",
		"876....",
		"987....",
	}

	got, _ := score(grid)
	want := 4

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	grid := []string{
		"10..9..",
		"2...8..",
		"3...7..",
		"4567654",
		"...8..3",
		"...9..2",
		".....01",
	}

	got, _ := score(grid)
	want := 3

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	grid := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}

	got, _ := score(grid)
	want := 36

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test6(t *testing.T) {
	grid := []string{
		".....0.",
		"..4321.",
		"..5..2.",
		"..6543.",
		"..7..4.",
		"..8765.",
		"..9....",
	}

	_, got := score(grid)
	want := 3

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test7(t *testing.T) {
	grid := []string{
		"..90..9",
		"...1.98",
		"...2..7",
		"6543456",
		"765.987",
		"876....",
		"987....",
	}

	_, got := score(grid)
	want := 13

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test8(t *testing.T) {
	grid := []string{
		"012345",
		"123456",
		"234567",
		"345678",
		"4.6789",
		"56789.",
	}

	_, got := score(grid)
	want := 227

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test9(t *testing.T) {
	grid := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}

	_, got := score(grid)
	want := 81

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
