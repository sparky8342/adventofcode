package main

import (
	"strings"
	"testing"
)

var data = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func TestPressure(t *testing.T) {
	start := parse_data(strings.Split(data, "\n"))

	got_pressure := find_best_pressure(start)
	want_pressure := 1651

	if got_pressure != want_pressure {
		t.Errorf("got %d, wanted %d", got_pressure, want_pressure)
	}
}

func TestPressureWithElephant(t *testing.T) {
	start := parse_data(strings.Split(data, "\n"))

	got_pressure := find_best_pressure_with_elephant(start)
	want_pressure := 1707

	if got_pressure != want_pressure {
		t.Errorf("got %d, wanted %d", got_pressure, want_pressure)
	}
}
