package day10

import (
	"fmt"
	"loader"
	"strings"
)

type Machine struct {
	goal    int
	buttons []int
}

type Entry struct {
	state   int
	presses int
}

func parse_data(data []string) []Machine {
	machines := []Machine{}

	for _, line := range data {
		parts := strings.Split(line, " ")

		str := parts[0][1 : len(parts[0])-1]
		goal := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '#' {
				goal = goal | (1 << i)
			}
		}

		buttons := []int{}
		for i := 1; i < len(parts)-1; i++ {
			str := parts[i][1 : len(parts[i])-1]
			num_strs := strings.Split(str, ",")
			button := 0
			for _, n_str := range num_strs {
				button = button | (1 << int(n_str[0]-'0'))
			}
			buttons = append(buttons, button)
		}

		machines = append(machines, Machine{
			goal:    goal,
			buttons: buttons,
		})
	}

	return machines
}

func bfs(machine Machine) int {
	queue := []Entry{Entry{}}
	visited := map[int]struct{}{}
	visited[0] = struct{}{}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		if entry.state == machine.goal {
			return entry.presses
		}

		for _, button := range machine.buttons {
			new_state := entry.state ^ button
			if _, ok := visited[new_state]; !ok {
				queue = append(queue, Entry{
					state:   new_state,
					presses: entry.presses + 1,
				})
				visited[new_state] = struct{}{}
			}
		}
	}

	return -1
}

func presses_needed(machines []Machine) int {
	presses := 0
	for _, machine := range machines {
		presses += bfs(machine)
	}
	return presses
}

func Run() {
	loader.Day = 10
	data := loader.GetStrings()
	machines := parse_data(data)
	part1 := presses_needed(machines)

	fmt.Printf("%d %d\n", part1, 0)
}
