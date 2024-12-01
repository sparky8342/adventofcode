package day1

import (
	"fmt"
	"loader"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func parse_data(data []string) ([]int, []int) {
	nums1 := make([]int, len(data))
	nums2 := make([]int, len(data))

	for i, line := range data {
		parts := strings.Fields(line)
		n, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		nums1[i] = n
		n, err = strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		nums2[i] = n
	}
	return nums1, nums2
}

func distance(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	total := 0
	for i := 0; i < len(nums1); i++ {
		total += abs(nums1[i] - nums2[i])
	}
	return total
}

func similarity(nums []int, nums2 []int) int {
	counts := map[int]int{}
	for _, n := range nums2 {
		counts[n]++
	}
	total := 0
	for _, n := range nums {
		total += n * counts[n]
	}
	return total
}

func Run() {
	loader.Day = 1
	data := loader.GetStrings()
	nums1, nums2 := parse_data(data)

	part1 := distance(nums1, nums2)
	part2 := similarity(nums1, nums2)

	fmt.Printf("%d %d\n", part1, part2)
}
