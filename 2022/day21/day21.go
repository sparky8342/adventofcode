package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Monkey struct {
	name          string
	number_monkey bool
	val           int
	left          *Monkey
	right         *Monkey
	operator      string
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) *Monkey {
	monkeys := map[string]*Monkey{}
	left := map[string]string{}
	right := map[string]string{}

	for _, line := range data {
		monkey := &Monkey{}

		parts := strings.Split(line, ": ")

		monkey.name = parts[0]

		values := strings.Split(parts[1], " ")
		if len(values) == 1 {
			num, _ := strconv.Atoi(values[0])
			monkey.number_monkey = true
			monkey.val = num
		} else {
			left[monkey.name] = values[0]
			monkey.operator = values[1]
			right[monkey.name] = values[2]
		}
		monkeys[monkey.name] = monkey
	}

	for _, monkey := range monkeys {
		monkey.left = monkeys[left[monkey.name]]
		monkey.right = monkeys[right[monkey.name]]
	}

	return monkeys["root"]
}

func eval(monkey *Monkey) int {
	if monkey.number_monkey {
		return monkey.val
	}

	left := eval(monkey.left)
	right := eval(monkey.right)

	switch monkey.operator {
	case "+":
		return left + right
	case "*":
		return left * right
	case "-":
		return left - right
	case "/":
		return left / right
	}

	return -1
}

func main() {
	data := load_data("input.txt")
	root := parse_data(data)
	fmt.Println(eval(root))
}
