package day1

import (
	"fmt"
	"loader"
	"strconv"
)

func turn_dial(data []string) int {
	dial := 50
	zeroes := 0

	for _, row := range data {
		n, err := strconv.Atoi(row[1:])
		if err != nil {
			panic(err)
		}
		if row[0] == 'L' {
			n *= -1
		}
		dial = (dial + n) % 100
		if dial == 0 {
			zeroes++
		}
	}

	return zeroes
}

func Run() {
	loader.Day = 1
	data := loader.GetStrings()
	part1 := turn_dial(data)

	fmt.Printf("%d %d\n", part1, 0)
}
