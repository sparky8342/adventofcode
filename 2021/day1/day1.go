package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func get_data() []int {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	numbers := []int{}
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}
	return numbers
}

func main() {
	numbers := get_data()
	part1 := 0
	part2 := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			part1++
		}
		if i < len(numbers)-3 {
			sum1 := numbers[i] + numbers[i+1] + numbers[i+2]
			sum2 := numbers[i+1] + numbers[i+2] + numbers[i+3]
			if sum2 > sum1 {
				part2++
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
