package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var incremental_moves []int

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

func max(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func populate_incremental_moves(max int) {
	incremental_moves = make([]int, max)
	cost := 0
	step_cost := 1
	for i := 1; i < max; i++ {
		cost += step_cost
		step_cost++
		incremental_moves[i] = cost
	}
}

func main() {
	crabs := get_data()

	max_val := max(crabs)
	populate_incremental_moves(max_val + 1)

	best_fuel := -1
	best_fuel_inc := -1
	for target := 0; target < max_val; target++ {
		fuel := 0
		fuel_inc := 0
		for _, crab := range crabs {
			fuel += abs(crab - target)
			fuel_inc += incremental_moves[abs(crab-target)]
		}
		if best_fuel == -1 || fuel < best_fuel {
			best_fuel = fuel
		}
		if best_fuel_inc == -1 || fuel_inc < best_fuel_inc {
			best_fuel_inc = fuel_inc
		}
	}

	fmt.Println(best_fuel)
	fmt.Println(best_fuel_inc)
}
