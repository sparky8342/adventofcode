package main

import (
	"testing"
)

var data = []string{
	"Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
	"Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
}

func TestQuality(t *testing.T) {
	blueprints := parse_data(data)
	got_quality := find_quality(blueprints)
	want_quality := 33

	if got_quality != want_quality {
		t.Errorf("got %d, wanted %d", got_quality, want_quality)
	}
}
