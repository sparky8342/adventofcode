package loader

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var Day int

func get_filename() string {
	return fmt.Sprintf("inputs/day%d.txt", Day)
}

func GetStrings() []string {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func GetStringGroups() [][]string {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	group_strs := strings.Split(string(data), "\n\n")
	groups := make([][]string, len(group_strs))
	for i, group_str := range group_strs {
		groups[i] = strings.Split(group_str, "\n")
	}
	return groups
}

func GetInts() []int {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}

	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	lines := strings.Split(string(data), "\n")
	ints := make([]int, len(lines))

	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}

	return ints
}

func GetIntColumns() [][]int {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}

	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	lines := strings.Split(string(data), "\n")
	no_cols := len(strings.Fields(lines[0]))
	ints := make([][]int, no_cols)
	for i := range ints {
		ints[i] = make([]int, len(lines))
	}

	for i, line := range lines {
		for j, n_str := range strings.Fields(line) {
			n, err := strconv.Atoi(n_str)
			if err != nil {
				panic(err)
			}
			ints[j][i] = n
		}
	}

	return ints
}

func GetIntRows() [][]int {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}

	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	lines := strings.Split(string(data), "\n")

	ints := make([][]int, len(lines))
	for i, line := range lines {
		strs := strings.Fields(line)
		row := make([]int, len(strs))
		for j, n_str := range strs {
			n, err := strconv.Atoi(n_str)
			if err != nil {
				panic(err)
			}
			row[j] = n
		}
		ints[i] = row
	}

	return ints
}

func GetIntLine() []int {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}

	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	ints := []int{}
	for _, n_str := range strings.Split(string(data), " ") {
		n, err := strconv.Atoi(n_str)
		if err != nil {
			panic(err)
		}
		ints = append(ints, n)
	}

	return ints
}

func GetOneLine() []byte {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		panic(err)
	}

	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	return data
}
