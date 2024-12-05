package day5

import "testing"

var data = [][]string{
	[]string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	},
	[]string{
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	},
}

func Test1(t *testing.T) {
	got_valid, got_corrected := valid_updates(data)
	want_valid := 143

	if want_valid != got_valid {
		t.Errorf("got %d, wanted %d", got_valid, want_valid)
	}

	want_corrected := 123
	if want_corrected != got_corrected {
		t.Errorf("got %d, wanted %d", got_corrected, want_corrected)
	}
}
