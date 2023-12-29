package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"jqt: rhn xhk nvd",
		"rsh: frs pzl lsr",
		"xhk: hfx",
		"cmg: qnr nvd lhk bvb",
		"rhn: xhk bvb hfx",
		"bvb: xhk hfx",
		"pzl: lsr hfx nvd",
		"qnr: nvd",
		"ntq: jqt hfx bvb xhk",
		"nvd: lhk",
		"lsr: lhk",
		"rzs: qnr cmg lsr rsh",
		"frs: qnr lhk lsr",
	}

	nodes := parse_data(data)
	got_product := find_split(nodes)
	want_product := 54

	if got_product != want_product {
		t.Errorf("got %d, wanted %d", got_product, want_product)
	}
}
