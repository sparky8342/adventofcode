package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}

	got_sum, got_unfolded_sum := total_sum(data)
	want_sum := 21
	want_unfolded_sum := 525152

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}

	if got_unfolded_sum != want_unfolded_sum {
		t.Errorf("got %d, wanted %d", got_unfolded_sum, want_unfolded_sum)
	}
}
