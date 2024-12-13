package day13

import "testing"

var machines = parse_data([][]string{
	[]string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
	},
	[]string{
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
	},
	[]string{
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
	},
	[]string{
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
	},
})

func Test1(t *testing.T) {
	got := tokens(machines)
	want := 480

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
