package day25

import (
	"fmt"
	"loader"
)

func parse_data(data [][]string) ([][]int, [][]int) {
	keys := [][]int{}
	locks := [][]int{}
	for _, group := range data {
		if group[0] == "#####" {
			lock := make([]int, 5)
			for x := 0; x < 5; x++ {
				for y := 1; y < 7; y++ {
					if group[y][x] == '.' {
						lock[x] = y - 1
						break
					}
				}
			}
			locks = append(locks, lock)
		} else if group[6] == "#####" {
			key := make([]int, 5)
			for x := 0; x < 5; x++ {
				for y := 5; y >= 0; y-- {
					if group[y][x] == '.' {
						key[x] = 5 - y
						break
					}
				}
			}
			keys = append(keys, key)
		}
	}

	return keys, locks
}

func fit(key []int, lock []int) bool {
	for i := 0; i < 5; i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func fit_combinations(keys [][]int, locks [][]int) int {
	total := 0
	for _, key := range keys {
		for _, lock := range locks {
			if fit(lock, key) {
				total++
			}
		}
	}
	return total
}

func Run() {
	loader.Day = 25
	data := loader.GetStringGroups()
	keys, locks := parse_data(data)

	part1 := fit_combinations(keys, locks)

	fmt.Printf("%d\n", part1)
}
