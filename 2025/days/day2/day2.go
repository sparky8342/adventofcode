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

func repeats2(n int) bool {
	str := strconv.Itoa(n)
	l := len(str)

	if l%2 != 0 {
		return false
	}

	half := l / 2
	for i := 0; i < half; i++ {
		if str[i] != str[i+half] {
			return false
		}
	}

	return true
}

func repeats_any(n int) bool {
	str := strconv.Itoa(n)
	l := len(str)

outer:
	for repeat_length := 1; repeat_length <= l/2; repeat_length++ {
		if l%repeat_length != 0 {
			continue
		}

		pattern := str[0:repeat_length]

		for i := repeat_length; i <= l-repeat_length; i += repeat_length {
			if pattern != str[i:i+repeat_length] {
				continue outer
			}
		}

		return true
	}

	return false
}

func total_invalid(ranges []pair, mode int) int {
	total := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			switch mode {
			case 1:
				if repeats2(i) {
					total += i
				}
			case 2:
				if repeats_any(i) {
					total += i
				}
			}
		}
	}
	return total
}

func Run() {
	loader.Day = 2
	data := loader.GetOneString()
	ranges := parse_data(data)
	part1 := total_invalid(ranges, 1)

	part2 := total_invalid(ranges, 2)

	fmt.Printf("%d %d\n", part1, part2)
}
