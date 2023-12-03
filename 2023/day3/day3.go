package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pos struct {
	x, y int
}

type Num struct {
	value    int
	star_pos Pos
}

var width, height int

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

func adjacent_to_symbol(data []string, x int, x2 int, y int) (bool, int, int) {
	for dx := x - 1; dx <= x2+1; dx++ {
		if dx < 0 || dx == width {
			continue
		}
		if y > 0 && data[y-1][dx] != '.' {
			return true, dx, y - 1
		}
		if y < height-1 && data[y+1][dx] != '.' {
			return true, dx, y + 1
		}
	}
	if x > 0 && data[y][x-1] != '.' {
		return true, x - 1, y
	}

	if x2 < width-1 && data[y][x2+1] != '.' {
		return true, x2 + 1, y
	}

	return false, -1, -1
}

func sum_and_ratio(data []string) (int, int) {
	height = len(data)
	width = len(data[0])

	sum := 0

	nums := []Num{}

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

			is_adj, adj_x, adj_y := adjacent_to_symbol(data, x, x2, y)
			if is_adj {
				num := int(data[y][x] - '0')
				for i := x + 1; i <= x2; i++ {
					num = num*10 + int(data[y][i]-'0')
				}
				sum += num

				if data[adj_y][adj_x] == '*' {
					nums = append(nums, Num{
						value:    num,
						star_pos: Pos{x: adj_x, y: adj_y},
					})
				}
			}

			x = x2 + 1
		}
	}

	ratio := 0

	seen := map[Pos]int{}
	for _, n := range nums {
		if val, ok := seen[n.star_pos]; ok {
			ratio += val * n.value
		} else {
			seen[n.star_pos] = n.value
		}
	}

	return sum, ratio
}

func main() {
	data := load_data("input.txt")
	sum, ratio := sum_and_ratio(data)
	fmt.Println(sum)
	fmt.Println(ratio)
}
