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

func calibration(data []string, find_words bool) int {
	words := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var first, last, total int

	for _, line := range data {

		// first digit
	outer:
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = int(line[i]-'0')
				break
			} else if find_words {
				for j := i + 2; j <= i+5 && j < len(line); j++ {
					sub := line[i:j]
					if val, exists := words[sub]; exists {
						first = val
						break outer
					}
				}
			}
		}

		// last digit
	outer2:
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = int(line[i] - '0')
				break
			} else if find_words {
				for j := i - 2; j >= i-4 && j >= 0; j-- {
					sub := line[j : i+1]
					if val, exists := words[sub]; exists {
						last = val
						break outer2
					}
				}
			}
		}

		total += (first * 10) + last
	}

	return total
}

func main() {
	data := load_data("input.txt")
	fmt.Println(calibration(data, false))
	fmt.Println(calibration(data, true))
}
