package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func sum2(numbers []int, target int) (int, int) {
	i := 0
	j := len(numbers) - 1
	k := numbers[i] + numbers[j]

	for k != target {
		if k > target {
			j--
		} else {
			i++
		}
		k = numbers[i] + numbers[j]
	}

	return numbers[i], numbers[j]
}

func sum3(numbers []int, target int) (int, int, int) {
	for _, n := range numbers {
		num1, num2 := sum2(numbers, target-n)
		if num1 != n && num2 != n {
			return n, num1, num2
		}
	}
	return 0, 0, 0
}

func get_data() []int {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	numbers := []int{}
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}
	sort.Ints(numbers)
	return numbers
}

func main() {
	numbers := get_data()
	target := 2020

	num1, num2 := sum2(numbers, target)
	fmt.Println(num1 * num2)

	num1, num2, num3 := sum3(numbers, target)
	fmt.Println(num1 * num2 * num3)
}
