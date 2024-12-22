package day22

import (
	"fmt"
	"loader"
)

func next(n int) int {
	n = n ^ (n*64)%16777216
	n = n ^ (n/32)%16777216
	n = n ^ (n*2048)%16777216
	return n
}

func sequence(n int, amount int) int {
	for i := 0; i < amount; i++ {
		n = next(n)
	}
	return n
}

func sequences(nums []int, amount int) int {
	total := 0
	for _, n := range nums {
		total += sequence(n, amount)
	}
	return total
}

func Run() {
	loader.Day = 22
	nums := loader.GetInts()

	part1 := sequences(nums, 2000)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
