package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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
	} else {
		for i, n := range found {
			if n > sizes[i] {
				return false
			}
		}
	}

	return true
}

func dfs(sequence []byte, sizes []int, valid *int) {
	if !is_valid(sequence, sizes) {
		return
	}

	wildcard := -1
	for i := 0; i < len(sequence); i++ {
		if sequence[i] == '?' {
			wildcard = i
			break
		}
	}
	if wildcard == -1 {
		*valid++
		return
	}

	next := make([]byte, len(sequence))
	copy(next, sequence)
	next[wildcard] = '.'
	dfs(next, sizes, valid)

	next = make([]byte, len(sequence))
	copy(next, sequence)
	next[wildcard] = '#'
	dfs(next, sizes, valid)
}

func get_sum(sequence []byte, sizes []int) int {
	valid := 0
	dfs(sequence, sizes, &valid)
	return valid
}

func total_sum(data []string) int {
	sum := 0
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
	}
	return sum
}

func main() {
	data := load_data("input.txt")
	total_sum := total_sum(data)
	fmt.Println(total_sum)
}
