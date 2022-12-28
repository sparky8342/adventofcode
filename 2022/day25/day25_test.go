package main

import (
	"strings"
	"testing"
)

const data = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

func Test(t *testing.T) {
	got_sum := get_sum(strings.Split(data, "\n"))
	want_sum := 4890

	if got_sum != want_sum {
		t.Errorf("got %d, wanted %d", got_sum, want_sum)
	}

	got_snafu := dec_to_snafu(want_sum)
	want_snafu := "2=-1=0"

	if got_snafu != want_snafu {
		t.Errorf("got %s, wanted %s", got_snafu, want_snafu)
	}
}
