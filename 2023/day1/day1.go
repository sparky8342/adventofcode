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

func main() {
	data := load_data("input.txt")
	fmt.Println(calibration(data))
}
