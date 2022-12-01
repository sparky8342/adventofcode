package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func most_calories(data string) (int, int) {
	sections := strings.Split(data, "\n\n")

	elves := []int{}
	for _, section := range sections {
		calories := 0
		lines := strings.Split(section, "\n")
		for _, line := range lines {
			num, _ := strconv.Atoi(line)
			calories += num
		}
		elves = append(elves, calories)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	return elves[0], elves[0] + elves[1] + elves[2]
}

func main() {
	data := load_data("input.txt")
	most, top_three := most_calories(data)
	fmt.Println(most)
	fmt.Println(top_three)
}
