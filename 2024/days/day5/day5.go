package day5

import (
	"fmt"
	"loader"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func valid_updates(data [][]string) (int, int) {
	rules := map[[2]int]struct{}{}
	for _, line := range data[0] {
		parts := strings.Split(line, "|")
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		rules[[2]int{n1, n2}] = struct{}{}
	}

	valid_sum := 0
	corrected_sum := 0

	for _, line := range data[1] {
		row := []int{}
		for _, str := range strings.Split(line, ",") {
			n, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}

		cpy := make([]int, len(row))
		copy(cpy, row)
		sort.Slice(cpy, func(i, j int) bool {
			_, ok := rules[[2]int{cpy[i], cpy[j]}]
			return ok
		})

		if reflect.DeepEqual(cpy, row) {
			valid_sum += row[len(row)/2]
		} else {
			corrected_sum += cpy[len(cpy)/2]
		}
	}

	return valid_sum, corrected_sum
}

func Run() {
	loader.Day = 5
	data := loader.GetStringGroups()
	part1, part2 := valid_updates(data)

	fmt.Printf("%d %d\n", part1, part2)
}
