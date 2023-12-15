package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func operations(data string) int {
	boxes := make([][]Lens, 256)

	for _, str := range strings.Split(data, ",") {
		for i, ru := range str {
			if ru == '=' {
				label := str[0:i]
				focal_length := int(str[i+1] - '0')
				box_id := hash(label)

				found := false
				for j, lens := range boxes[box_id] {
					if lens.label == label {
						boxes[box_id][j].focal_length = focal_length
						found = true
						break
					}
				}
				if !found {
					boxes[box_id] = append(boxes[box_id], Lens{label: label, focal_length: focal_length})
				}

			} else if ru == '-' {
				label := str[0:i]
				box_id := hash(label)

				for j, lens := range boxes[box_id] {
					if lens.label == label {
						boxes[box_id] = append(boxes[box_id][:j], boxes[box_id][j+1:]...)
						break
					}
				}
			}
		}
	}

	power := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < len(boxes[i]); j++ {
			power += (i + 1) * (j + 1) * boxes[i][j].focal_length
		}
	}

	return power
}

func main() {
	data := load_data("input.txt")
	fmt.Println(hash_sum(data))
	fmt.Println(operations(data))
}
