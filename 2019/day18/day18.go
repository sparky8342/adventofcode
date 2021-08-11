package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	x int
	y int
}

type Grid struct {
	squares [][]byte
	player  Pos
	goal    int
}

type State struct {
	x    int
	y    int
	keys [26]bool
}

type QueueEntry struct {
	state    State
	distance int
}

func get_grid() Grid {
	squares := [][]byte{{}}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		squares = append(squares, []byte(line))
	}
	file.Close()

	var player Pos
	goal := 0

	for y, row := range squares {
		for x, space := range row {
			if space == '@' {
				player.x = x
				player.y = y
			} else if space >= 'a' && space <= 'z' && int(space-'a') > goal {
				goal = int(space - 'a')
			}

		}
	}

	grid := Grid{squares: squares, player: player, goal: goal}
	return grid
}

func bfs(grid Grid) int {
	start := State{x: grid.player.x, y: grid.player.y}
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

		if grid.squares[state.y][state.x] == byte('#') {
			continue
		}

		keys := state.keys
		space := grid.squares[state.y][state.x]

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

		queue = append(queue, QueueEntry{state: State{x: state.x + 1, y: state.y, keys: keys}, distance: entry.distance + 1})
		queue = append(queue, QueueEntry{state: State{x: state.x - 1, y: state.y, keys: keys}, distance: entry.distance + 1})
		queue = append(queue, QueueEntry{state: State{x: state.x, y: state.y + 1, keys: keys}, distance: entry.distance + 1})
		queue = append(queue, QueueEntry{state: State{x: state.x, y: state.y - 1, keys: keys}, distance: entry.distance + 1})
	}

	return 0
}

func main() {
	grid := get_grid()
	distance := bfs(grid)
	fmt.Println(distance)
}
