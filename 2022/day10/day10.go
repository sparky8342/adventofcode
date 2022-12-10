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

func check_signal_strength(cycle int, x int) int {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return cycle * x
	}
	return 0
}

func run_program(program []string) int {
	x := 1
	cycle := 1
	sum := 0

	for _, line := range program {
		parts := strings.Split(line, " ")
		if parts[0] == "noop" {
			cycle++
			sum += check_signal_strength(cycle, x)
		} else {
			val, _ := strconv.Atoi(parts[1])
			cycle++
			sum += check_signal_strength(cycle, x)
			cycle++
			x += val
			sum += check_signal_strength(cycle, x)
		}
	}

	return sum
}

func main() {
	data := load_data("input.txt")
	sum := run_program(data)
	fmt.Println(sum)
}
