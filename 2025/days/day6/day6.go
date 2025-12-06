package day6

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

func parse_data_part1(data []string) ([][]int, []byte) {
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

func parse_data_part2(data []string) ([][]int, []byte) {
	height := len(data)
	width := len(data[0])

	nums := [][]int{}
	operators := []byte{}

	col := []int{}
	for x := width - 1; x >= 0; x-- {
		num := 0
		for y := 0; y < height-1; y++ {
			if data[y][x] != ' ' {
				num = num*10 + int(data[y][x]-'0')
			}
		}
		col = append(col, num)

		if data[height-1][x] != ' ' {
			operators = append(operators, data[height-1][x])
			nums = append(nums, col)
			col = []int{}
			x--
		}
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
	nums, operators := parse_data_part1(data)
	part1 := calculate(nums, operators)

	nums, operators = parse_data_part2(data)
	part2 := calculate(nums, operators)

	fmt.Printf("%d %d\n", part1, part2)
}
