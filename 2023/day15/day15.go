package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func hash(str string) int {
	h := 0
	for _, ru := range str {
		h = ((h + int(ru)) * 17) % 256
	}
	return h
}

func hash_sum(data string) int {
	sum := 0
	for _, str := range strings.Split(data, ",") {
		sum += hash(str)
	}
	return sum
}

func main() {
	data := load_data("input.txt")
	fmt.Println(hash_sum(data))
}
