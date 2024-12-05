package day5

import (
	"fmt"
	"loader"
	"os"
	"strconv"
	"strings"
)

func valid_order(rules map[[2]int]struct{}, nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if _, ok := rules[[2]int{nums[i], nums[j]}]; !ok {
				return false
			}
		}
	}
	return true
}

func sort_update(rules map[[2]int]struct{}, nums []int) {
	done := false
	for !done {
		done = true
		for i := 0; i < len(nums)-1; i++ {
			if _, ok := rules[[2]int{nums[i], nums[i+1]}]; !ok {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				done = false
			}
		}
	}
}

func valid_updates(data [][]string) (int, int) {
	rules := map[[2]int]struct{}{}
	for _, line := range data[0] {
		parts := strings.Split(line, "|")
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		rules[[2]int{n1, n2}] = struct{}{}
	}

	valid_sum := 0
	corrected_sum := 0

	for _, line := range data[1] {
		row := []int{}
		for _, str := range strings.Split(line, ",") {
			n, err := strconv.Atoi(str)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			row = append(row, n)
		}

		if valid_order(rules, row) {
			valid_sum += row[len(row)/2]
		} else {
			sort_update(rules, row)
			corrected_sum += row[len(row)/2]
		}
	}

	return valid_sum, corrected_sum
}

func Run() {
	loader.Day = 5
	data := loader.GetStringGroups()
	part1, part2 := valid_updates(data)

	fmt.Printf("%d %d\n", part1, part2)
}
