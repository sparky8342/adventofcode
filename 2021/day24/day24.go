package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var div_z [14]int
var add_x [14]int
var add_y [14]int

func get_data() {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")

	for i := 0; i < 14; i++ {
		l := lines[i*18+4]
		parts := strings.Split(l, " ")
		div_z[i], _ = strconv.Atoi(parts[2])

		l = lines[i*18+5]
		parts = strings.Split(l, " ")
		add_x[i], _ = strconv.Atoi(parts[2])

		l = lines[i*18+15]
		parts = strings.Split(l, " ")
		add_y[i], _ = strconv.Atoi(parts[2])
	}
}

func run_section(section int, z int, digit int) int {
	x := z % 26
	z = z / div_z[section]
	x += add_x[section]
	if x == digit {
		x = 0
	} else {
		x = 1
	}
	y := 25*x + 1
	z = z * y
	y = digit + add_y[section]
	y = y * x
	z = z + y
	return z
}

func pow10(p int) int {
	r := 1
	for i := 0; i < p; i++ {
		r = r * 10
	}
	return r
}

func main() {
	get_data()

	max_z_map := map[int]int{0: 0}
	min_z_map := map[int]int{0: 0}

	// run the sections in reverse, keeping tracking of
	// which inputs lead to 0 at the last step
	for section := 13; section >= 0; section-- {
		new_max_z_map := map[int]int{}
		new_min_z_map := map[int]int{}
		for z := 0; z <= 10000000; z++ {
			for digit := 1; digit <= 9; digit++ {
				new_z := run_section(section, z, digit)

				if val, found := max_z_map[new_z]; found {
					new_val := val + (digit * pow10(13-section))
					if val2, found2 := new_max_z_map[z]; found2 {
						if new_val > val2 {
							new_max_z_map[z] = new_val
						}
					} else {
						new_max_z_map[z] = new_val
					}
				}

				if val, found := min_z_map[new_z]; found {
					new_val := val + (digit * pow10(13-section))
					if val2, found2 := new_min_z_map[z]; found2 {
						if new_val < val2 {
							new_min_z_map[z] = new_val
						}
					} else {
						new_min_z_map[z] = new_val
					}
				}

			}
		}
		max_z_map = new_max_z_map
		min_z_map = new_min_z_map
	}

	fmt.Println(max_z_map[0])
	fmt.Println(min_z_map[0])
}
