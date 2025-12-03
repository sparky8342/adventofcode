package day3

import (
	"fmt"
	"loader"
)

func max_joltage(bank string) int {
	max := 0
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			joltage := int(bank[i]-'0')*10 + int(bank[j]-'0')
			if joltage > max {
				max = joltage
			}
		}
	}
	return max
}

func total_joltage(banks []string) int {
	total := 0
	for _, bank := range banks {
		total += max_joltage(bank)
	}
	return total
}

func Run() {
	loader.Day = 3
	banks := loader.GetStrings()
	part1 := total_joltage(banks)

	fmt.Printf("%d %d\n", part1, 0)
}
