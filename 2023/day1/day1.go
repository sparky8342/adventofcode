package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func calibration(data []string) int {
	var num, total int

	for _, line := range data {
		for i := 0; i < len(line); i++ {
			if line[i] >= '1' && line[i] <= '9' {
				num = 10 * int(line[i]-'0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '1' && line[i] <= '9' {
				num += int(line[i] - '0')
				break
			}
		}
		total += num
	}

	return total
}

func calibration2(data []string) int {
	re := regexp.MustCompile("(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine)")
	re2 := regexp.MustCompile(".*(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine)")

	words := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var total, num int

	for _, line := range data {
		match := re.FindString(line)

		if len(match) == 1 {
			num = 10 * int(match[0]-'0')
		} else {
			num = 10 * words[match]
		}

		m := re2.FindStringSubmatch(line)
		match = m[1]

		if len(match) == 1 {
			num += int(match[0] - '0')
		} else {
			num += words[match]
		}

		total += num
	}

	return total
}

func main() {
	data := load_data("input.txt")
	fmt.Println(calibration(data))
	fmt.Println(calibration2(data))
}
