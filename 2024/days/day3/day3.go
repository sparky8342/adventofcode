package day3

import (
	"fmt"
	"loader"
	"os"
	"regexp"
	"strconv"
)

func find_valid(data []string, use_conditionals bool) int {
	var r *regexp.Regexp
	if use_conditionals {
		r = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")
	} else {
		r = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	}

	total := 0
	on := true

	for _, line := range data {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				on = true
			} else if match[0] == "don't()" {
				on = false
			} else if on {
				n1, err := strconv.Atoi(match[1])
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
					os.Exit(1)
				}
				n2, err := strconv.Atoi(match[2])
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
					os.Exit(1)
				}
				total += n1 * n2
			}
		}
	}

	return total
}

func Run() {
	loader.Day = 3
	data := loader.GetStrings()

	part1 := find_valid(data, false)
	part2 := find_valid(data, true)

	fmt.Printf("%d %d\n", part1, part2)
}
