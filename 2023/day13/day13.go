package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func get_notes(data string) (int, int) {
	patterns := strings.Split(data, "\n\n")

	total := 0
	total_second := 0
	for _, pattern_str := range patterns {
		lines := strings.Split(pattern_str, "\n")
		pattern := [][]byte{}
		for _, line := range lines {
			pattern = append(pattern, []byte(line))
		}
		amount := process_pattern(pattern, -1)
		total += amount
		second := find_second_mirror(pattern, amount)
		total_second += second
	}
	return total, total_second
}

func flip(b byte) byte {
	if b == '#' {
		return '.'
	} else {
		return '#'
	}
}

func find_second_mirror(pattern [][]byte, previous_amount int) int {
	height := len(pattern)
	width := len(pattern[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pattern[y][x] = flip(pattern[y][x])
			amount := process_pattern(pattern, previous_amount)
			if amount != -1 {
				return amount
			}
			pattern[y][x] = flip(pattern[y][x])
		}
	}

	return previous_amount
}

func is_horizontal_mirror(pattern [][]byte, x int, x2 int) bool {
	x--
	x2++
	for x >= 0 && x2 < len(pattern) {
		if string(pattern[x]) != string(pattern[x2]) {
			return false
		}
		x--
		x2++
	}
	return true
}

func is_vertical_mirror(pattern [][]byte, y int, y2 int) bool {
	y--
	y2++
	for y >= 0 && y2 < len(pattern[0]) {
		col := []byte{}
		col2 := []byte{}
		for j := 0; j < len(pattern); j++ {
			col = append(col, pattern[j][y])
			col2 = append(col2, pattern[j][y2])
		}
		if string(col) != string(col2) {
			return false
		}
		y--
		y2++
	}
	return true
}

func process_pattern(pattern [][]byte, ignore int) int {
	height := len(pattern)
	width := len(pattern[0])

	// rows
	for i := 0; i < height-1; i++ {
		if string(pattern[i]) == string(pattern[i+1]) {
			if is_horizontal_mirror(pattern, i, i+1) {
				note := (i + 1) * 100
				if note != ignore {
					return note
				}
			}
		}
	}

	// columns
	for i := 0; i < width-1; i++ {
		col := []byte{}
		col2 := []byte{}
		for j := 0; j < height; j++ {
			col = append(col, pattern[j][i])
			col2 = append(col2, pattern[j][i+1])
		}
		if string(col) == string(col2) {
			if is_vertical_mirror(pattern, i, i+1) {
				note := i + 1
				if note != ignore {
					return note
				}
			}
		}
	}

	return -1
}

func pp(pattern [][]byte) {
	for _, row := range pattern {
		fmt.Println(string(row))
	}
}

func main() {
	data := load_data("input.txt")
	notes, notes2 := get_notes(data)
	fmt.Println(notes)
	fmt.Println(notes2)
}
