package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"sync"
)

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func is_valid(sequence []byte, sizes []int) bool {
	found := []int{}

	current := 0
	all_scanned := true
	for _, b := range sequence {
		if b == '?' {
			all_scanned = false
			break
		} else if b == '#' {
			current++
		} else if b == '.' && current > 0 {
			found = append(found, current)
			current = 0
		}
	}
	if current > 0 {
		found = append(found, current)
	}

	if len(found) > len(sizes) {
		return false
	}

	if all_scanned {
		if len(found) != len(sizes) {
			return false
		}
		for i := 0; i < len(found); i++ {
			if found[i] != sizes[i] {
				return false
			}
		}
		return true

	} else if len(found) > 0 {
		last_found := len(found) - 1
		if found[last_found] > sizes[last_found] {
			return false
		}

		for i := 0; i < last_found-1; i++ {
			if found[i] != sizes[i] {
				return false
			}
		}
	}

	return true
}

func dfs(sequence []byte, wildcards []int, wildcard_index int, sizes []int) int {
	if !is_valid(sequence, sizes) {
		return 0
	}

	if wildcard_index == len(wildcards) {
		return 1
	}

	sequence[wildcards[wildcard_index]] = '#'
	valid := dfs(sequence, wildcards, wildcard_index+1, sizes)

	sequence[wildcards[wildcard_index]] = '.'
	valid += dfs(sequence, wildcards, wildcard_index+1, sizes)

	sequence[wildcards[wildcard_index]] = '?'

	return valid
}

func get_sum(sequence []byte, sizes []int) int {
	wildcards := []int{}
	for i, b := range sequence {
		if b == '?' {
			wildcards = append(wildcards, i)
		}
	}
	sum := dfs(sequence, wildcards, 0, sizes)
	fmt.Println(string(sequence), sum)
	return sum
}

func unfold(sequence []byte, sizes []int) ([]byte, []int) {
	new_sequence := sequence
	for i := 0; i < 4; i++ {
		new_sequence = append(new_sequence, '?')
		new_sequence = append(new_sequence, sequence...)
	}
	new_sizes := []int{}
	for i := 0; i < 5; i++ {
		new_sizes = append(new_sizes, sizes...)
	}
	return new_sequence, new_sizes
}

func total_sum(data []string) (int, int) {
	sum := 0
	unfolded_sum := 0

	var wg sync.WaitGroup

	for _, line := range data {
		parts := strings.Split(line, " ")
		sequence := []byte(parts[0])
		size_strs := strings.Split(parts[1], ",")
		sizes := []int{}
		for _, size_str := range size_strs {
			n, _ := strconv.Atoi(size_str)
			sizes = append(sizes, n)
		}

		wg.Add(1)

		go func() {
			defer wg.Done()

			s := get_sum(sequence, sizes)
			sum += s

			if sequence[0] != '?' && sequence[len(sequence)-1] != '?' {
				unfolded := int(math.Pow(float64(s), 5) * 16)
				unfolded_sum += unfolded
			} else {
				unfolded_sequence, unfolded_sizes := unfold(sequence, sizes)
				unfolded_sum += get_sum(unfolded_sequence, unfolded_sizes)
			}
		}()
	}

	wg.Wait()

	return sum, unfolded_sum
}

func main() {
	data := load_data("input.txt")
	sum, unfolded_sum := total_sum(data)
	fmt.Println(sum)
	fmt.Println(unfolded_sum)
}
