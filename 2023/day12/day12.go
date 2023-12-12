package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type CacheKey struct {
	seq_i   int
	size_i  int
	current int
}

var cache map[CacheKey]int

func init() {
	cache = map[CacheKey]int{}
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

// thanks to jonathan paulson for this algorithm
func combinations(sequence []byte, sizes []int, seq_i int, size_i int, current int) int {
	key := CacheKey{seq_i: seq_i, size_i: size_i, current: current}
	if value, exists := cache[key]; exists {
		return value
	}

	if seq_i == len(sequence) {
		if size_i == len(sizes) && current == 0 {
			return 1
		} else if size_i == len(sizes)-1 && sizes[size_i] == current {
			return 1
		} else {
			return 0
		}
	}

	ans := 0
	for _, c := range []byte{'.', '#'} {
		if sequence[seq_i] == c || sequence[seq_i] == '?' {
			if c == '.' && current == 0 {
				ans += combinations(sequence, sizes, seq_i+1, size_i, 0)
			} else if c == '.' && current > 0 && size_i < len(sizes) && sizes[size_i] == current {
				ans += combinations(sequence, sizes, seq_i+1, size_i+1, 0)
			} else if c == '#' {
				ans += combinations(sequence, sizes, seq_i+1, size_i, current+1)
			}
		}
	}

	cache[key] = ans
	return ans
}

func unfold(sequence []byte, sizes []int) ([]byte, []int) {
	unfolded_sequence := sequence
	for i := 0; i < 4; i++ {
		unfolded_sequence = append(unfolded_sequence, '?')
		unfolded_sequence = append(unfolded_sequence, sequence...)
	}
	unfolded_sizes := []int{}
	for i := 0; i < 5; i++ {
		unfolded_sizes = append(unfolded_sizes, sizes...)
	}
	return unfolded_sequence, unfolded_sizes
}

func get_sum(sequence []byte, sizes []int) int {
	cache = map[CacheKey]int{}
	return combinations(sequence, sizes, 0, 0, 0)
}

func total_sum(data []string) (int, int) {
	sum := 0
	unfolded_sum := 0

	for _, line := range data {
		parts := strings.Split(line, " ")
		sequence := []byte(parts[0])
		size_strs := strings.Split(parts[1], ",")
		sizes := []int{}
		for _, size_str := range size_strs {
			n, _ := strconv.Atoi(size_str)
			sizes = append(sizes, n)
		}

		sum += get_sum(sequence, sizes)

		unfolded_sequence, unfolded_sizes := unfold(sequence, sizes)
		unfolded_sum += get_sum(unfolded_sequence, unfolded_sizes)
	}

	return sum, unfolded_sum
}

func main() {
	data := load_data("input.txt")
	total_sum, total_unfolded_sum := total_sum(data)
	fmt.Println(total_sum)
	fmt.Println(total_unfolded_sum)
}
