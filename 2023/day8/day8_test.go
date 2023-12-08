package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	instructions, nodes := parse_data(data)

	got_steps := follow_instructions(instructions, nodes)
	want_steps := 2

	if got_steps != want_steps {
		t.Errorf("got %d, wanted %d", got_steps, want_steps)
	}

	data = []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	instructions, nodes = parse_data(data)

	got_steps = follow_instructions(instructions, nodes)
	want_steps = 6

	if got_steps != want_steps {
		t.Errorf("got %d, wanted %d", got_steps, want_steps)
	}
}
