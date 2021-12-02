package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	amount    int
}

func get_data() []instruction {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	instructions := []instruction{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])
		ins := instruction{direction: parts[0], amount: amount}
		instructions = append(instructions, ins)
	}
	return instructions
}

func main() {
	instructions := get_data()

	// part 1
	pos := 0
	depth := 0
	for _, instruction := range instructions {
		if instruction.direction == "forward" {
			pos += instruction.amount
		} else if instruction.direction == "down" {
			depth += instruction.amount
		} else if instruction.direction == "up" {
			depth -= instruction.amount
		}
	}
	fmt.Println(pos * depth)

	// part 2
	pos = 0
	depth = 0
	aim := 0
	for _, instruction := range instructions {
		if instruction.direction == "forward" {
			pos += instruction.amount
			depth += instruction.amount * aim
		} else if instruction.direction == "down" {
			aim += instruction.amount
		} else if instruction.direction == "up" {
			aim -= instruction.amount
		}
	}
	fmt.Println(pos * depth)
}
