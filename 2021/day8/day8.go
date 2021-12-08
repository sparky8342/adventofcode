package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type entry struct {
	patterns     []string
	output_value []string
}

func get_data() []entry {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	entries := []entry{}
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		patterns := strings.Split(parts[0], " ")
		output_value := strings.Split(parts[1], " ")
		entries = append(entries, entry{patterns: patterns, output_value: output_value})
	}
	return entries
}

func main() {
	entries := get_data()

	unique_numbers := map[int]bool{2: true, 3: true, 4: true, 7: true}
	count := 0
	for _, entry := range entries {
		for _, segment := range entry.output_value {
			if _, found := unique_numbers[len(segment)]; found {
				count++
			}
		}
	}
	fmt.Println(count)
}
