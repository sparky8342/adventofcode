package day1

import "testing"

func Test1(t *testing.T) {
	data := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	nums1, nums2 := parse_data(data)

	got := distance(nums1, nums2)
	want := 11

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = similarity(nums1, nums2)
	want = 31

	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
