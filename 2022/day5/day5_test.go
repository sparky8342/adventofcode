package main

import "testing"

var data = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestTopOfStack(t *testing.T) {
	stacks, instructions := parse_data(data)
	process_instructions(stacks, instructions)
	got_top := top_of_stacks(stacks)
	want_top := "CMZ"

	if got_top != want_top {
		t.Errorf("got %s, wanted %s", got_top, want_top)
	}
}

func TestTopOfStackWith9001(t *testing.T) {
	stacks, instructions := parse_data(data)
	process_instructions_9001(stacks, instructions)
	got_top := top_of_stacks(stacks)
	want_top := "MCD"

	if got_top != want_top {
		t.Errorf("got %s, wanted %s", got_top, want_top)
	}
}
