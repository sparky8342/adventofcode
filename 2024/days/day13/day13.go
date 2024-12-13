package day13

import (
	"fmt"
	"loader"
	"os"
	"regexp"
	"strconv"
)

type Pos struct {
	x int
	y int
}

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

func parse_data(data [][]string) [][]int {
	machines := [][]int{}
	for _, group := range data {
		machine := make([]int, 6)
		machine[0], machine[1] = parse_line(group[0])
		machine[2], machine[3] = parse_line(group[1])
		machine[4], machine[5] = parse_line(group[2])
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

func min_cost(machine []int) int {
	goal_x := machine[4]
	goal_y := machine[5]

	a := []int{machine[0], machine[1]}
	b := []int{machine[2], machine[3]}

	dp := map[Pos]int{}

	dp[Pos{x: a[0], y: a[1]}] = 3
	dp[Pos{x: b[0], y: b[1]}] = 1

	for y := 0; y <= goal_y; y++ {
		for x := 0; x <= goal_x; x++ {
			pos := Pos{x: x, y: y}
			if _, ok := dp[pos]; ok {
				next_pos := Pos{x: pos.x + a[0], y: pos.y + a[1]}
				if next_pos.x <= goal_x && next_pos.y <= goal_y {
					if _, ok := dp[next_pos]; !ok {
						dp[next_pos] = dp[pos] + 3
					} else {
						dp[next_pos] = min(dp[next_pos], dp[pos]+3)
					}
				}
				next_pos = Pos{x: pos.x + b[0], y: pos.y + b[1]}
				if next_pos.x <= goal_x && next_pos.y <= goal_y {
					if _, ok := dp[next_pos]; !ok {
						dp[next_pos] = dp[pos] + 1
					} else {
						dp[next_pos] = min(dp[next_pos], dp[pos]+1)
					}
				}

			}
		}
	}

	return dp[Pos{x: machine[4], y: machine[5]}]
}

func tokens(machines [][]int) int {
	total := 0
	for _, machine := range machines {
		total += min_cost(machine)
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
