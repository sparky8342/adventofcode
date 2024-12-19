package day19

import (
	"fmt"
	"loader"
	"strings"
)

func possible(pattern string, towels map[string]struct{}) bool {
	if pattern == "" {
		return true
	}
	for i := 1; i <= len(pattern); i++ {
		part := pattern[0:i]
		if _, ok := towels[part]; ok {
			if possible(pattern[i:], towels) {
				return true
			}
		}
	}
	return false
}

func possible_patterns(data []string) int {
	towels := map[string]struct{}{}
	for _, str := range strings.Split(data[0], ", ") {
		towels[str] = struct{}{}
	}

	total := 0
	for i := 2; i < len(data); i++ {
		if possible(data[i], towels) {
			total++
		}
	}

	return total
}

func Run() {
	loader.Day = 19
	data := loader.GetStrings()

	part1 := possible_patterns(data)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
