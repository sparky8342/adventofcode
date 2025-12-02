package day2

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type pair [2]int

func parse_data(data string) []pair {
	ranges := []pair{}
	lines := strings.Split(data, ",")
	for _, line := range lines {
		parts := strings.Split(line, "-")
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, pair{n1, n2})
	}
	return ranges
}

func invalid(n int) bool {
	str := strconv.Itoa(n)
	l := len(str)
	half := l / 2

	if l%2 == 1 {
		return false
	}
	for i := 0; i < half; i++ {
		if str[i] != str[i+half] {
			return false
		}
	}

	return true
}

func total_invalid(ranges []pair) int {
	total := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if invalid(i) {
				total += i
			}
		}
	}
	return total
}

func Run() {
	loader.Day = 2
	data := loader.GetOneString()
	ranges := parse_data(data)
	part1 := total_invalid(ranges)

	fmt.Printf("%d %d\n", part1, 0)
}
