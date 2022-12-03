package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Empty struct {
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func priority_sum(data []string) int {
	sum := 0

	for _, line := range data {
		letters := map[byte]Empty{}
		for i := 0; i < len(line)/2; i++ {
			letters[line[i]] = Empty{}
		}

		for i := len(line) / 2; i < len(line); i++ {
			if _, seen := letters[line[i]]; seen {
				if line[i] >= 'a' {
					sum += int(line[i]) - 'a' + 1
				} else {
					sum += int(line[i]) - 'A' + 27
				}
				break
			}
		}
	}
	return sum
}

func badge_sum(data []string) int {
	sum := 0

	for i := 0; i < len(data); i += 3 {
		letter_sets := []map[byte]Empty{}
		for j := 0; j < 3; j++ {
			line := data[i+j]
			letter_set := map[byte]Empty{}
			for k := 0; k < len(line); k++ {
				letter_set[line[k]] = Empty{}
			}
			letter_sets = append(letter_sets, letter_set)
		}
		for letter, _ := range letter_sets[0] {
			if _, seen := letter_sets[1][letter]; seen {
				if _, seen2 := letter_sets[2][letter]; seen2 {
					if letter >= 'a' {
						sum += int(letter) - 'a' + 1
					} else {
						sum += int(letter) - 'A' + 27
					}
				}
			}
		}
	}
	return sum
}

func main() {
	data := load_data("input.txt")
	sum := priority_sum(data)
	fmt.Println(sum)
	sum = badge_sum(data)
	fmt.Println(sum)
}
