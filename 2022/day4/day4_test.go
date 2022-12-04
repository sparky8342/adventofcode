package main

import "testing"

var data = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

var pairs = get_pairs(data)

func TestRangeContains(t *testing.T) {
	got_amount := range_contains(pairs)
	want_amount := 2

	if got_amount != want_amount {
		t.Errorf("got %d, wanted %d", got_amount, want_amount)
	}
}

func TestRangeOverlaps(t *testing.T) {
	got_amount := range_overlaps(pairs)
	want_amount := 4

	if got_amount != want_amount {
		t.Errorf("got %d, wanted %d", got_amount, want_amount)
	}
}
