package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

type Transmission struct {
	packet               string
	packet_version_total int64
}

func (tr *Transmission) parse() *big.Int {
	packet_version, _ := strconv.ParseInt(tr.packet[0:3], 2, 64)
	packet_type_id, _ := strconv.ParseInt(tr.packet[3:6], 2, 64)

	tr.packet_version_total += packet_version

	var values []*big.Int
	if packet_type_id != 4 {
		length_type_id := tr.packet[6]
		if length_type_id == '0' {
			length, _ := strconv.ParseInt(tr.packet[7:22], 2, 64)
			tr.packet = tr.packet[22:]
			bin_length := len(tr.packet)

			for bin_length-len(tr.packet) < int(length) {
				// check for padding zeroes
				check, _ := strconv.ParseInt(tr.packet, 2, 64)
				if check == 0 {
					break
				}
				values = append(values, tr.parse())
			}
		} else {
			sub_packets, _ := strconv.ParseInt(tr.packet[7:18], 2, 64)
			tr.packet = tr.packet[18:]
			for len(values) < int(sub_packets) {
				values = append(values, tr.parse())
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
		min := big.NewInt(0)
		first := true
		for _, value := range values {
			if first {
				min = value
				first = false
				continue
			}
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
			start = tr.packet[i]
			value_str += tr.packet[i+1 : i+5]
			i += 5
		}
		value, _ := strconv.ParseInt(value_str, 2, 64)
		tr.packet = tr.packet[i:]
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

func get_data() string {
	data, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSuffix(string(data), "\n")
	return line
}

func main() {
	line := get_data()

	tr := Transmission{packet_version_total: 0}
	for _, ru := range line {
		num, _ := strconv.ParseUint(string(ru), 16, 32)
		tr.packet += fmt.Sprintf("%04b", num)
	}

	val := tr.parse()

	fmt.Println(tr.packet_version_total)
	fmt.Println(val)
}
