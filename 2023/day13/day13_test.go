package main

import "testing"

func Test(t *testing.T) {
	data := "#.##..##.\n" +
		"..#.##.#.\n" +
		"##......#\n" +
		"##......#\n" +
		"..#.##.#.\n" +
		"..##..##.\n" +
		"#.#.##.#.\n" +
		"\n" +
		"#...##..#\n" +
		"#....#..#\n" +
		"..##..###\n" +
		"#####.##.\n" +
		"#####.##.\n" +
		"..##..###\n" +
		"#....#..#"

	got_notes := get_notes(data)
	want_notes := 405

	if got_notes != want_notes {
		t.Errorf("got %d, wanted %d", got_notes, want_notes)
	}
}
