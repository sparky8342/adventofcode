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

func contains_at(str string, str2 string, pos int) bool {
	if len(str2)+pos > len(str) {
		return false
	}
	for i := 0; i < len(str2); i++ {
		if str[pos+i] != str2[i] {
			return false
		}
	}
	return true
}

func calibration(data []string) (int, int) {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	total := 0
	total2 := 0

	for _, line := range data {
		digits := []int{}
		digits2 := []int{}

		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				n := int(line[i] - '0')
				digits = append(digits, n)
				digits2 = append(digits2, n)
				continue
			}
			for word_index, word := range words {
				if contains_at(line, word, i) {
					digits2 = append(digits2, word_index+1)
					break
				}
			}

		}

		if len(digits) > 0 {
			total += digits[0]*10 + digits[len(digits)-1]
		}
		if len(digits2) > 0 {
			total2 += digits2[0]*10 + digits2[len(digits2)-1]
		}
	}

	return total, total2
}

func main() {
	data := load_data("input.txt")
	total, total2 := calibration(data)
	fmt.Println(total)
	fmt.Println(total2)
}
