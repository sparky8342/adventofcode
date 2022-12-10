package main

import (
	"fmt"
	"testing"
)

func TestProgram1(t *testing.T) {
	data := []string{
		"addx 15",
		"addx -11",
		"addx 6",
		"addx -3",
		"addx 5",
		"addx -1",
		"addx -8",
		"addx 13",
		"addx 4",
		"noop",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx -35",
		"addx 1",
		"addx 24",
		"addx -19",
		"addx 1",
		"addx 16",
		"addx -11",
		"noop",
		"noop",
		"addx 21",
		"addx -15",
		"noop",
		"noop",
		"addx -3",
		"addx 9",
		"addx 1",
		"addx -3",
		"addx 8",
		"addx 1",
		"addx 5",
		"noop",
		"noop",
		"noop",
		"noop",
		"noop",
		"addx -36",
		"noop",
		"addx 1",
		"addx 7",
		"noop",
		"noop",
		"noop",
		"addx 2",
		"addx 6",
		"noop",
		"noop",
		"noop",
		"noop",
		"noop",
		"addx 1",
		"noop",
		"noop",
		"addx 7",
		"addx 1",
		"noop",
		"addx -13",
		"addx 13",
		"addx 7",
		"noop",
		"addx 1",
		"addx -33",
		"noop",
		"noop",
		"noop",
		"addx 2",
		"noop",
		"noop",
		"noop",
		"addx 8",
		"noop",
		"addx -1",
		"addx 2",
		"addx 1",
		"noop",
		"addx 17",
		"addx -9",
		"addx 1",
		"addx 1",
		"addx -3",
		"addx 11",
		"noop",
		"noop",
		"addx 1",
		"noop",
		"addx 1",
		"noop",
		"noop",
		"addx -13",
		"addx -19",
		"addx 1",
		"addx 3",
		"addx 26",
		"addx -30",
		"addx 12",
		"addx -1",
		"addx 3",
		"addx 1",
		"noop",
		"noop",
		"noop",
		"addx -9",
		"addx 18",
		"addx 1",
		"addx 2",
		"noop",
		"noop",
		"addx 9",
		"noop",
		"noop",
		"noop",
		"addx -1",
		"addx 2",
		"addx -37",
		"addx 1",
		"addx 3",
		"noop",
		"addx 15",
		"addx -21",
		"addx 22",
		"addx -6",
		"addx 1",
		"noop",
		"addx 2",
		"addx 1",
		"noop",
		"addx -10",
		"noop",
		"noop",
		"addx 20",
		"addx 1",
		"addx 2",
		"addx 2",
		"addx -6",
		"addx -11",
		"noop",
		"noop",
		"noop",
	}

	got_sum := run_program(data)
	want_sum := 13140

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}

	got_crt := run_program_part2(data)
	crt_strings := []string{
		"##  ##  ##  ##  ##  ##  ##  ##  ##  ##  ",
		"###   ###   ###   ###   ###   ###   ### ",
		"####    ####    ####    ####    ####    ",
		"#####     #####     #####     #####     ",
		"######      ######      ######      ####",
		"#######       #######       #######     ",
	}
	want_crt := Crt{}
	for row, line := range crt_strings {
		for col, ch := range line {
			if ch == '#' {
				want_crt[row][col] = true
			} else {
				want_crt[row][col] = false
			}
		}
	}

	if got_crt != want_crt {
		t.Errorf("crt doesn't match")
		fmt.Println("got:")
		print_crt(&got_crt)
		fmt.Println("wanted:")
		print_crt(&want_crt)
	}
}
