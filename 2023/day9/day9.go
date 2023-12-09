package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func load_data(filename string) [][]int {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	lines := strings.Split(string(data), "\n")
	nums := [][]int{}

	for _, line := range lines {
		row := []int{}
		for _, n_str := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(n_str)
			row = append(row, n)
		}
		nums = append(nums, row)
	}

	return nums
}

func find_next_value(row []int) int {
	sequences := [][]int{row}

	seq_no := 0
	for {
		previous := sequences[seq_no]
		all_zeroes := true
		seq := []int{}
		for i := 0; i < len(previous)-1; i++ {
			diff := previous[i+1] - previous[i]
			if diff != 0 {
				all_zeroes = false
			}
			seq = append(seq, diff)
		}
		sequences = append(sequences, seq)
		if all_zeroes {
			break
		}
		seq_no++
	}

	sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)

	for i := len(sequences) - 2; i >= 0; i-- {
		last := sequences[i][len(sequences[i])-1]
		last_below := sequences[i+1][len(sequences[i+1])-1]
		sequences[i] = append(sequences[i], last+last_below)
	}

	return sequences[0][len(sequences[0])-1]
}

func calculate_sum(nums [][]int) int {
	sum := 0
	for _, row := range nums {
		sum += find_next_value(row)
	}
	return sum
}

func main() {
	nums := load_data("input.txt")
	fmt.Println(calculate_sum(nums))
}
