package day9

import "testing"

var data = []byte("2333133121414131402")

func Test1(t *testing.T) {
	got := defrag(data)
	want := 1928

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := defrag_files(data)
	want := 2858

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
