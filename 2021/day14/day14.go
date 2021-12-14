package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Rules map[[2]byte]byte

type CacheKey struct {
	code  [2]byte
	steps int
}

var cache map[CacheKey]map[byte]int

func init() {
	cache = make(map[CacheKey]map[byte]int)
}

func get_data() (string, Rules) {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	sections := strings.Split(string(data), "\n\n")

	polymer := sections[0]

	rules := Rules{}
	for _, line := range strings.Split(sections[1], "\n") {
		rules[[2]byte{line[0], line[1]}] = line[6]
	}

	return polymer, rules
}

func calculate(rules Rules, polymer string, steps int) int {
	result := map[byte]int{}
	for i := 0; i < len(polymer)-1; i++ {
		r := get_amount(rules, [2]byte{polymer[i], polymer[i+1]}, steps)
		for k, v := range r {
			result[k] += v
		}
	}
	result[polymer[0]]++

	max := 0
	min := -1
	for _, count := range result {
		if count > max {
			max = count
		} else if min == -1 || count < min {
			min = count
		}
	}
	return max - min
}

func get_amount(rules Rules, code [2]byte, steps int) map[byte]int {
	key := CacheKey{code: code, steps: steps}
	if val, found := cache[key]; found {
		return val
	}

	new_byte := rules[code]

	if steps == 1 {
		result := map[byte]int{}
		result[code[1]]++
		result[new_byte]++
		return result
	}

	result1 := get_amount(rules, [2]byte{code[0], new_byte}, steps-1)
	result2 := get_amount(rules, [2]byte{new_byte, code[1]}, steps-1)

	result := map[byte]int{}
	for k, v := range result1 {
		result[k] += v
	}
	for k, v := range result2 {
		result[k] += v
	}

	cache[key] = result
	return result
}

func main() {
	polymer, rules := get_data()

	fmt.Println(calculate(rules, polymer, 10))
	fmt.Println(calculate(rules, polymer, 40))
}
