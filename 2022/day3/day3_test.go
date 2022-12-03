package main

import "testing"

var data = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestPrioritySum(t *testing.T) {
	got_sum := priority_sum(data)
	want_sum := 157

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}
}

func TestBadgeSum(t *testing.T) {
	got_sum := badge_sum(data)
	want_sum := 70

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}
}
