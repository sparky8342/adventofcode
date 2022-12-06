package main

import "testing"

func TestExample1(t *testing.T) {
	buffer := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	got_marker := find_marker(buffer, 4)
	want_marker := 7

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample2(t *testing.T) {
	buffer := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	got_marker := find_marker(buffer, 4)
	want_marker := 5

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample3(t *testing.T) {
	buffer := "nppdvjthqldpwncqszvftbrmjlhg"
	got_marker := find_marker(buffer, 4)
	want_marker := 6

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample4(t *testing.T) {
	buffer := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	got_marker := find_marker(buffer, 4)
	want_marker := 10

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample5(t *testing.T) {
	buffer := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	got_marker := find_marker(buffer, 4)
	want_marker := 11

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample6(t *testing.T) {
	buffer := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	got_marker := find_marker(buffer, 14)
	want_marker := 19

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample7(t *testing.T) {
	buffer := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	got_marker := find_marker(buffer, 14)
	want_marker := 23

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample8(t *testing.T) {
	buffer := "nppdvjthqldpwncqszvftbrmjlhg"
	got_marker := find_marker(buffer, 14)
	want_marker := 23

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample9(t *testing.T) {
	buffer := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	got_marker := find_marker(buffer, 14)
	want_marker := 29

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}

func TestExample10(t *testing.T) {
	buffer := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	got_marker := find_marker(buffer, 14)
	want_marker := 26

	if got_marker != want_marker {
		t.Errorf("got %d, wanted %d", got_marker, want_marker)
	}
}
