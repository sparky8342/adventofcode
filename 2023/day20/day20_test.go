package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"broadcaster -> a, b, c",
		"%a -> b",
		"%b -> c",
		"%c -> inv",
		"&inv -> a",
	}

	broadcaster := parse_data(data)
	got_answer := press_button(broadcaster, 1000, false)
	want_answer := 32000000

	if got_answer != want_answer {
		t.Errorf("got %d, wanted %d", got_answer, want_answer)
	}

	data = []string{
		"broadcaster -> a",
		"%a -> inv, con",
		"&inv -> b",
		"%b -> con",
		"&con -> output",
	}

	broadcaster = parse_data(data)
	got_answer = press_button(broadcaster, 1000, false)
	want_answer = 11687500

	if got_answer != want_answer {
		t.Errorf("got %d, wanted %d", got_answer, want_answer)
	}
}
