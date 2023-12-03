package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func is_digit(b byte) bool {
	return b >= '0' && b <= '9'
}

func adjacent_to_symbol(data []string, x int, x2 int, y int) bool {
	height := len(data)
	width := len(data[0])

	for dx := x - 1; dx <= x2+1; dx++ {
		if dx < 0 || dx == width {
			continue
		}
		if y > 0 && data[y-1][dx] != '.' {
			return true
		}
		if y < height-1 && data[y+1][dx] != '.' {
			return true
		}
	}
	if (x > 0 && data[y][x-1] != '.') || (x2 < width-1 && data[y][x2+1] != '.') {
		return true
	}

	return false
}

func part_sum(data []string) int {
	height := len(data)
	width := len(data[0])

	sum := 0

	for y := 0; y < height; y++ {
		x := 0

		for x < width {
			for x < width && !is_digit(data[y][x]) {
				x++
			}
			x2 := x
			for x2 < width && is_digit(data[y][x2]) {
				x2++
			}
			x2--

			if x == width {
				continue
			}

			if adjacent_to_symbol(data, x, x2, y) {
				num := int(data[y][x] - '0')
				for i := x + 1; i <= x2; i++ {
					num = num*10 + int(data[y][i]-'0')
				}
				sum += num
			}

			x = x2 + 1
		}
	}

	return sum
}

func main() {
	data := load_data("input.txt")
	fmt.Println(part_sum(data))

}
