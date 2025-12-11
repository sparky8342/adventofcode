package day11

import "testing"

func Test1(t *testing.T) {
	data := []string{
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

	you, _ := parse_data(data)

	got := count_paths(you)
	want := 5

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	data := []string{
		"svr: aaa bbb",
		"aaa: fft",
		"fft: ccc",
		"bbb: tty",
		"tty: ccc",
		"ccc: ddd eee",
		"ddd: hub",
		"hub: fff",
		"eee: dac",
		"dac: fff",
		"fff: ggg hhh",
		"ggg: out",
		"hhh: out",
	}

	_, svr := parse_data(data)

	got := count_paths_through(svr)
	want := 2

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
