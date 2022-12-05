package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	amount      int
	source      int
	destination int
}

type Stacks [][]byte

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func parse_data(data string) (Stacks, []Instruction) {
	parts := strings.Split(data, "\n\n")

	// parse crate setup
	lines := strings.Split(parts[0], "\n")
	lines = lines[:len(lines)-1]

	stacks := Stacks{}

	no_of_stacks := (len(lines[0]) / 4) + 1
	for i := 0; i < no_of_stacks; i++ {
		stacks = append(stacks, []byte{})
	}

	for _, line := range lines {
		for i := 0; i < no_of_stacks; i++ {
			pos := i*4 + 1
			if line[pos] != ' ' {
				stacks[i] = append([]byte{line[pos]}, stacks[i]...)
			}
		}
	}

	/*
		for i := 0; i < no_of_stacks; i++ {
			for _, crate := range stacks[i] {
				fmt.Print(string(crate))
			}
			fmt.Println()
		}
	*/

	// parse instructions
	lines = strings.Split(parts[1], "\n")
	r := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	instructions := []Instruction{}
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		amount, _ := strconv.Atoi(match[1])
		source, _ := strconv.Atoi(match[2])
		destination, _ := strconv.Atoi(match[3])
		instructions = append(instructions, Instruction{
			amount:      amount,
			source:      source,
			destination: destination,
		})
	}

	return stacks, instructions
}

func move_crate(stacks Stacks, source int, destination int) {
	source--
	destination--
	crate := stacks[source][len(stacks[source])-1]
	stacks[source] = stacks[source][:len(stacks[source])-1]
	stacks[destination] = append(stacks[destination], crate)
}

func move(stacks Stacks, instruction Instruction) {
	for i := 1; i <= instruction.amount; i++ {
		move_crate(stacks, instruction.source, instruction.destination)
	}
}

func process_instructions(stacks Stacks, instructions []Instruction) {
	for _, ins := range instructions {
		move(stacks, ins)
	}
}

func top_of_stacks(stacks Stacks) string {
	top := ""
	for i := 0; i < len(stacks); i++ {
		top += string(stacks[i][len(stacks[i])-1])
	}
	return top
}

func main() {
	data := load_data("input.txt")
	stacks, instructions := parse_data(data)

	process_instructions(stacks, instructions)
	top := top_of_stacks(stacks)
	fmt.Println(top)
}
