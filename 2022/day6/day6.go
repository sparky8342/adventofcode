package main

import (
	"fmt"
	"io/ioutil"
)

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func find_marker(data string, no_chars int) int {
	letters := map[byte]int{}
	for i := 0; i < no_chars; i++ {
		letters[data[i]]++
	}

	for i := no_chars; i < len(data); i++ {
		if len(letters) == no_chars {
			return i
		}

		letters[data[i]]++

		previous := data[i-no_chars]
		letters[previous]--
		if letters[previous] == 0 {
			delete(letters, previous)
		}
	}
	return -1
}

func main() {
	data := load_data("input.txt")

	marker := find_marker(data, 4)
	fmt.Println(marker)

	marker = find_marker(data, 14)
	fmt.Println(marker)
}
