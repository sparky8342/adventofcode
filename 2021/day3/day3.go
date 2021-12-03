package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func get_data() []string {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	return lines
}

func main() {
	numbers := get_data()
	width := len(numbers[0])
	amount := len(numbers)

	ones := make([]int, width)
	for _, number := range numbers {
		for i := 0; i < width; i++ {
			if number[i] == '1' {
				ones[i]++
			}
		}
	}

	gamma_str := ""
	for i := 0; i < width; i++ {
		if ones[i] > amount/2 {
			gamma_str += "1"
		} else {
			gamma_str += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gamma_str, 2, 64)
	epsilon := gamma ^ int64(math.Pow(2, float64(width))-1)
	fmt.Println(gamma * epsilon)
}
