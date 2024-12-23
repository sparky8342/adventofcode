package day22

import (
	"fmt"
	"loader"
)

const ITERATIONS = 2000

func next(n int) int {
	n = n ^ (n*64)%16777216
	n = n ^ (n/32)%16777216
	n = n ^ (n*2048)%16777216
	return n
}

func sequence(n int) int {
	for i := 0; i < ITERATIONS; i++ {
		n = next(n)
	}
	return n
}

func sequences(nums []int) int {
	total := 0
	for _, n := range nums {
		total += sequence(n)
	}
	return total
}

func change_sequence(n int) map[[4]int]int {
	sequences := map[[4]int]int{}
	last_digit := 0

	sequence := [4]int{}

	for i := 0; i < ITERATIONS; i++ {
		n = next(n)
		digit := n % 10
		change := digit - last_digit
		for j := 0; j < 3; j++ {
			sequence[j] = sequence[j+1]
		}
		sequence[3] = change

		if i > 3 {
			if _, ok := sequences[sequence]; !ok {
				sequences[sequence] = digit
			}
		}

		last_digit = digit
	}

	return sequences
}

func best_sequence(nums []int) int {
	all_sequences := map[[4]int]struct{}{}
	change_sequences := make([]map[[4]int]int, len(nums))
	for i, n := range nums {
		change_sequences[i] = change_sequence(n)
		for k := range change_sequences[i] {
			all_sequences[k] = struct{}{}
		}
	}

	max := 0
	for seq := range all_sequences {
		bananas := 0
		for _, change_seq := range change_sequences {
			bananas += change_seq[seq]
		}
		if bananas > max {
			max = bananas
		}
	}

	return max
}

func Run() {
	loader.Day = 22
	nums := loader.GetInts()

	part1 := sequences(nums)
	part2 := best_sequence(nums)

	fmt.Printf("%d %d\n", part1, part2)
}
