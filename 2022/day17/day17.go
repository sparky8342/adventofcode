package main

import (
	"fmt"
	"io/ioutil"
)

type Pos struct {
	x int
	y int
}

type Empty struct {
}

type Chamber struct {
	jets     string
	jet_pos  int
	spaces   map[Pos]Empty
	top_y    int
	piece_id int
}

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func NewChamber(jets string) *Chamber {
	c := new(Chamber)
	c.jets = jets
	c.spaces = map[Pos]Empty{}
	c.top_y = -1
	//for x := 0; x < 7; x++ {
	//	c.spaces[Pos{x: x, y: 0}] = Empty{}
	//}
	return c
}

func (chamber *Chamber) drop_piece() {
	y := chamber.top_y + 4
	var parts []*Pos

	switch chamber.piece_id {
	case 0:
		parts = []*Pos{&Pos{2, y}, &Pos{3, y}, &Pos{4, y}, &Pos{5, y}}
	case 1:
		parts = []*Pos{&Pos{3, y}, &Pos{2, y + 1}, &Pos{3, y + 1}, &Pos{4, y + 1}, &Pos{3, y + 2}}
	case 2:
		parts = []*Pos{&Pos{2, y}, &Pos{3, y}, &Pos{4, y}, &Pos{4, y + 1}, &Pos{4, y + 2}}
	case 3:
		parts = []*Pos{&Pos{2, y}, &Pos{2, y + 1}, &Pos{2, y + 2}, &Pos{2, y + 3}}
	case 4:
		parts = []*Pos{&Pos{2, y}, &Pos{3, y}, &Pos{2, y + 1}, &Pos{3, y + 1}}
	}

outer:
	for {

		//fmt.Println(parts[0])

		// jet move
		var dx int
		if chamber.jets[chamber.jet_pos] == '<' {
			dx = -1
		} else if chamber.jets[chamber.jet_pos] == '>' {
			dx = 1
		}
		chamber.jet_pos++
		if chamber.jet_pos == len(chamber.jets) {
			chamber.jet_pos = 0
		}

		can_move := true
		for _, part := range parts {
			if part.x+dx < 0 || part.x+dx > 6 {
				can_move = false
				break
			}
			if _, rock := chamber.spaces[Pos{x: part.x + dx, y: part.y}]; rock {
				can_move = false
				break
			}
		}
		if can_move {
			for _, part := range parts {
				part.x += dx
			}
		}

		//fmt.Println(parts[0])

		// move down
		for _, part := range parts {
			if part.y == 0 {
				break outer
			}
			if _, rock := chamber.spaces[Pos{x: part.x, y: part.y - 1}]; rock {
				break outer
			}
		}
		for _, part := range parts {
			part.y--
		}
	}

	max_y := 0
	for _, part := range parts {
		chamber.spaces[*part] = Empty{}
		if part.y > max_y {
			max_y = part.y
		}
	}
	if max_y > chamber.top_y {
		chamber.top_y = max_y
	}

	chamber.piece_id = (chamber.piece_id + 1) % 5
}

func (chamber *Chamber) draw() {
	for y := chamber.top_y; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if _, rock := chamber.spaces[Pos{x: x, y: y}]; rock {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	data := load_data("input.txt")

	chamber := NewChamber(data)

	for i := 0; i < 2022; i++ {
		chamber.drop_piece()
	}

	fmt.Println(chamber.top_y + 1)
}
