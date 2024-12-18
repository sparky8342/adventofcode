package day17

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type Computer struct {
	a       int
	b       int
	c       int
	program []int
	pc      int
	output  []int
}

func parse_data(data []string) Computer {
	computer := Computer{}

	_, err := fmt.Sscanf(data[0], "Register A: %d", &computer.a)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Sscanf(data[1], "Register B: %d", &computer.b)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Sscanf(data[2], "Register C: %d", &computer.c)
	if err != nil {
		panic(err)
	}

	for _, str := range strings.Split(data[4][9:], ",") {
		computer.program = append(computer.program, int(str[0]-'0'))
	}

	return computer
}

func (c *Computer) combo_operand(op int) int {
	if op >= 0 && op <= 3 {
		return op
	} else if op == 4 {
		return c.a
	} else if op == 5 {
		return c.b
	} else if op == 6 {
		return c.c
	}
	return -1
}

func (c *Computer) run_program() string {
	for c.pc >= 0 && c.pc < len(c.program) {
		ins := c.program[c.pc]

		var operand int
		if ins == 0 || ins == 2 || ins == 5 || ins == 6 || ins == 7 {
			operand = c.combo_operand(c.program[c.pc+1])
		} else {
			operand = c.program[c.pc+1]
		}

		switch ins {
		case 0:
			c.a = c.a / (1 << operand)
		case 1:
			c.b ^= operand
		case 2:
			c.b = operand % 8
		case 3:
			if c.a != 0 {
				c.pc = operand
				continue
			}
		case 4:
			c.b ^= c.c
		case 5:
			c.output = append(c.output, operand%8)
		case 6:
			c.b = c.a / (1 << operand)
		case 7:
			c.c = c.a / (1 << operand)
		}

		c.pc += 2
	}

	out := []string{}
	for _, n := range c.output {
		out = append(out, strconv.Itoa(n))
	}
	return strings.Join(out, ",")
}

func (c *Computer) reset() {
	c.a = 0
	c.b = 0
	c.c = 0
	c.pc = 0
	c.output = []int{}
}

func find_quine(computer Computer) int {
	nums := []int{0}
	pos := len(computer.program) - 1

	for pos >= 0 {
		digit := computer.program[pos]

		next_nums := []int{}
		for _, n := range nums {
			for i := 0; i < 8; i++ {
				computer.reset()
				a := n*8 + i
				computer.a = a
				_ = computer.run_program()
				if computer.output[0] == digit {
					next_nums = append(next_nums, a)
				}
			}
		}

		pos--
		nums = next_nums
	}

	return nums[0]
}

func Run() {
	loader.Day = 17
	data := loader.GetStrings()
	computer := parse_data(data)

	part1 := computer.run_program()
	part2 := find_quine(computer)

	fmt.Printf("%s %d\n", part1, part2)
}
