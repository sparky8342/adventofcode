package day23

import "testing"

var data = []string{
	"kh-tc",
	"qp-kh",
	"de-cg",
	"ka-co",
	"yn-aq",
	"qp-ub",
	"cg-tb",
	"vc-aq",
	"tb-ka",
	"wh-tc",
	"yn-cg",
	"kh-ub",
	"ta-co",
	"de-co",
	"tc-td",
	"tb-wq",
	"wh-td",
	"ta-ka",
	"td-qp",
	"aq-cg",
	"wq-ub",
	"ub-vc",
	"de-ta",
	"wq-aq",
	"wq-vc",
	"wh-yn",
	"ka-de",
	"kh-ta",
	"co-tc",
	"wh-qp",
	"tb-vc",
	"td-yn",
}

var nodes = parse_data(data)

func Test1(t *testing.T) {
	got := find_three(nodes)
	want := 7

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := find_largest_group(nodes)
	want := "co,de,ka,ta"

	if want != got {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
