package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Crt [6][40]bool

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

func set_pixel(crt *Crt, cycle int, x int) {
	if cycle > 240 {
		return
	}

	row := (cycle - 1) / 40
	col := (cycle - 1) % 40

	if col >= x-1 && col <= x+1 {
		crt[row][col] = true
	} else {
		crt[row][col] = false
	}
}

func run_program_part2(program []string) Crt {
	crt := Crt{}
	x := 1
	cycle := 1
	set_pixel(&crt, cycle, x)

	for _, line := range program {
		parts := strings.Split(line, " ")
		if parts[0] == "noop" {
			cycle++
			set_pixel(&crt, cycle, x)
		} else {
			val, _ := strconv.Atoi(parts[1])
			cycle++
			set_pixel(&crt, cycle, x)
			cycle++
			x += val
			set_pixel(&crt, cycle, x)
		}
	}

	return crt
}

func print_crt(crt *Crt) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if crt[i][j] == true {
				fmt.Print("\u2B1C")
			} else {
				fmt.Print("\u2B1B")
			}
		}
		fmt.Println()
	}
}

func main() {
	data := load_data("input.txt")
	sum := run_program(data)
	fmt.Println(sum)

	crt := run_program_part2(data)
	print_crt(&crt)
}
