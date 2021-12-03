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

func find_number(numbers []string, most bool) string {
	pos := 0
	for len(numbers) > 1 {
		ones := 0
		zeroes := 0
		for _, number := range numbers {
			if number[pos] == '1' {
				ones++
			} else {
				zeroes++
			}
		}
		var match byte
		if most {
			if ones >= zeroes {
				match = '1'
			} else {
				match = '0'
			}
		} else {
			if ones < zeroes {
				match = '1'
			} else {
				match = '0'
			}
		}

		new_numbers := []string{}
		for _, number := range numbers {
			if number[pos] == match {
				new_numbers = append(new_numbers, number)
			}
		}

		numbers = new_numbers
		pos++
	}
	return numbers[0]
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

	ogr_str := find_number(numbers, true)
	ogr, _ := strconv.ParseInt(ogr_str, 2, 64)
	csr_str := find_number(numbers, false)
	csr, _ := strconv.ParseInt(csr_str, 2, 64)
	fmt.Println(ogr * csr)
}
