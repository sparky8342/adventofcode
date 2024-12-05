package day5

import (
	"fmt"
	"loader"
	"os"
	"strconv"
	"strings"
)

func valid_updates(data [][]string) int {
	rules := map[[2]int]struct{}{}
	for _, line := range data[0] {
		parts := strings.Split(line, "|")
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		rules[[2]int{n1, n2}] = struct{}{}
	}

	sum := 0

outer:
	for _, line := range data[1] {
		row := []int{}
		for _, str := range strings.Split(line, ",") {
			n, err := strconv.Atoi(str)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			row = append(row, n)
		}

		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				if _, ok := rules[[2]int{row[i], row[j]}]; !ok {
					continue outer
				}
			}
		}

		sum += row[len(row)/2]
	}

	return sum
}

func Run() {
	loader.Day = 5
	data := loader.GetStringGroups()
	part1 := valid_updates(data)

	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
