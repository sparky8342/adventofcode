package day3

import (
	"fmt"
	"loader"
)

type Digit struct {
	value int
	on    bool
}

type Bank struct {
	digits []Digit
}

func (b Bank) number() int {
	n := 0
	for _, digit := range b.digits {
		if digit.on {
			n = n*10 + digit.value
		}
	}
	return n
}

func max_joltage(bank_str string, batteries int) int {
	bank := Bank{}
	for _, b := range bank_str {
		bank.digits = append(bank.digits, Digit{
			value: int(b - '0'),
			on:    false,
		})
	}

	for i := 0; i < batteries; i++ {
		bank.digits[i].on = true
	}

	current := bank.number()

	done := false
	for !done {
		done = true
		for i := 0; i < len(bank.digits); i++ {
			for j := i + 1; j < len(bank.digits); j++ {
				if bank.digits[i].on && !bank.digits[j].on {
					bank.digits[i].on = false
					bank.digits[j].on = true
					n := bank.number()
					if n < current {
						bank.digits[i].on = true
						bank.digits[j].on = false
					} else {
						current = n
						done = false
					}
				}
			}
		}
	}

	return current
}

func total_joltage(banks []string, batteries int) int {
	total := 0
	for _, bank := range banks {
		total += max_joltage(bank, batteries)
	}
	return total
}

func Run() {
	loader.Day = 3
	banks := loader.GetStrings()
	part1 := total_joltage(banks, 2)
	part2 := total_joltage(banks, 12)

	fmt.Printf("%d %d\n", part1, part2)
}
