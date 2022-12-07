package main

import (
	"strings"
	"testing"
)

const data = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestGetSize(t *testing.T) {
	tree := create_tree(strings.Split(data, "\n"))
	got_size, got_smallest := get_sizes(tree)
	want_size, want_smallest := 95437, 24933642

	if got_size != want_size {
		t.Errorf("got %d, wanted %d", got_size, want_size)
	}

	if got_smallest != want_smallest {
		t.Errorf("got %d, wanted %d", got_smallest, want_smallest)
	}
}
