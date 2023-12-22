package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"1,0,1~1,2,1",
		"0,0,2~2,0,2",
		"0,2,3~2,2,3",
		"0,0,4~0,2,4",
		"2,0,5~2,2,5",
		"0,1,6~2,1,6",
		"1,1,8~1,1,9",
	}

	bricks, positions := parse_data(data)
	move_down(bricks, positions)

	got_bricks := free_bricks(bricks, positions)
	want_bricks := 5

	if got_bricks != want_bricks {
		t.Errorf("got %d, wanted %d", got_bricks, want_bricks)
	}
}
