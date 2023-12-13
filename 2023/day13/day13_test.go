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

	got_notes, got_notes2 := get_notes(data)
	want_notes := 405
	want_notes2 := 400

	if got_notes != want_notes {
		t.Errorf("got %d, wanted %d", got_notes, want_notes)
	}

	if got_notes2 != want_notes2 {
		t.Errorf("got %d, wanted %d", got_notes2, want_notes2)
	}
}
