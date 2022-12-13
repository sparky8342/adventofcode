package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const True = 1
const False = 2
const Same = 3

func load_data(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return string(data)
}

func parse_data(data string) [][]string {
	pairs := [][]string{}

	sections := strings.Split(data, "\n\n")

	for _, section := range sections {
		pair := strings.Split(section, "\n")
		pairs = append(pairs, pair)
	}

	return pairs
}

func parse_packet(packet string) []interface{} {
	var result []interface{}
	json.Unmarshal([]byte(packet), &result)
	return result
}

func compare_pair(a string, b string) bool {
	a_data := parse_packet(a)
	b_data := parse_packet(b)
	r := compare(a_data, b_data)
	if r == False {
		return false
	}
	return true
}

func compare(a []interface{}, b []interface{}) int {
	pos := 0

	for pos < len(a) && pos < len(b) {
		a_val := a[pos]
		b_val := b[pos]

		a_val_type := fmt.Sprintf("%T", a_val)
		b_val_type := fmt.Sprintf("%T", b_val)

		if a_val_type != b_val_type {
			var result int
			if a_val_type == "float64" {
				result = compare([]interface{}{a_val}, b_val.([]interface{}))
			} else if b_val_type == "float64" {
				result = compare(a_val.([]interface{}), []interface{}{b_val})
			}
			if result != Same {
				return result
			}
		} else if a_val_type == "float64" && b_val_type == "float64" {
			if a_val.(float64) < b_val.(float64) {
				return True
			} else if a_val.(float64) > b_val.(float64) {
				return False
			}
		} else if a_val_type == "[]interface {}" && b_val_type == "[]interface {}" {
			result := compare(a_val.([]interface{}), b_val.([]interface{}))
			if result != Same {
				return result
			}
		}
		pos++
	}
	if len(a) < len(b) {
		return True
	} else if len(a) > len(b) {
		return False
	}

	return Same
}

func compare_pairs(pairs [][]string) int {
	sum := 0
	for i, pair := range pairs {
		result := compare_pair(pair[0], pair[1])
		if result {
			sum += i + 1
		}
	}
	return sum
}

func main() {
	data := load_data("input.txt")

	pairs := parse_data(data)
	sum := compare_pairs(pairs)
	fmt.Println(sum)
}
