package loader

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var Day int

func get_filename() string {
	return fmt.Sprintf("inputs/day%d.txt", Day)
}

func GetStrings() []string {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func GetInts() []int {
	data, err := ioutil.ReadFile(get_filename())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	lines := strings.Split(string(data), "\n")
	ints := make([]int, len(lines))

	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		ints[i] = n
	}

	return ints
}
