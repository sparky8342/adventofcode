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
	category       byte
	min            int
	max            int
	min_fail       int
	max_fail       int
	destination    string
}

type Workflow struct {
	name  string
	rules []Rule
	def   string
}

type Part struct {
	categories map[byte]int
}

type State struct {
	name    string
	rule_no int
	min     map[byte]int
	max     map[byte]int
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

			var min, max, min_fail, max_fail int
			if data_parts2[0][1] == '<' {
				min = 1
				max = amount - 1
				min_fail = amount
				max_fail = 4000
			} else {
				min = amount + 1
				max = 4000
				min_fail = 1
				max_fail = amount
			}

			rules = append(rules, Rule{
				category:       data_parts2[0][0],
				min:            min,
				max:            max,
				min_fail:       min_fail,
				max_fail:       max_fail,
				destination:    data_parts2[1],
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
			if part.categories[rule.category] >= rule.min && part.categories[rule.category] <= rule.max {
				name = rule.destination
				break
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

func copy_map(mp map[byte]int) map[byte]int {
	cp := map[byte]int{}
	for k, v := range mp {
		cp[k] = v
	}
	return cp
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func find_paths(workflows map[string]Workflow) int {
	start := State{
		name:    "in",
		rule_no: 0,
		min:     map[byte]int{'x': 1, 'm': 1, 'a': 1, 's': 1},
		max:     map[byte]int{'x': 4000, 'm': 4000, 'a': 4000, 's': 4000},
	}

	sum := 0

	queue := []State{start}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.name == "A" {
			s := 1
			s *= (state.max['x'] - state.min['x'] + 1)
			s *= (state.max['m'] - state.min['m'] + 1)
			s *= (state.max['a'] - state.min['a'] + 1)
			s *= (state.max['s'] - state.min['s'] + 1)
			sum += s
			continue
		} else if state.name == "R" {
			continue
		}

		workflow := workflows[state.name]
		rules := workflow.rules
		rule := rules[state.rule_no]

		// success
		next_min := copy_map(state.min)
		next_min[rule.category] = max(next_min[rule.category], rule.min)

		next_max := copy_map(state.max)
		next_max[rule.category] = min(next_max[rule.category], rule.max)

		next_state := State{
			name:    rule.destination,
			rule_no: 0,
			min:     next_min,
			max:     next_max,
		}
		queue = append(queue, next_state)

		// fail
		next_min = copy_map(state.min)
		next_min[rule.category] = max(next_min[rule.category], rule.min_fail)

		next_max = copy_map(state.max)
		next_max[rule.category] = min(next_max[rule.category], rule.max_fail)

		if state.rule_no == len(rules)-1 {
			next_state = State{
				name:    workflow.def,
				rule_no: 0,
				min:     next_min,
				max:     next_max,
			}
		} else {
			next_state = State{
				name:    state.name,
				rule_no: state.rule_no + 1,
				min:     next_min,
				max:     next_max,
			}
		}
		queue = append(queue, next_state)
	}

	return sum
}

func main() {
	data := load_data("input.txt")
	workflows, parts := parse_data(data)
	fmt.Println(process_parts(workflows, parts))
	fmt.Println(find_paths(workflows))
}
