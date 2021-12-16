package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func get_data() string {
	data, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSuffix(string(data), "\n")
	return line
}

func parse_bin(bin_str string, packet_version_total *int64) string {
	// check for 0 padding at end
	check, _ := strconv.ParseInt(bin_str, 2, 64)
	if check == 0 {
		return ""
	}

	packet_version, _ := strconv.ParseInt(bin_str[0:3], 2, 64)
	packet_type_id, _ := strconv.ParseInt(bin_str[3:6], 2, 64)
	is_literal_value := packet_type_id == 4

	*packet_version_total += packet_version

	if is_literal_value {
		value_str := ""
		i := 6
		var start byte
		start = '1'
		for start == '1' {
			start = bin_str[i]
			value_str += bin_str[i+1:i+5]
			i += 5
		}
		value, _ := strconv.ParseInt(value_str, 2, 64)
		_ = value
		return bin_str[i:]
	} else {
		// is operator packet
		length_type_id := bin_str[6]

		if length_type_id == '0' {
			return bin_str[22:]
		} else if length_type_id == '1' {
			return bin_str[18:]
		}
	}

	return ""
}

func main() {
	line := get_data()

	bin_str := ""
	for _, ru := range line {
		num, _ := strconv.ParseUint(string(ru), 16, 32)
		bin_str += fmt.Sprintf("%04b", num)
	}

	var packet_version_total int64
	packet_version_total = 0

	for len(bin_str) > 0 {
		bin_str = parse_bin(bin_str, &packet_version_total)
	}

	fmt.Println(packet_version_total)
}
