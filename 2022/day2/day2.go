package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func get_score(data []string) int {
	scoring := map[byte]map[byte]int{
		'A': map[byte]int{'X': 4, 'Y': 8, 'Z': 3},
		'B': map[byte]int{'X': 1, 'Y': 5, 'Z': 9},
		'C': map[byte]int{'X': 7, 'Y': 2, 'Z': 6},
	}

	score := 0
	for _, line := range data {
		opponent := line[0]
		player := line[2]
		score += scoring[opponent][player]
	}

	return score
}

func get_score_with_strategy(data []string) int {
	scoring := map[byte]map[byte]int{
		'A': map[byte]int{'X': 3, 'Y': 4, 'Z': 8},
		'B': map[byte]int{'X': 1, 'Y': 5, 'Z': 9},
		'C': map[byte]int{'X': 2, 'Y': 6, 'Z': 7},
	}

	score := 0
	for _, line := range data {
		opponent := line[0]
		player := line[2]
		score += scoring[opponent][player]
	}

	return score
}

func main() {
	data := load_data("input.txt")

	score := get_score(data)
	fmt.Println(score)

	score = get_score_with_strategy(data)
	fmt.Println(score)
}
