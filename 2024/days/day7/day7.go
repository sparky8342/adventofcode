package day7

import (
	"fmt"
	"loader"
	"os"
	"strconv"
	"strings"
)

func parse_data(data []string) [][]int {
	equations := make([][]int, len(data))
	for i, line := range data {
		strs := strings.Split(line, " ")
		strs[0] = strs[0][:len(strs[0])-1]
		equations[i] = make([]int, len(strs))
		for j, n_str := range strs {
			n, err := strconv.Atoi(n_str)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %vn", err)
				os.Exit(1)
			}
			equations[i][j] = n
		}
	}
	return equations
}

func get_sequences(n int, use_concat bool) [][][]byte {
	sequences := make([][][]byte, n)
	sequences[1] = [][]byte{
		[]byte{'+'},
		[]byte{'*'},
	}
	symbols := []byte{'+', '*'}

	if use_concat {
		sequences[1] = append(sequences[1], []byte{'|'})
		symbols = append(symbols, '|')
	}

	for i := 2; i < n; i++ {
		for _, seq := range sequences[i-1] {
			for _, symbol := range symbols {
				cpy := make([]byte, len(seq))
				copy(cpy, seq)
				cpy = append(cpy, symbol)
				sequences[i] = append(sequences[i], cpy)
			}
		}
	}

	return sequences
}

func total_calibration(equations [][]int, use_concat bool) int {
	total := 0

	max := 0
	for _, equation := range equations {
		if len(equation) > max {
			max = len(equation)
		}
	}
	sequences := get_sequences(max+1, use_concat)

	for _, equation := range equations {
		target := equation[0]

		for _, seq := range sequences[len(equation)-2] {
			n := equation[1]
			for i := 2; i < len(equation); i++ {
				if seq[i-2] == '+' {
					n = n + equation[i]
				} else if seq[i-2] == '*' {
					n = n * equation[i]
				} else if seq[i-2] == '|' {
					next := equation[i]
					for next > 0 {
						next /= 10
						n *= 10
					}
					n = n + equation[i]
				}
			}
			if n == target {
				total += target
				break
			}
		}
	}

	return total
}

func Run() {
	loader.Day = 7
	data := loader.GetStrings()
	equations := parse_data(data)

	part1 := total_calibration(equations, false)
	part2 := total_calibration(equations, true)

	fmt.Printf("%d %d\n", part1, part2)
}
