package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Elf struct {
	x int
	y int
}

type Empty struct {
}

type Grove struct {
	elves       map[Elf]Empty
	check_order []byte
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) Grove {
	grove := Grove{elves: map[Elf]Empty{}, check_order: []byte{'N', 'S', 'W', 'E'}}

	for y, line := range data {
		for x, ru := range line {
			if ru == '#' {
				grove.elves[Elf{x: x, y: y}] = Empty{}
			}
		}
	}

	return grove
}

func (grove *Grove) NW(elf Elf) bool {
	_, exists := grove.elves[Elf{x: elf.x - 1, y: elf.y - 1}]
	return exists
}

func (grove *Grove) N(elf Elf) bool {
	_, exists := grove.elves[Elf{x: elf.x, y: elf.y - 1}]
	return exists
}

func (grove *Grove) NE(elf Elf) bool {
	_, exists := grove.elves[Elf{x: elf.x + 1, y: elf.y - 1}]
	return exists
}

func (grove *Grove) W(elf Elf) bool {
	_, exists := grove.elves[Elf{x: elf.x - 1, y: elf.y}]
	return exists
}

func (grove *Grove) E(elf Elf) bool {
	_, exists := grove.elves[Elf{x: elf.x + 1, y: elf.y}]
	return exists
}

func (grove *Grove) SW(elf Elf) bool {
	_, exists := grove.elves[Elf{x: elf.x - 1, y: elf.y + 1}]
	return exists
}

func (grove *Grove) S(elf Elf) bool {
	_, exists := grove.elves[Elf{x: elf.x, y: elf.y + 1}]
	return exists
}

func (grove *Grove) SE(elf Elf) bool {
	_, exists := grove.elves[Elf{x: elf.x + 1, y: elf.y + 1}]
	return exists
}

func (grove *Grove) rotate_check_order() {
	first := grove.check_order[0]
	for i := 1; i < 4; i++ {
		grove.check_order[i-1] = grove.check_order[i]
	}
	grove.check_order[3] = first
}

func (grove *Grove) move_elves() bool {
	proposed_destinations := map[Elf][]Elf{}

	for elf := range grove.elves {
		if !grove.NW(elf) && !grove.N(elf) && !grove.NE(elf) && !grove.E(elf) && !grove.W(elf) && !grove.SE(elf) && !grove.S(elf) && !grove.SW(elf) {
			continue
		}

		for _, check := range grove.check_order {
			if check == 'N' && (!grove.NW(elf) && !grove.N(elf) && !grove.NE(elf)) {
				proposed_move := Elf{x: elf.x, y: elf.y - 1}
				proposed_destinations[proposed_move] = append(proposed_destinations[proposed_move], elf)
				break
			} else if check == 'S' && (!grove.SW(elf) && !grove.S(elf) && !grove.SE(elf)) {
				proposed_move := Elf{x: elf.x, y: elf.y + 1}
				proposed_destinations[proposed_move] = append(proposed_destinations[proposed_move], elf)
				break
			} else if check == 'W' && (!grove.NW(elf) && !grove.W(elf) && !grove.SW(elf)) {
				proposed_move := Elf{x: elf.x - 1, y: elf.y}
				proposed_destinations[proposed_move] = append(proposed_destinations[proposed_move], elf)
				break
			} else if check == 'E' && (!grove.NE(elf) && !grove.E(elf) && !grove.SE(elf)) {
				proposed_move := Elf{x: elf.x + 1, y: elf.y}
				proposed_destinations[proposed_move] = append(proposed_destinations[proposed_move], elf)
				break
			}
		}
	}

	moved := false
	for destination, elves := range proposed_destinations {
		if len(elves) == 1 {
			delete(grove.elves, elves[0])
			grove.elves[destination] = Empty{}
			moved = true
		}
	}
	return moved
}

func (grove *Grove) print_grove() int {
	x_min := math.MaxInt32
	x_max := math.MinInt32
	y_min := math.MaxInt32
	y_max := math.MinInt32

	for elf := range grove.elves {
		if elf.x < x_min {
			x_min = elf.x
		}
		if elf.x > x_max {
			x_max = elf.x
		}
		if elf.y < y_min {
			y_min = elf.y
		}
		if elf.y > y_max {
			y_max = elf.y
		}
	}

	empty := 0

	for x := x_min; x <= x_max; x++ {
		fmt.Print("-")
	}
	fmt.Println()
	for y := y_min; y <= y_max; y++ {
		for x := x_min; x <= x_max; x++ {
			if _, exists := grove.elves[Elf{x: x, y: y}]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
				empty++
			}
		}
		fmt.Println()
	}
	for x := x_min; x <= x_max; x++ {
		fmt.Print("-")
	}
	fmt.Println()

	return empty
}

func main() {
	data := load_data("input.txt")
	grove := parse_data(data)

	i := 0
	moved := true
	for moved {
		moved = grove.move_elves()
		grove.rotate_check_order()
		i++
		if i == 10 {
			fmt.Println(grove.print_grove())
		}
	}
	fmt.Println(i)
}
