package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	patterns     []string
	output_value []string
}

var digit_groups map[string]int
var wires []string

func init() {
	digit_groups = map[string]int{
		"012456":  0,
		"25":      1,
		"02346":   2,
		"02356":   3,
		"1235":    4,
		"01356":   5,
		"013456":  6,
		"025":     7,
		"0123456": 8,
		"012356":  9,
	}

	wires = []string{"a", "b", "c", "d", "e", "f", "g"}
}

func get_data() []Entry {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	entries := []Entry{}
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		patterns := strings.Split(parts[0], " ")
		output_value := strings.Split(parts[1], " ")
		entries = append(entries, Entry{patterns: patterns, output_value: output_value})
	}
	return entries
}

func is_valid(entry Entry, solution [7]string) bool {
	letters := map[string]int{}
	for _, letter := range solution {
		if letter != "" {
			letters[letter]++
			if letters[letter] > 1 {
				return false
			}
		}
	}

	for _, pattern := range append(entry.patterns, entry.output_value...) {
		switch len(pattern) {
		case 2: //segment 2 and 5
			for _, ch := range pattern {
				for i := 0; i <= 6; i++ {
					if solution[i] == string(ch) && i != 2 && i != 5 {
						return false
					}
				}
			}
		case 3: // 0, 2, 5
			for _, ch := range pattern {
				for i := 0; i <= 6; i++ {
					if solution[i] == string(ch) && i != 0 && i != 2 && i != 5 {
						return false
					}
				}
			}
		case 4: // 1, 2, 3, 5
			for _, ch := range pattern {
				for i := 0; i <= 6; i++ {
					if solution[i] == string(ch) && i != 1 && i != 2 && i != 3 && i != 5 {
						return false
					}
				}
			}
		}
	}

	if solution[6] != "" {
		// check for valid display digits
		for _, pattern := range append(entry.patterns, entry.output_value...) {
			segments_on := []int{}
			for _, ch := range pattern {
				for i := 0; i <= 6; i++ {
					if solution[i] == string(ch) {
						segments_on = append(segments_on, i)
					}
				}
			}
			sort.Ints(segments_on)
			str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(segments_on)), ""), "[]")
			_, found := digit_groups[str]
			if !found {
				return false
			}
		}

	}
	return true
}

func backtrack(entry Entry, solution [7]string, segment_no int) (bool, [7]string) {
	/*
		 00
		1  2
		1  2
		 33
		4  5
		4  5
		 66
	*/

	if segment_no == 7 {
		return true, solution
	}

	for _, letter := range wires {
		solution[segment_no] = letter
		if is_valid(entry, solution) {
			done, solution_found := backtrack(entry, solution, segment_no+1)
			if done {
				return done, solution_found
			}
		}
	}
	solution[segment_no] = ""
	return false, [7]string{}
}

func translate_numbers(entry Entry, solution [7]string) int {
	output_digits := []int{}

	for _, str := range entry.output_value {
		digits := []int{}
		for _, ch := range str {
			for i := 0; i <= 6; i++ {
				if solution[i] == string(ch) {
					digits = append(digits, i)
				}
			}
		}
		sort.Ints(digits)
		str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(digits)), ""), "[]")
		output_digits = append(output_digits, digit_groups[str])
	}

	s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output_digits)), ""), "[]")
	num, _ := strconv.Atoi(s)
	return num
}

func main() {
	entries := get_data()

	// part 1
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

	// part 2
	total := 0
	for _, entry := range entries {
		_, solution := backtrack(entry, [7]string{}, 0)
		total += translate_numbers(entry, solution)
	}
	fmt.Println(total)
}
