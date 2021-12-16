package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

var packet_version_total int64
var bin_str string

func get_data() string {
	data, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSuffix(string(data), "\n")
	return line
}

func parse_bin() *big.Int {
	packet_version, _ := strconv.ParseInt(bin_str[0:3], 2, 64)
	packet_type_id, _ := strconv.ParseInt(bin_str[3:6], 2, 64)

	packet_version_total += packet_version

	var values []*big.Int
	if packet_type_id != 4 {
		length_type_id := bin_str[6]
		if length_type_id == '0' {
			length, _ := strconv.ParseInt(bin_str[7:22], 2, 64)
			bin_str = bin_str[22:]
			bin_length := len(bin_str)

			for bin_length-len(bin_str) < int(length) {
				// check for padding zeroes
				check, _ := strconv.ParseInt(bin_str, 2, 64)
				if check == 0 {
					break
				}
				values = append(values, parse_bin())
			}
		} else {
			sub_packets, _ := strconv.ParseInt(bin_str[7:18], 2, 64)
			bin_str = bin_str[18:]
			for len(values) < int(sub_packets) {
				values = append(values, parse_bin())
			}
		}
	}

	switch packet_type_id {
	case 0:
		sum := big.NewInt(0)
		for _, value := range values {
			sum.Add(sum, value)
		}
		return sum
	case 1:
		product := big.NewInt(1)
		for _, value := range values {
			product.Mul(product, value)
		}
		return product
	case 2:
		min := big.NewInt(999999)
		for _, value := range values {
			if value.Cmp(min) == -1 {
				min = value
			}
		}
		return min
	case 3:
		max := big.NewInt(0)
		for _, value := range values {
			if value.Cmp(max) == 1 {
				max = value
			}
		}
		return max
	case 4: // literal value
		value_str := ""
		i := 6
		var start byte
		start = '1'
		for start == '1' {
			start = bin_str[i]
			value_str += bin_str[i+1 : i+5]
			i += 5
		}
		value, _ := strconv.ParseInt(value_str, 2, 64)
		bin_str = bin_str[i:]
		return big.NewInt(value)
	case 5:
		if values[0].Cmp(values[1]) == 1 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	case 6:
		if values[0].Cmp(values[1]) == -1 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	case 7:
		if values[0].Cmp(values[1]) == 0 {
			return big.NewInt(1)
		} else {
			return big.NewInt(0)
		}
	}

	return big.NewInt(0)
}

func main() {
	line := get_data()

	bin_str = ""
	for _, ru := range line {
		num, _ := strconv.ParseUint(string(ru), 16, 32)
		bin_str += fmt.Sprintf("%04b", num)
	}

	packet_version_total = 0

	val := parse_bin()

	fmt.Println(packet_version_total)
	fmt.Println(val)
}
