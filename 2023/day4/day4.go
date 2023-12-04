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

func points(data []string) (int, int) {
	re := regexp.MustCompile(".*?: (.*?) \\| (.*)")

	score := 0

	dp := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		dp[i] = 1
	}

	for i, line := range data {
		match := re.FindStringSubmatch(line)

		winning := map[int]struct{}{}

		for _, num_str := range strings.Fields(match[1]) {
			num, _ := strconv.Atoi(num_str)
			winning[num] = struct{}{}
		}

		matching_numbers := 0
		card_score := 0

		for _, num_str := range strings.Fields(match[2]) {
			num, _ := strconv.Atoi(num_str)
			if _, exists := winning[num]; exists {
				matching_numbers++
				if card_score == 0 {
					card_score = 1
				} else {
					card_score *= 2
				}
			}
		}

		score += card_score

		if matching_numbers > 0 {
			for j := 1; j <= matching_numbers; j++ {
				dp[i+j] += dp[i]
			}
		}

	}

	cards_won := 0
	for _, n := range dp {
		cards_won += n
	}

	return score, cards_won
}

func main() {
	data := load_data("input.txt")
	score, cards_won := points(data)
	fmt.Println(score)
	fmt.Println(cards_won)
}
