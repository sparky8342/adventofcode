package day13

import (
	"fmt"
	"loader"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Pos struct {
	x int
	y int
}

type Machine struct {
	a_x    int
	a_y    int
	b_x    int
	b_y    int
	goal_x int
	goal_y int
}

var cache map[[2]int]int

func parse_line(line string) (int, int) {
	r := regexp.MustCompile(".*?(\\d+).*?(\\d+)")
	matches := r.FindStringSubmatch(line)
	n1, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	n2, err := strconv.Atoi(matches[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return n1, n2
}

func parse_data(data [][]string) []Machine {
	machines := []Machine{}
	for _, group := range data {
		a_x, a_y := parse_line(group[0])
		b_x, b_y := parse_line(group[1])
		goal_x, goal_y := parse_line(group[2])
		machine := Machine{
			a_x:    a_x,
			a_y:    a_y,
			b_x:    b_x,
			b_y:    b_y,
			goal_x: goal_x,
			goal_y: goal_y,
		}
		machines = append(machines, machine)
	}
	return machines
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func dfs(machine Machine, x int, y int, cost int) int {
	if val, ok := cache[[2]int{x, y}]; ok {
		return val
	}

	if x == machine.goal_x && y == machine.goal_y {
		return cost
	} else if x > machine.goal_x || y > machine.goal_y {
		return math.MaxInt32
	}

	c := dfs(machine, x+machine.a_x, y+machine.a_y, cost+3)
	c2 := dfs(machine, x+machine.b_x, y+machine.b_y, cost+1)
	if c2 < c {
		c = c2
	}

	cache[[2]int{x, y}] = c
	return c
}

func cost(machine Machine) int {
	cache = map[[2]int]int{}
	c := dfs(machine, 0, 0, 0)
	if c == math.MaxInt32 {
		return 0
	} else {
		return c
	}
}

func tokens(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		total += cost(machine)
	}
	return total
}

func Run() {
	loader.Day = 13
	data := loader.GetStringGroups()
	machines := parse_data(data)

	part1 := tokens(machines)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
