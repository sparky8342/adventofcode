package main

import (
	"strings"
	"testing"
)

const data = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func TestRoot(t *testing.T) {
	root, _ := parse_data(strings.Split(data, "\n"))
	got_root := eval(root)
	want_root := 152

	if got_root != want_root {
		t.Errorf("got %d, wanted %d", got_root, want_root)
	}
}

func TestFindHumn(t *testing.T) {
	root, humn := parse_data(strings.Split(data, "\n"))
	got_humn := find_humn_simple(root, humn)
	want_humn := 301

	if got_humn != want_humn {
		t.Errorf("got %d, wanted %d", got_humn, want_humn)
	}
}
