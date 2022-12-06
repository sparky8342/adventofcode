package main

import "testing"

func TestExample1(t *testing.T) {
	buffer := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	got_marker := find_marker(buffer)
	want_marker := 7

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample2(t *testing.T) {
	buffer := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	got_marker := find_marker(buffer)
	want_marker := 5

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample3(t *testing.T) {
	buffer := "nppdvjthqldpwncqszvftbrmjlhg"
	got_marker := find_marker(buffer)
	want_marker := 6

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample4(t *testing.T) {
	buffer := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	got_marker := find_marker(buffer)
	want_marker := 10

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample5(t *testing.T) {
	buffer := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	got_marker := find_marker(buffer)
	want_marker := 11

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}
