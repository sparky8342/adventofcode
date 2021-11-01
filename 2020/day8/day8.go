package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Ins struct {
	op  string
	arg int
}

type Program struct {
	acm          int
	pos          int
	instructions []Ins
}

func get_program() Program {
	instructions := []Ins{}

	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		op := parts[0]
		arg, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, Ins{op: op, arg: arg})
	}
	p := Program{instructions: instructions}
	return p
}

func run(program *Program) (bool, int) {
	program.acm = 0
	program.pos = 0
	seen := make(map[int]bool)

	for program.pos >= 0 && program.pos < len(program.instructions) {
		if seen[program.pos] {
			return true, program.acm
		}
		seen[program.pos] = true

		ins := program.instructions[program.pos]

		switch ins.op {
		case "acc":
			{
				program.acm += ins.arg
				program.pos++
			}
		case "jmp":
			{
				program.pos += ins.arg
			}
		case "nop":
			{
				program.pos++
			}
		}
	}

	return false, program.acm
}

func main() {
	program := get_program()

	// part 1
	_, acm := run(&program)
	fmt.Println(acm)

	// part 2
	for i, _ := range program.instructions {
		ins := program.instructions[i]

		switch ins.op {
		case "nop":
			{
				program.instructions[i].op = "jmp"
			}
		case "jmp":
			{
				program.instructions[i].op = "nop"
			}
		default:
			{
				continue
			}
		}

		looped, acm := run(&program)

		if looped == false {
			fmt.Println(acm)
			break
		}

		program.instructions[i].op = ins.op
	}
}
