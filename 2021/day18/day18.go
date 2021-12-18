package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func get_data() [][]string {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")

	numbers := [][]string{}
	for _, line := range lines {
		num := strings.Split(line, "")
		numbers = append(numbers, num)
	}

	return numbers
}

func add(num1 []string, num2 []string) []string {
	added := append([]string{"["}, num1...)
	added = append(added, ",")
	added = append(added, num2...)
	added = append(added, "]")
	return added
}

func explode(num []string) ([]string, bool) {
	exploded := false
	for {
		level := 0
		explode_place := -1
		for i, ch := range num {
			if ch == "[" {
				level++
				if level == 5 {
					explode_place = i
					exploded = true
					break
				}
			} else if ch == "]" {
				level--
			}
		}

		if explode_place == -1 {
			return num, exploded
		}

		first_part := num[0:explode_place]
		for i := len(first_part) - 1; i >= 0; i-- {
			ch := first_part[i]
			if ch == "[" || ch == "]" || ch == "," {
				continue
			}
			a, _ := strconv.Atoi(ch)
			b, _ := strconv.Atoi(num[explode_place+1])
			first_part[i] = strconv.Itoa(a + b)
			break
		}
		first_part = append(first_part, "0")

		second_part := num[explode_place+5:]
		for i := 0; i < len(second_part); i++ {
			ch := second_part[i]
			if ch == "[" || ch == "]" || ch == "," {
				continue
			}
			a, _ := strconv.Atoi(ch)
			b, _ := strconv.Atoi(num[explode_place+3])
			second_part[i] = strconv.Itoa(a + b)
			break
		}

		new_num := append(first_part, second_part...)
		num = new_num
	}
}

func split(num []string) ([]string, bool) {
	split_point := -1
	n := 0
	for i, ch := range num {
		if ch == "[" || ch == "]" || ch == "," {
			continue
		}
		n, _ = strconv.Atoi(ch)
		if n >= 10 {
			split_point = i
			break
		}
	}

	if split_point > -1 {
		val1 := n / 2
		val2 := (n + 1) / 2

		new_num := []string{}
		for i, ch := range num {
			if i == split_point {
				new_num = append(new_num, []string{"[", strconv.Itoa(val1), ",", strconv.Itoa(val2), "]"}...)
			} else {
				new_num = append(new_num, ch)
			}
		}
		return new_num, true
	}

	return num, false
}

func reduce(num []string) []string {
	exploded := true
	was_split := true
	for exploded || was_split {
		num, exploded = explode(num)
		num, was_split = split(num)
	}
	return num
}

func magnitude(num []string) int {
	changed := true
	for changed == true {
		changed = false

		new_num := []string{}

		for i := 0; i < len(num); i++ {
			if num[i] == "[" && num[i+4] == "]" {
				a, _ := strconv.Atoi(num[i+1])
				b, _ := strconv.Atoi(num[i+3])
				mag := 3*a + 2*b
				new_num = append(new_num, strconv.Itoa(mag))
				i += 4
				changed = true
			} else {
				new_num = append(new_num, num[i])
			}
		}
		num = new_num
	}
	n, _ := strconv.Atoi(num[0])
	return n
}

func main() {
	numbers := get_data()

	total := numbers[0]
	for i := 1; i < len(numbers); i++ {
		total = add(total, numbers[i])
		total = reduce(total)
	}

	mag := magnitude(total)
	fmt.Println(mag)

	max_mag := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			total := reduce(add(numbers[i], numbers[j]))
			mag := magnitude(total)
			if mag > max_mag {
				max_mag = mag
			}
		}
	}
	fmt.Println(max_mag)
}
