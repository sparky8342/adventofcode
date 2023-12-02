package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func possible(data []string) (int, int) {
	id_total := 0
	power_total := 0

	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, line := range data {
		parts := strings.Split(line, ": ")
		prefix := strings.Split(parts[0], " ")
		id, _ := strconv.Atoi(prefix[1])

		sets := strings.Split(parts[1], "; ")

		valid := true
		max_found := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, set := range sets {
			cube_strs := strings.Split(set, ", ")
			cubes := map[string]int{}
			for _, cube_str := range cube_strs {
				cube_str_parts := strings.Split(cube_str, " ")
				n, _ := strconv.Atoi(cube_str_parts[0])
				cubes[cube_str_parts[1]] = n
			}

			for key, value := range cubes {
				if value > max[key] {
					valid = false
				}
				if value > max_found[key] {
					max_found[key] = value
				}
			}
		}

		if valid {
			id_total += id
		}

		power := 1
		for _, value := range max_found {
			power *= value
		}
		power_total += power

	}

	return id_total, power_total
}

func main() {
	data := load_data("input.txt")
	possible, power_total := possible(data)
	fmt.Println(possible)
	fmt.Println(power_total)
}
