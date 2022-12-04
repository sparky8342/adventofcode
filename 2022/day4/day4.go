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

type Range struct {
	start int
	end   int
}

type Pair struct {
	range1 Range
	range2 Range
}

func get_pairs(data []string) []Pair {
	pairs := []Pair{}

	for _, line := range data {
		parts := strings.Split(line, ",")
		range1_parts := strings.Split(parts[0], "-")
		range2_parts := strings.Split(parts[1], "-")

		range1_a, _ := strconv.Atoi(range1_parts[0])
		range1_b, _ := strconv.Atoi(range1_parts[1])
		range2_a, _ := strconv.Atoi(range2_parts[0])
		range2_b, _ := strconv.Atoi(range2_parts[1])

		range1 := Range{start: range1_a, end: range1_b}
		range2 := Range{start: range2_a, end: range2_b}

		if range1.start > range2.start {
			range1, range2 = range2, range1
		}

		pairs = append(pairs, Pair{range1: range1, range2: range2})
	}

	return pairs
}

func range_contains(pairs []Pair) int {
	contains := 0
	for _, pair := range pairs {
		if (pair.range1.start <= pair.range2.start && pair.range1.end >= pair.range2.end) || (pair.range2.start <= pair.range1.start && pair.range2.end >= pair.range1.end) {
			contains++
		}
	}
	return contains
}

func range_overlaps(pairs []Pair) int {
	overlaps := 0
	for _, pair := range pairs {
		if pair.range1.end >= pair.range2.start {
			overlaps++
		}
	}
	return overlaps
}

func main() {
	data := load_data("input.txt")
	pairs := get_pairs(data)

	contains := range_contains(pairs)
	fmt.Println(contains)

	overlaps := range_overlaps(pairs)
	fmt.Println(overlaps)
}
