package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func pow(a int, b int) int {
	if b == 0 {
		return 1
	}
	p := a
	for i := 1; i < b; i++ {
		p = p * a
	}
	return p
}

func snafu_to_dec(snafu string) int {
	dec := 0
	p := 0

	for i := len(snafu) - 1; i >= 0; i-- {
		var digit int
		switch snafu[i] {
		case '=':
			digit = -2
		case '-':
			digit = -1
		case '0':
			digit = 0
		case '1':
			digit = 1
		case '2':
			digit = 2
		}

		dec += digit * pow(5, p)
		p++
	}

	return dec
}

func dec_to_snafu(dec int) string {
	snafu := strings.Split(strconv.FormatInt(int64(dec), 5), "")

	for i := len(snafu) - 1; i >= 0; i-- {
		if snafu[i] == "3" {
			snafu[i-1] = string(snafu[i-1][0] + 1)
			snafu[i] = "="
		} else if snafu[i] == "4" {
			snafu[i-1] = string(snafu[i-1][0] + 1)
			snafu[i] = "-"
		} else if snafu[i] == "5" {
			snafu[i-1] = string(snafu[i-1][0] + 1)
			snafu[i] = "0"
		}
	}

	return strings.Join(snafu, "")
}

func get_sum(data []string) int {
	total := 0
	for _, line := range data {
		total += snafu_to_dec(line)
	}
	return total
}

func main() {
	data := load_data("input.txt")
	sum := get_sum(data)
	fmt.Println(dec_to_snafu(sum))
}
