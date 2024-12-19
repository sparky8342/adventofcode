package day19

import (
	"fmt"
	"loader"
	"strings"
)

var cache map[string]int

func init() {
	cache = map[string]int{}
}

func possible(pattern string, towels map[string]struct{}) int {
	if pattern == "" {
		return 1
	}
	if amount, ok := cache[pattern]; ok {
		return amount
	}
	ways := 0
	for i := 1; i <= len(pattern); i++ {
		part := pattern[0:i]
		if _, ok := towels[part]; ok {
			ways += possible(pattern[i:], towels)
		}
	}
	cache[pattern] = ways
	return ways
}

func possible_patterns(data []string) (int, int) {
	towels := map[string]struct{}{}
	for _, str := range strings.Split(data[0], ", ") {
		towels[str] = struct{}{}
	}

	total := 0
	ways := 0
	for i := 2; i < len(data); i++ {
		w := possible(data[i], towels)
		if w > 0 {
			total++
			ways += w
		}
	}

	return total, ways
}

func Run() {
	loader.Day = 19
	data := loader.GetStrings()

	part1, part2 := possible_patterns(data)

	fmt.Printf("%d %d\n", part1, part2)
}
