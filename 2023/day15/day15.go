package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type BoxLine struct {
	boxes [256][]Lens
}

type Lens struct {
	label        string
	focal_length int
}

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func hash(str string) int {
	h := 0
	for _, ru := range str {
		h = ((h + int(ru)) * 17) % 256
	}
	return h
}

func hash_sum(data string) int {
	sum := 0
	for _, str := range strings.Split(data, ",") {
		sum += hash(str)
	}
	return sum
}

func (box_line *BoxLine) set(label string, focal_length int) {
	box_id := hash(label)

	found := false
	for j, lens := range box_line.boxes[box_id] {
		if lens.label == label {
			box_line.boxes[box_id][j].focal_length = focal_length
			found = true
			return
		}
	}
	if !found {
		box_line.boxes[box_id] = append(box_line.boxes[box_id], Lens{label: label, focal_length: focal_length})
	}
}

func (box_line *BoxLine) del(label string) {
	box_id := hash(label)

	for i, lens := range box_line.boxes[box_id] {
		if lens.label == label {
			box_line.boxes[box_id] = append(box_line.boxes[box_id][:i], box_line.boxes[box_id][i+1:]...)
			return
		}
	}
}

func (box_line *BoxLine) power() int {
	power := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < len(box_line.boxes[i]); j++ {
			power += (i + 1) * (j + 1) * box_line.boxes[i][j].focal_length
		}
	}
	return power
}

func operations(data string) int {
	box_line := BoxLine{boxes: [256][]Lens{}}

	for _, str := range strings.Split(data, ",") {
		for i, ru := range str {
			if ru == '=' {
				label := str[0:i]
				focal_length := int(str[i+1] - '0')
				box_line.set(label, focal_length)
				break
			} else if ru == '-' {
				label := str[0:i]
				box_line.del(label)
			}
		}
	}

	return box_line.power()
}

func main() {
	data := load_data("input.txt")
	fmt.Println(hash_sum(data))
	fmt.Println(operations(data))
}
