package main

import "testing"

func Test(t *testing.T) {
	data := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	got_calibration, _ := calibration(data)
	want_calibration := 142

	if got_calibration != want_calibration {
		t.Errorf("got %d, wanted %d", got_calibration, want_calibration)
	}

	data = []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	_, got_calibration = calibration(data)
	want_calibration = 281

	if got_calibration != want_calibration {
		t.Errorf("got %d, wanted %d", got_calibration, want_calibration)
	}
}
