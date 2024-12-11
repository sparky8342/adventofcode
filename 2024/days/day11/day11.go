package day11

import (
	"fmt"
	"loader"
)

func blink(nums []int, amount int) int {
	stones := map[int]int{}
	for _, n := range nums {
		stones[n]++
	}

	for i := 0; i < amount; i++ {
		next_stones := map[int]int{}
		for val, count := range stones {
			if val == 0 {
				next_stones[1] += count
				continue
			}
			digits := 0
			n := val
			for n > 0 {
				n /= 10
				digits++
			}
			if digits%2 == 0 {
				n1 := val
				for j := 0; j < digits/2; j++ {
					n1 /= 10
				}
				n3 := n1
				for j := 0; j < digits/2; j++ {
					n3 *= 10
				}
				n2 := val - n3
				next_stones[n1] += count
				next_stones[n2] += count
			} else {
				next_stones[val*2024] += count
			}
		}
		stones = next_stones
	}

	total := 0
	for _, count := range stones {
		total += count
	}

	return total
}

func Run() {
	loader.Day = 11
	nums := loader.GetIntLine()

	part1 := blink(nums, 25)
	part2 := blink(nums, 75)

	fmt.Printf("%d %d\n", part1, part2)
}
