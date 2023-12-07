package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	hands := parse_data(data)

	got_winnings := winnings(hands)
	want_winnings := 6440

	if got_winnings != want_winnings {
		t.Errorf("got %d, wanted %d", got_winnings, want_winnings)
	}
}
