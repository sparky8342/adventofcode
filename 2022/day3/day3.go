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
			fmt.Printf(string(line[i]))
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

func main() {
	data := load_data("input.txt")
	sum := priority_sum(data)
	fmt.Println(sum)
}
