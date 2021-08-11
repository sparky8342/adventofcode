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

func get_program() []Ins {
	program := []Ins{}

	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		op := parts[0]
		arg, _ := strconv.Atoi(parts[1])
		program = append(program, Ins{op: op, arg: arg})
	}
	return program
}

func run(program []Ins) (bool, int) {
	acm := 0
	pos := 0
	seen := make(map[int]bool)

	for pos >= 0 && pos < len(program) {
		if seen[pos] {
			return true, acm
		}
		seen[pos] = true

		ins := program[pos]

		switch ins.op {
		case "acc":
			{
				acm += ins.arg
				pos++
			}
		case "jmp":
			{
				pos += ins.arg
			}
		case "nop":
			{
				pos++
			}
		}
	}

	return false, acm
}

func main() {
	program := get_program()

	// part 1
	_, acm := run(program)
	fmt.Println(acm)

	// part 2
	for i, _ := range program {
		ins := program[i]

		switch ins.op {
		case "nop":
			{
				program[i].op = "jmp"
			}
		case "jmp":
			{
				program[i].op = "nop"
			}
		default:
			{
				continue
			}
		}

		looped, acm := run(program)

		if looped == false {
			fmt.Println(acm)
			break
		}

		program[i].op = ins.op
	}
}
