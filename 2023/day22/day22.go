package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
	z int
}

type Brick struct {
	id       int
	cubes    []Pos
	base     Pos
	top      Pos
	vertical bool
	lowest_z int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) ([]Brick, map[Pos]int) {
	bricks := []Brick{}
	positions := map[Pos]int{}

	brick_id := 0
	for _, line := range data {
		parts := strings.Split(line, "~")
		parts2 := strings.Split(parts[0], ",")
		parts3 := strings.Split(parts[1], ",")

		start_x, _ := strconv.Atoi(parts2[0])
		start_y, _ := strconv.Atoi(parts2[1])
		start_z, _ := strconv.Atoi(parts2[2])

		end_x, _ := strconv.Atoi(parts3[0])
		end_y, _ := strconv.Atoi(parts3[1])
		end_z, _ := strconv.Atoi(parts3[2])

		brick := Brick{id: brick_id, lowest_z: start_z}

		if start_x != end_x {
			for x := start_x; x <= end_x; x++ {
				pos := Pos{x: x, y: start_y, z: start_z}
				brick.cubes = append(brick.cubes, pos)
				positions[pos] = brick_id
			}
		} else if start_y != end_y {
			for y := start_y; y <= end_y; y++ {
				pos := Pos{x: start_x, y: y, z: start_z}
				brick.cubes = append(brick.cubes, pos)
				positions[pos] = brick_id
			}
		} else {
			for z := start_z; z <= end_z; z++ {
				pos := Pos{x: start_x, y: start_y, z: z}
				brick.cubes = append(brick.cubes, pos)
				positions[pos] = brick_id
			}
			brick.vertical = true
			brick.base = Pos{x: start_x, y: start_y, z: start_z}
			brick.top = Pos{x: start_x, y: start_y, z: end_z}
		}

		bricks = append(bricks, brick)
		brick_id++
	}

	return bricks, positions
}

func any_below(brick Brick, positions map[Pos]int) bool {
	if brick.lowest_z == 1 {
		return true
	}

	if brick.vertical {
		pos := Pos{x: brick.base.x, y: brick.base.y, z: brick.base.z - 1}
		_, exists := positions[pos]
		return exists
	}

	for _, cube := range brick.cubes {
		pos := Pos{x: cube.x, y: cube.y, z: cube.z - 1}
		if _, exists := positions[pos]; exists {
			return true
		}
	}

	return false
}

func move_down(bricks []Brick, positions map[Pos]int) {
	moves := true
	for moves {

		moves = false

		for i := 0; i < len(bricks); i++ {
			for !any_below(bricks[i], positions) {
				for j := range bricks[i].cubes {
					delete(positions, bricks[i].cubes[j])
					bricks[i].cubes[j].z--
					positions[bricks[i].cubes[j]] = bricks[i].id
				}
				if bricks[i].vertical {
					bricks[i].base.z--
					bricks[i].top.z--
				}
				bricks[i].lowest_z--
				moves = true
			}
		}

	}
}

func amount_below(brick Brick, positions map[Pos]int) int {
	if brick.lowest_z == 1 {
		return 0
	}

	if brick.vertical {
		pos := Pos{x: brick.base.x, y: brick.base.y, z: brick.base.z - 1}
		_, exists := positions[pos]
		if exists {
			return 1
		} else {
			return 0
		}
	}

	ids := map[int]struct{}{}
	for _, cube := range brick.cubes {
		pos := Pos{x: cube.x, y: cube.y, z: cube.z - 1}
		if id, exists := positions[pos]; exists {
			ids[id] = struct{}{}
		}
	}

	return len(ids)
}

func above(brick Brick, positions map[Pos]int) []int {
	if brick.vertical {
		pos := Pos{x: brick.top.x, y: brick.top.y, z: brick.top.z + 1}
		if id, exists := positions[pos]; exists {
			return []int{id}
		} else {
			return []int{}
		}
	}

	ids := map[int]struct{}{}
	for _, cube := range brick.cubes {
		pos := Pos{x: cube.x, y: cube.y, z: cube.z + 1}
		if id, exists := positions[pos]; exists {
			ids[id] = struct{}{}
		}
	}

	b := []int{}
	for id := range ids {
		b = append(b, id)
	}
	return b
}

func can_remove(bricks []Brick, brick_id int, positions map[Pos]int) bool {
	brick := bricks[brick_id]
	if brick.vertical {
		pos := Pos{x: brick.top.x, y: brick.top.y, z: brick.top.z + 1}
		if id, exists := positions[pos]; !exists {
			return true
		} else {
			return amount_below(bricks[id], positions) > 1
		}
	}

	bricks_above := above(brick, positions)
	for _, id := range bricks_above {
		if amount_below(bricks[id], positions) == 1 {
			return false
		}
	}

	return true
}

func free_bricks(bricks []Brick, positions map[Pos]int) int {
	free := 0
	for _, brick := range bricks {
		if can_remove(bricks, brick.id, positions) {
			free++
		}
	}
	return free
}

func main() {
	data := load_data("input.txt")
	bricks, positions := parse_data(data)
	move_down(bricks, positions)
	fmt.Println(free_bricks(bricks, positions))
}
