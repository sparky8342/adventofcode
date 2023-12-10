package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func load_data(filename string) [][]int {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	lines := strings.Split(string(data), "\n")
	nums := [][]int{}

	for _, line := range lines {
		row := []int{}
		for _, n_str := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(n_str)
			row = append(row, n)
		}
		nums = append(nums, row)
	}

	return nums
}

func find_next_value(nums []int) int {
	next := make([]int, len(nums)-1)
	all_zero := true
	for i := 0; i < len(nums)-1; i++ {
		next[i] = nums[i+1] - nums[i]
		if next[i] != 0 {
			all_zero = false
		}
	}
	next_value := 0
	if !all_zero {
		next_value = find_next_value(next)
	}
	return nums[len(nums)-1] + next_value
}

func reverse(nums []int) []int {
	l := len(nums)
	rev := make([]int, l)
	for i := 0; i < l; i++ {
		rev[i] = nums[l-i-1]
	}
	return rev
}

func calculate_sum(nums [][]int) (int, int) {
	sum := 0
	sum2 := 0
	for _, row := range nums {
		sum += find_next_value(row)
		sum2 += find_next_value(reverse(row))
	}
	return sum, sum2
}

func main() {
	nums := load_data("input.txt")
	next_sum, previous_sum := calculate_sum(nums)
	fmt.Println(next_sum)
	fmt.Println(previous_sum)
}
