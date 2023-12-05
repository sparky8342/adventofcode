package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Mapping struct {
	destination int
	source      int
	length      int
	diff        int
}

type Conversion struct {
	mappings []Mapping
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) ([]int, []Conversion) {
	parts := strings.Fields(data[0])
	seeds := []int{}
	for i := 1; i < len(parts); i++ {
		n, _ := strconv.Atoi(parts[i])
		seeds = append(seeds, n)
	}

	conversions := []Conversion{}
	line_no := 3

	for line_no < len(data) {
		conversion := Conversion{}
		for line_no < len(data) && data[line_no] != "" {
			parts := strings.Fields(data[line_no])
			destination, _ := strconv.Atoi(parts[0])
			source, _ := strconv.Atoi(parts[1])
			length, _ := strconv.Atoi(parts[2])
			mapping := Mapping{destination: destination, source: source, length: length, diff: destination - source}
			conversion.mappings = append(conversion.mappings, mapping)
			line_no++
		}
		conversions = append(conversions, conversion)
		line_no += 2
	}

	return seeds, conversions
}

func convert_seed(seed int, conversions []Conversion) int {
	for _, conversion := range conversions {
		for _, mapping := range conversion.mappings {
			if mapping.source <= seed && seed < mapping.source+mapping.length {
				seed += mapping.diff
				break
			}
		}
	}
	return seed
}

func lowest_location(data []string) (int, int) {
	seeds, conversions := parse_data(data)

	lowest := convert_seed(seeds[0], conversions)
	lowest_from_range := lowest

	for i := 1; i < len(seeds); i++ {
		location := convert_seed(seeds[i], conversions)
		if location < lowest {
			lowest = location
		}
	}

	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			location := convert_seed(j, conversions)
			if location < lowest_from_range {
				lowest_from_range = location
			}
		}
	}

	return lowest, lowest_from_range
}

func main() {
	data := load_data("input.txt")
	lowest, lowest_from_range := lowest_location(data)
	fmt.Println(lowest)
	fmt.Println(lowest_from_range)
}
