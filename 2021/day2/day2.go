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

type sub struct {
	pos   int
	depth int
	aim   int
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

	part1 := sub{pos: 0, depth: 0}
	part2 := sub{pos: 0, depth: 0, aim: 0}

	for _, instruction := range instructions {
		switch instruction.direction {
		case "forward":
			{
				part1.pos += instruction.amount
				part2.pos += instruction.amount
				part2.depth += instruction.amount * part2.aim
			}
		case "down":
			{
				part1.depth += instruction.amount
				part2.aim += instruction.amount
			}
		case "up":
			{
				part1.depth -= instruction.amount
				part2.aim -= instruction.amount
			}
		}
	}
	fmt.Println(part1.pos * part1.depth)
	fmt.Println(part2.pos * part2.depth)
}
