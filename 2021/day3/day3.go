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

func count_at_pos(numbers []string, pos int) (int, int) {
	ones := 0
	for _, number := range numbers {
		if number[pos] == '1' {
			ones++
		}
	}
	return len(numbers) - ones, ones
}

func get_oxygen_generator_rating(numbers []string) int64 {
	return find_number(numbers, true)
}

func get_co2_scrubber_rating(numbers []string) int64 {
	return find_number(numbers, false)
}

func find_number(numbers []string, most bool) int64 {
	pos := 0
	for len(numbers) > 1 {
		zeroes, ones := count_at_pos(numbers, pos)

		var match byte

		if (most && ones >= zeroes) || (!most && ones < zeroes) {
			match = '1'
		} else {
			match = '0'
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
	num, _ := strconv.ParseInt(numbers[0], 2, 64)
	return num
}

func main() {
	numbers := get_data()

	// part 1
	width := len(numbers[0])
	gamma_str := ""
	for i := 0; i < width; i++ {
		zeroes, ones := count_at_pos(numbers, i)
		if ones > zeroes {
			gamma_str += "1"
		} else {
			gamma_str += "0"
		}
	}
	gamma, _ := strconv.ParseInt(gamma_str, 2, 64)
	epsilon := gamma ^ int64(math.Pow(2, float64(width))-1)
	fmt.Println(gamma * epsilon)

	// part 2
	ogr := get_oxygen_generator_rating(numbers)
	csr := get_co2_scrubber_rating(numbers)
	fmt.Println(ogr * csr)
}
