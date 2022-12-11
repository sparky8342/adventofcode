package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id            int
	items         []int
	multiply_self bool
	multiply      int
	add           int
	test          int
	True          int
	False         int
	inspected     int
}

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func parse_data(data string) ([]*Monkey, int) {
	mod_no := 1

	monkeys := []*Monkey{}

	sections := strings.Split(data, "\n\n")
	for _, section := range sections {
		lines := strings.Split(section, "\n")

		monkey := Monkey{}

		parts := strings.Split(lines[0], " ")
		monkey.id = int(parts[1][0]) - '0'

		parts = strings.Split(lines[1], " ")
		for i := 4; i < len(parts); i++ {
			str := parts[i]
			if str[len(str)-1] == ',' {
				str = str[:len(str)-1]
			}
			n, _ := strconv.Atoi(str)
			monkey.items = append(monkey.items, n)
		}

		parts = strings.Split(lines[2], " ")
		operand_str := parts[7]
		if operand_str == "old" {
			monkey.multiply_self = true
		} else {
			operand, _ := strconv.Atoi(parts[7])
			if parts[6] == "*" {
				monkey.multiply = operand
				monkey.add = 0
			} else if parts[6] == "+" {
				monkey.multiply = 1
				monkey.add = operand
			}
		}

		parts = strings.Split(lines[3], " ")
		monkey.test, _ = strconv.Atoi(parts[5])

		mod_no *= monkey.test

		parts = strings.Split(lines[4], " ")
		monkey.True, _ = strconv.Atoi(parts[9])

		parts = strings.Split(lines[5], " ")
		monkey.False, _ = strconv.Atoi(parts[9])

		monkeys = append(monkeys, &monkey)
	}

	return monkeys, mod_no
}

func one_round(monkeys []*Monkey, divide bool, mod_no int) {
	for _, monkey := range monkeys {
		for _, item := range monkey.items {
			if monkey.multiply_self {
				item = item * item
			} else {
				item = item*monkey.multiply + monkey.add
			}
			if divide {
				item /= 3
			}

			item = item % mod_no

			var dest_monkey *Monkey
			if item%monkey.test == 0 {
				dest_monkey = monkeys[monkey.True]
			} else {
				dest_monkey = monkeys[monkey.False]
			}
			dest_monkey.items = append(dest_monkey.items, item)

			monkey.inspected++
		}
		monkey.items = []int{}
	}
}

func main() {
	data := load_data("input.txt")

	for part := 1; part <= 2; part++ {
		divide := true
		rounds := 20
		if part == 2 {
			divide = false
			rounds = 10000
		}

		monkeys, mod_no := parse_data(data)

		for i := 0; i < rounds; i++ {
			one_round(monkeys, divide, mod_no)
		}

		sort.Slice(monkeys, func(i, j int) bool {
			return monkeys[i].inspected > monkeys[j].inspected
		})

		fmt.Println(monkeys[0].inspected * monkeys[1].inspected)
	}
}
