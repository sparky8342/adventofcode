package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func get_data() []int {
	data, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSuffix(string(data), "\n")
	parts := strings.Split(line, ",")
	numbers := []int{}
	for _, num_str := range parts {
		num, _ := strconv.Atoi(num_str)
		numbers = append(numbers, num)
	}
	return numbers
}

func day_passes(fish []int) {
	end_of_cycle := fish[0]
	for i := 0; i < 8; i++ {
		fish[i] = fish[i+1]
	}
	fish[8] = end_of_cycle
	fish[6] += end_of_cycle
}

func sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {
	fish_nums := get_data()

	fish := make([]int, 9)
	for _, num := range fish_nums {
		fish[num]++
	}

	part1_cycles := 79
	part2_cycles := 256

	for i := 0; i < part2_cycles; i++ {
		day_passes(fish)
		if i == part1_cycles {
			fmt.Println(sum(fish))
		}
	}
	fmt.Println(sum(fish))
}
