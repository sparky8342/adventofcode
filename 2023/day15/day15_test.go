package main

import "testing"

func Test(t *testing.T) {
	data := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

	got_hash_sum := hash_sum(data)
	want_hash_sum := 1320

	if got_hash_sum != want_hash_sum {
		t.Errorf("got %d, wanted %d", got_hash_sum, want_hash_sum)
	}

	got_power := operations(data)
	want_power := 145

	if got_power != want_power {
		t.Errorf("got %d, wanted %d", got_power, want_power)
	}
}
