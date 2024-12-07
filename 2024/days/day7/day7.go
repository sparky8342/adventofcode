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

func get_sequences(n int) [][]byte {
	sequences := [][]byte{
		[]byte{},
	}

	for i := 0; i < n; i++ {
		new_sequences := [][]byte{}
		for _, seq := range sequences {
			for _, symbol := range []byte{'+', '*'} {
				cpy := make([]byte, len(seq))
				copy(cpy, seq)
				cpy = append(cpy, symbol)
				new_sequences = append(new_sequences, cpy)
			}
		}
		sequences = new_sequences
	}

	return sequences
}

func total_calibration(equations [][]int) int {
	total := 0

	for _, equation := range equations {
		target := equation[0]
		sequences := get_sequences(len(equation) - 2)

		for _, seq := range sequences {
			n := equation[1]
			for i := 2; i < len(equation); i++ {
				if seq[i-2] == '+' {
					n = n + equation[i]
				} else if seq[i-2] == '*' {
					n = n * equation[i]
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

	part1 := total_calibration(equations)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
