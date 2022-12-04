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

func range_contains(data []string) int {
	contains := 0

	for _, line := range data {
		parts := strings.Split(line, ",")
		range1 := strings.Split(parts[0], "-")
		range2 := strings.Split(parts[1], "-")

		range1_a, _ := strconv.Atoi(range1[0])
		range1_b, _ := strconv.Atoi(range1[1])
		range2_a, _ := strconv.Atoi(range2[0])
		range2_b, _ := strconv.Atoi(range2[1])

		ranges := []Range{
			Range{start: range1_a, end: range1_b},
			Range{start: range2_a, end: range2_b},
		}

		if (ranges[0].start <= ranges[1].start && ranges[0].end >= ranges[1].end) || (ranges[1].start <= ranges[0].start && ranges[1].end >= ranges[0].end) {
			contains++
		}
	}

	return contains
}

func main() {
	data := load_data("input.txt")
	contains := range_contains(data)
	fmt.Println(contains)
}
