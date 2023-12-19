package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

type Rule struct {
	category    byte
	less_than   bool
	amount      int
	destination string
}

type Workflow struct {
	name  string
	rules []Rule
	def   string
}

type Part struct {
	categories map[byte]int
}

func parse_data(data string) (map[string]Workflow, []Part) {
	workflows := map[string]Workflow{}

	data_parts := strings.Split(data, "\n\n")

	for _, line := range strings.Split(data_parts[0], "\n") {

		workflow := Workflow{}

		pos := 0
		for line[pos] != '{' {
			pos++
		}
		workflow.name = line[0:pos]

		rule_strs := strings.Split(line[pos+1:len(line)-1], ",")

		rules := []Rule{}

		for i := 0; i < len(rule_strs)-1; i++ {
			data_parts2 := strings.Split(rule_strs[i], ":")
			amount, _ := strconv.Atoi(data_parts2[0][2:])

			rules = append(rules, Rule{
				category:    data_parts2[0][0],
				less_than:   data_parts2[0][1] == '<',
				amount:      amount,
				destination: data_parts2[1],
			})
		}

		workflow.rules = rules
		workflow.def = rule_strs[len(rule_strs)-1]

		workflows[workflow.name] = workflow
	}

	parts := []Part{}

	for _, line := range strings.Split(data_parts[1], "\n") {
		part := Part{categories: map[byte]int{}}

		for _, cat_str := range strings.Split(line[1:len(line)-1], ",") {
			cat_parts := strings.Split(cat_str, "=")
			amount, _ := strconv.Atoi(cat_parts[1])
			part.categories[cat_parts[0][0]] = amount
		}

		parts = append(parts, part)
	}

	return workflows, parts
}

func process_part(workflows map[string]Workflow, part Part) int {
	name := "in"

	for {
		workflow := workflows[name]
		name = workflow.def
		for _, rule := range workflow.rules {
			if rule.less_than {
				if part.categories[rule.category] < rule.amount {
					name = rule.destination
					break
				}
			} else {
				if part.categories[rule.category] > rule.amount {
					name = rule.destination
					break
				}
			}
		}

		if name == "A" {
			return part.categories['x'] + part.categories['m'] + part.categories['a'] + part.categories['s']
		} else if name == "R" {
			return 0
		}
	}
}

func process_parts(workflows map[string]Workflow, parts []Part) int {
	sum := 0
	for _, part := range parts {
		sum += process_part(workflows, part)
	}
	return sum
}

func main() {
	data := load_data("input.txt")
	workflows, parts := parse_data(data)
	fmt.Println(process_parts(workflows, parts))

	//228141 too low
}
