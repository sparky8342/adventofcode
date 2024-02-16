package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Update struct {
	value bool
	left  bool
	state byte
}

type Card struct {
	zero Update
	one  Update
}

type Cell struct {
	value bool
	left  *Cell
	right *Cell
}

func get_data(filename string) (byte, int, map[byte]Card) {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	sections := strings.Split(string(data), "\n\n")

	prefix := strings.Split(sections[0], "\n")
	start_state := prefix[0][15]
	steps, _ := strconv.Atoi(strings.Split(prefix[1], " ")[5])

	sections = sections[1:]

	cards := map[byte]Card{}

	for _, section := range sections {
		lines := strings.Split(section, "\n")
		state := lines[0][9]

		value := false
		if lines[2][22] == '1' {
			value = true
		}
		left := true
		if lines[3][27] == 'r' {
			left = false
		}
		st := lines[4][26]
		zero_update := Update{
			value: value,
			left:  left,
			state: st,
		}

		value = false
		if lines[6][22] == '1' {
			value = true
		}
		left = true
		if lines[7][27] == 'r' {
			left = false
		}
		st = lines[8][26]
		one_update := Update{
			value: value,
			left:  left,
			state: st,
		}

		card := Card{
			zero: zero_update,
			one:  one_update,
		}
		cards[state] = card
	}

	return start_state, steps, cards
}

func main() {
	state, steps, cards := get_data("input.txt")

	cell := &Cell{}

	for i := 0; i < steps; i++ {
		card := cards[state]

		var update Update
		if cell.value {
			update = card.one
		} else {
			update = card.zero
		}

		cell.value = update.value
		state = update.state

		if update.left {
			if cell.left == nil {
				cell.left = &Cell{}
				cell.left.right = cell
			}
			cell = cell.left
		} else {
			if cell.right == nil {
				cell.right = &Cell{}
				cell.right.left = cell
			}
			cell = cell.right
		}
	}

	for cell.left != nil {
		cell = cell.left
	}

	count := 0
	if cell.value {
		count++
	}
	for cell.right != nil {
		cell = cell.right
		if cell.value {
			count++
		}
	}

	fmt.Println(count)
}
