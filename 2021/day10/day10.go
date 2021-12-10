package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func get_data() []string {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	return lines
}

func main() {
	left := map[string]bool{"(": true, "[": true, "{": true, "<": true}
	right := map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
	errors := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	complete := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}

	lines := get_data()

	error_score := 0             // part 1
	completion_scores := []int{} // part2

	for _, line := range lines {
		stack := []string{}
		for _, ru := range line {
			char := string(ru)
			if _, found := left[char]; found {
				stack = append(stack, char)
			} else {
				matching := right[char]
				if stack[len(stack)-1] != matching {
					error_score += errors[char]
					stack = []string{}
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}

		// part 2
		if len(stack) > 0 {
			complete_score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				complete_score *= 5
				complete_score += complete[stack[i]]
			}
			completion_scores = append(completion_scores, complete_score)
		}
	}

	fmt.Println(error_score)

	sort.Ints(completion_scores)
	fmt.Println(completion_scores[len(completion_scores)/2])
}
