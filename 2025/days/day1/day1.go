package day1

import (
	"fmt"
	"loader"
	"strconv"
)

func turn_dial(data []string) (int, int) {
	dial := 50
	zeroes := 0
	zeroes_at_end := 0

	for _, row := range data {
		n, err := strconv.Atoi(row[1:])
		if err != nil {
			panic(err)
		}

		zeroes += n / 100
		n %= 100

		if row[0] == 'L' {
			if dial == 0 {
				dial = (dial - n) + 100
			} else {
				dial -= n
				if dial == 0 {
					zeroes++
				} else if dial < 0 {
					dial += 100
					zeroes++
				}
			}
		} else if row[0] == 'R' {
			dial += n
			if dial > 99 {
				dial -= 100
				zeroes++
			}
		}
		if dial == 0 {
			zeroes_at_end++
		}
	}

	return zeroes_at_end, zeroes
}

func Run() {
	loader.Day = 1
	data := loader.GetStrings()
	part1, part2 := turn_dial(data)

	fmt.Printf("%d %d\n", part1, part2)
}
