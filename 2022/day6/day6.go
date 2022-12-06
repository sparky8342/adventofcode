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

func find_marker(data string) int {
	letters := map[byte]int{}
	for i := 0; i < 4; i++ {
		letters[data[i]]++
	}

	for i := 4; i < len(data); i++ {
		if len(letters) == 4 {
			return i
		}

		letters[data[i]]++

		previous := data[i-4]
		letters[previous]--
		if letters[previous] == 0 {
			delete(letters, previous)
		}
	}
	return -1
}

func main() {
	data := load_data("input.txt")
	marker := find_marker(data)
	fmt.Println(marker)
}
