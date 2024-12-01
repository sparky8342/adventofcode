package day1

import (
	"fmt"
	"loader"
	"sort"
	"utils"
)

func distance(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	total := 0
	for i := 0; i < len(nums1); i++ {
		total += utils.Abs(nums1[i] - nums2[i])
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
	nums := loader.GetIntColumns()

	part1 := distance(nums[0], nums[1])
	part2 := similarity(nums[0], nums[1])

	fmt.Printf("%d %d\n", part1, part2)
}
