package day2

import (
	"fmt"
	"loader"
	"utils"
)

func is_safe(report []int) bool {
	asc := false
	desc := false
	for i := 0; i < len(report)-1; i++ {
		if utils.Abs(report[i]-report[i+1]) > 3 {
			return false
		} else if report[i] == report[i+1] {
			return false
		}
		if report[i] > report[i+1] {
			desc = true
		} else if report[i] < report[i+1] {
			asc = true
		}
	}
	return !(asc && desc)
}

func count_safe(nums [][]int) int {
	safe := 0
	for _, report := range nums {
		if is_safe(report) {
			safe++
		}
	}
	return safe
}

func count_safe_with_tolerance(nums [][]int) int {
	safe := 0
	for _, report := range nums {
		this_safe := false
		for i := 0; i < len(report); i++ {
			check_report := make([]int, len(report))
			copy(check_report, report)
			if is_safe(append(check_report[0:i], check_report[i+1:]...)) {
				this_safe = true
				break
			}
		}
		if this_safe {
			safe++
		}
	}
	return safe
}

func Run() {
	loader.Day = 2
	nums := loader.GetIntRows()

	part1 := count_safe(nums)
	part2 := count_safe_with_tolerance(nums)

	fmt.Printf("%d %d\n", part1, part2)
}
