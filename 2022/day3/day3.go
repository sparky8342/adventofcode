package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func letter_to_bit(letter byte) int {
	if letter >= 'a' {
		return int(letter) - 'a'
	} else {
		return int(letter) - 'A' + 26
	}
}

func priority_sum(data []string) int {
	sum := 0

	for _, line := range data {
		var num int64 = 0

		for i := 0; i < len(line)/2; i++ {
			bit := letter_to_bit(line[i])
			num = num | (1 << bit)
		}

		for i := len(line) / 2; i < len(line); i++ {
			bit := letter_to_bit(line[i])
			if num&(1<<bit) != 0 {
				sum += bit + 1
				break
			}
		}
	}
	return sum
}

func badge_sum(data []string) int {
	sum := 0

	for i := 0; i < len(data); i += 3 {
		nums := [3]int64{}
		for j := 0; j < 3; j++ {
			var num int64 = 0
			for _, ru := range data[i+j] {
				bit := letter_to_bit(byte(ru))
				num = num | (1 << bit)
			}
			nums[j] = num
		}
		intersection := nums[0] & nums[1] & nums[2]
		for i := 0; i < 52; i++ {
			if intersection&(1<<i) != 0 {
				sum += i + 1
				break
			}
		}
	}
	return sum
}

func main() {
	data := load_data("input.txt")
	sum := priority_sum(data)
	fmt.Println(sum)
	sum = badge_sum(data)
	fmt.Println(sum)
}
