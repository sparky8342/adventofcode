package day6

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

func parse_data(data []string) ([][]int, []byte) {
	no_cols := len(strings.Fields(data[0]))
	nums := make([][]int, no_cols)
	for i := range nums {
		nums[i] = make([]int, len(data)-1)
	}

	for i := 0; i < len(data)-1; i++ {
		for j, n_str := range strings.Fields(data[i]) {
			n, err := strconv.Atoi(n_str)
			if err != nil {
				panic(err)
			}
			nums[j][i] = n
		}
	}

	strs := strings.Fields(data[len(data)-1])
	operators := make([]byte, len(strs))
	for i, str := range strs {
		operators[i] = str[0]
	}

	return nums, operators
}

func calculate(nums [][]int, operators []byte) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		if operators[i] == '*' {
			product := 1
			for _, n := range nums[i] {
				product *= n
			}
			total += product
		} else if operators[i] == '+' {
			sum := 0
			for _, n := range nums[i] {
				sum += n
			}
			total += sum
		}
	}
	return total
}

func Run() {
	loader.Day = 6
	data := loader.GetStrings()
	nums, operators := parse_data(data)
	part1 := calculate(nums, operators)

	fmt.Printf("%d %d\n", part1, 0)
}
