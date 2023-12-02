package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var max_allowed Set

type Set struct {
	cube map[string]int
}

type Game struct {
	id      int
	sets    []Set
	max_set Set
	valid   bool
	power   int
}

func init() {
	max_allowed = Set{cube: map[string]int{"red": 12, "green": 13, "blue": 14}}
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_line(line string) Game {
	game := Game{
		valid:   true,
		power:   1,
		max_set: Set{cube: map[string]int{"red": 0, "green": 0, "blue": 0}},
	}

	parts := strings.Split(line, ": ")
	prefix := strings.Split(parts[0], " ")
	id, _ := strconv.Atoi(prefix[1])
	game.id = id

	sets_str := strings.Split(parts[1], "; ")

	for _, set_str := range sets_str {
		set := Set{cube: map[string]int{}}
		cube_strs := strings.Split(set_str, ", ")

		for _, cube_str := range cube_strs {
			cube_str_parts := strings.Split(cube_str, " ")
			amount_str, colour := cube_str_parts[0], cube_str_parts[1]
			amount, _ := strconv.Atoi(amount_str)
			set.cube[colour] = amount

			if amount > game.max_set.cube[colour] {
				game.max_set.cube[colour] = amount
			}
			if amount > max_allowed.cube[colour] {
				game.valid = false
			}
		}
		game.sets = append(game.sets, set)
	}

	for _, amount := range game.max_set.cube {
		game.power *= amount
	}

	return game
}

func possible(data []string) (int, int) {
	id_total := 0
	power_total := 0

	for _, line := range data {
		game := parse_line(line)
		if game.valid {
			id_total += game.id
		}
		power_total += game.power
	}

	return id_total, power_total
}

func main() {
	data := load_data("input.txt")
	possible, power_total := possible(data)
	fmt.Println(possible)
	fmt.Println(power_total)
}
