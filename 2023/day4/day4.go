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

func points(data []string) (int, int) {
	score := 0

	dp := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		dp[i] = 1
	}

	for i, line := range data {
		parts := strings.Fields(line)

		winning := map[int]struct{}{}

		pos := 2
		for parts[pos] != "|" {
			num, _ := strconv.Atoi(parts[pos])
			winning[num] = struct{}{}
			pos++
		}

		matching_numbers := 0
		card_score := 0

		for j := pos + 1; j < len(parts); j++ {
			num, _ := strconv.Atoi(parts[j])
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
