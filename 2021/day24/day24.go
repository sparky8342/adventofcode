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

func dedupe(nums []int) []int {
	mp := map[int]struct{}{}
	for _, num := range nums {
		mp[num] = struct{}{}
	}
	ret := []int{}
	for num, _ := range mp {
		ret = append(ret, num)
	}
	return ret
}

func main() {
	get_data()

	z_map := map[int][]int{0: []int{0}}

	//max_z_map := map[int]int{0: 0}
	//min_z_map := map[int]int{0: 0}

	// run the sections in reverse, keeping tracking of
	// which inputs lead to 0 at the last step
	for section := 13; section >= 0; section-- {
		new_z_map := map[int][]int{}
		for z := 0; z <= 10000000; z++ {
			for digit := 1; digit <= 9; digit++ {
				new_z := run_section(section, z, digit)

				if vals, found := z_map[new_z]; found {
					new_vals := []int{}
					for _, val := range vals {
						new_val := val + (digit * pow10(13-section))
						new_vals = append(new_vals, new_val)
					}

					if vals2, found2 := new_z_map[z]; found2 {
						new_vals = dedupe(append(new_vals, vals2...))
					}
					new_z_map[z] = new_vals
				}
			}
		}
		z_map = new_z_map
	}

	ids := z_map[0]
	min := ids[0]
	max := ids[0]
	for i := 1; i < len(ids); i++ {
		if ids[i] < min {
			min = ids[i]
		}
		if ids[i] > max {
			max = ids[i]
		}
	}
	fmt.Println(max)
	fmt.Println(min)
}
