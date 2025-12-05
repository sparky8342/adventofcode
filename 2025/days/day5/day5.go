package day5

import (
	"fmt"
	"loader"
	"strconv"
	"strings"
)

type pair [2]int

func parse_data(data []string) ([]pair, []int) {
	ranges := []pair{}

	var i int
	for i = 0; i < len(data); i++ {
		if data[i] == "" {
			break
		}
		parts := strings.Split(data[i], "-")
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, pair{n1, n2})
	}

	ingredients := []int{}

	for i = i + 1; i < len(data); i++ {
		n, err := strconv.Atoi(data[i])
		if err != nil {
			panic(err)
		}
		ingredients = append(ingredients, n)
	}

	return ranges, ingredients
}

func in_ranges(ranges []pair, n int) bool {
	for _, r := range ranges {
		if n >= r[0] && n <= r[1] {
			return true
		}
	}
	return false
}

func count_fresh(ranges []pair, ingredients []int) int {
	fresh := 0
	for _, n := range ingredients {
		if in_ranges(ranges, n) {
			fresh++
		}
	}
	return fresh
}

func Run() {
	loader.Day = 5
	data := loader.GetStrings()
	ranges, ingredients := parse_data(data)
	part1 := count_fresh(ranges, ingredients)

	fmt.Printf("%d %d\n", part1, 0)
}
