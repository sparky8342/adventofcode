package day5

import (
	"fmt"
	"loader"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func parse_data(data []string) ([]Range, []int) {
	var i int

	ranges := []Range{}
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
		ranges = append(ranges, Range{start: n1, end: n2})
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

func in_ranges(ranges []Range, n int) bool {
	for _, r := range ranges {
		if n >= r.start && n <= r.end {
			return true
		}
	}
	return false
}

func count_fresh(ranges []Range, ingredients []int) int {
	fresh := 0
	for _, n := range ingredients {
		if in_ranges(ranges, n) {
			fresh++
		}
	}
	return fresh
}

func fresh_ids(ranges []Range) int {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start == ranges[j].start {
			return ranges[i].end < ranges[j].end
		} else {
			return ranges[i].start < ranges[j].start
		}
	})

	fresh := 0

	current := ranges[0]
	for i := 1; i < len(ranges); i++ {
		this_range := ranges[i]
		if current.end >= this_range.start {
			if current.end >= this_range.end {
				continue
			} else {
				current.end = this_range.end
			}
		} else {
			fresh += current.end - current.start + 1
			current = this_range
		}
	}
	fresh += current.end - current.start + 1

	return fresh
}

func Run() {
	loader.Day = 5
	data := loader.GetStrings()
	ranges, ingredients := parse_data(data)
	part1 := count_fresh(ranges, ingredients)
	part2 := fresh_ids(ranges)

	fmt.Printf("%d %d\n", part1, part2)
}
