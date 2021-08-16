package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Dir struct {
	dx int
	dy int
}

type Pos struct {
	x int
	y int
}

type Grid struct {
	squares     [][]byte
	robots      [4]Pos
	robot_count int
	player      Pos
	goal        int
}

type State struct {
	robots      [4]Pos
	robot_count int
	keys        [26]bool
}

type QueueEntry struct {
	state    State
	distance int
}

var dirs [4]Dir

func init() {
	dirs = [4]Dir{
		Dir{dx: 1, dy: 0},
		Dir{dx: -1, dy: 0},
		Dir{dx: 0, dy: 1},
		Dir{dx: 0, dy: -1},
	}
}

func load_grid() Grid {
	squares := [][]byte{{}}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		squares = append(squares, []byte(line))
	}
	grid := Grid{squares: squares}
	return grid
}

func find_elements(grid Grid) Grid {
	robots := [4]Pos{}
	robot_count := 0
	goal := 0

	for y, row := range grid.squares {
		for x, space := range row {
			if space == '@' {
				robots[robot_count].x = x
				robots[robot_count].y = y
				robot_count++
			} else if space >= 'a' && space <= 'z' && int(space-'a') > goal {
				goal = int(space - 'a')
			}

		}
	}

	grid.robots = robots
	grid.robot_count = robot_count
	grid.goal = goal
	return grid
}

func bfs(grid Grid) int {
	start := State{robots: grid.robots, robot_count: grid.robot_count}
	entry := QueueEntry{state: start, distance: 0}
	queue := []QueueEntry{entry}

	visited := make(map[State]bool)

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		state := entry.state

		if visited[state] {
			continue
		}
		visited[state] = true

		for i := 0; i < state.robot_count; i++ {
			robot := state.robots[i]
			keys := state.keys

			space := grid.squares[robot.y][robot.x]

			if space == byte('#') {
				continue
			}

			if space >= byte('A') && space <= byte('Z') && keys[space-65] == false {
				// at door without key
				continue
			}

			if space >= byte('a') && space <= byte('z') {
				// found a key
				keys[space-97] = true

				all_found := true
				for i := 0; i <= grid.goal; i++ {
					if keys[i] == false {
						all_found = false
						break
					}
				}
				if all_found {
					return entry.distance
				}
			}

			for _, dir := range dirs {
				robots_copy := state.robots
				robots_copy[i].x += dir.dx
				robots_copy[i].y += dir.dy
				new_state := State{
					robots:      robots_copy,
					robot_count: state.robot_count,
					keys:        keys,
				}
				queue = append(queue, QueueEntry{state: new_state, distance: entry.distance + 1})
			}
		}
	}

	return 0
}

func main() {
	grid := load_grid()
	grid = find_elements(grid)
	distance := bfs(grid)
	fmt.Println(distance)
}
