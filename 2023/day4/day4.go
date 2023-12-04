package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

func points(data []string) int {
	re := regexp.MustCompile("Card\\s+(\\d+): (.*?) \\| (.*)")

	score := 0

	for _, line := range data {
		match := re.FindStringSubmatch(line)

		//id, _ := strconv.Atoi(match[1])

		winning := map[int]struct{}{}

		for _, num_str := range strings.Fields(match[2]) {
			num, _ := strconv.Atoi(num_str)
			winning[num] = struct{}{}
		}

		card_score := 0

		for _, num_str := range strings.Fields(match[3]) {
			num, _ := strconv.Atoi(num_str)
			if _, exists := winning[num]; exists {
				if card_score == 0 {
					card_score = 1
				} else {
					card_score *= 2
				}
			}
		}

		score += card_score
	}

	return score
}

func main() {
	data := load_data("input.txt")
	fmt.Println(points(data))
}
