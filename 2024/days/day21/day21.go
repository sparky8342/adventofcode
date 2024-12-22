package day21

import (
	"fmt"
	"loader"
	"math"
	"sort"
	"strings"
)

type CacheEntry struct {
	sequence   string
	iterations int
}

var cache map[CacheEntry]int

var buttons = []byte{'^', 'v', '<', '>'}

func init() {
	cache = map[CacheEntry]int{}
}

func keypad_button(pos byte, button byte) byte {
	switch pos {
	case '0':
		switch button {
		case '^':
			return '2'
		case '>':
			return 'A'
		default:
			return ' '
		}
	case '1':
		switch button {
		case '^':
			return '4'
		case '>':
			return '2'
		default:
			return ' '
		}
	case '2':
		switch button {
		case '^':
			return '5'
		case '>':
			return '3'
		case 'v':
			return '0'
		case '<':
			return '1'
		}
	case '3':
		switch button {
		case '^':
			return '6'
		case 'v':
			return 'A'
		case '<':
			return '2'
		case '>':
			return ' '
		}
	case '4':
		switch button {
		case '^':
			return '7'
		case 'v':
			return '1'
		case '<':
			return ' '
		case '>':
			return '5'
		}
	case '5':
		switch button {
		case '^':
			return '8'
		case 'v':
			return '2'
		case '<':
			return '4'
		case '>':
			return '6'
		}
	case '6':
		switch button {
		case '^':
			return '9'
		case 'v':
			return '3'
		case '<':
			return '5'
		case '>':
			return ' '
		}
	case '7':
		switch button {
		case '>':
			return '8'
		case 'v':
			return '4'
		default:
			return ' '
		}
	case '8':
		switch button {
		case '<':
			return '7'
		case '>':
			return '9'
		case 'v':
			return '5'
		case '^':
			return ' '
		}
	case '9':
		switch button {
		case '<':
			return '8'
		case 'v':
			return '6'
		default:
			return ' '
		}
	case 'A':
		switch button {
		case '^':
			return '3'
		case '<':
			return '0'
		default:
			return ' '
		}
	}
	return ' '
}

func direction_keypad(from byte, to byte) []byte {
	switch from {
	case '^':
		switch to {
		case 'v':
			return []byte{'v'}
		case '<':
			return []byte{'v', '<'}
		case '>':
			return []byte{'v', '>'}
		case 'A':
			return []byte{'>'}
		}
	case 'v':
		switch to {
		case 'v':
			return []byte{'^'}
		case '<':
			return []byte{'<'}
		case '>':
			return []byte{'>'}
		case 'A':
			return []byte{'^', '>'}
		}
	case '<':
		switch to {
		case 'v':
			return []byte{'>'}
		case '>':
			return []byte{'>', '>'}
		case '^':
			return []byte{'>', '^'}
		case 'A':
			return []byte{'>', '>', '^'}
		}
	case '>':
		switch to {
		case 'v':
			return []byte{'<'}
		case '<':
			return []byte{'<', '<'}
		case '^':
			return []byte{'<', '^'}
		case 'A':
			return []byte{'^'}
		}
	case 'A':
		switch to {
		case 'v':
			return []byte{'<', 'v'}
		case '<':
			return []byte{'v', '<', '<'}
		case '^':
			return []byte{'<'}
		case '>':
			return []byte{'v'}
		}
	}
	return []byte{}
}

func keypad_moves(from byte, to byte, seq []byte, visited map[byte]struct{}) [][]byte {
	if from == to {
		return [][]byte{seq}
	}

	sequences := [][]byte{}
	for _, button := range buttons {
		next := keypad_button(from, button)
		if _, ok := visited[next]; !ok {
			visited[next] = struct{}{}
			new_seq := make([]byte, len(seq))
			copy(new_seq, seq)
			new_seq = append(new_seq, button)
			sq := keypad_moves(next, to, new_seq, visited)
			sequences = append(sequences, sq...)
			delete(visited, next)
		}
	}

	return sequences
}

func find_sequence(code string, robots int) int {
	code = "A" + code

	sequences := [][]byte{[]byte{}}

	for i := 0; i < len(code)-1; i++ {
		visited := map[byte]struct{}{}
		seq := keypad_moves(code[i], code[i+1], []byte{}, visited)

		new_sequences := [][]byte{}
		for _, sequence := range sequences {
			for _, sq := range seq {
				new_seq := make([]byte, len(sequence))
				copy(new_seq, sequence)
				new_seq = append(new_seq, sq...)
				new_seq = append(new_seq, 'A')
				new_sequences = append(new_sequences, new_seq)
			}
		}
		sequences = new_sequences
	}

	sort.Slice(sequences, func(i, j int) bool {
		return len(sequences[i]) < len(sequences[j])
	})
	sequences = sequences[:100]

	shortest := math.MaxInt64

outer:
	for _, sequence := range sequences {
		l := sequence_length(sequence, robots)

		if l > shortest {
			continue outer
		}
		if l < shortest {
			shortest = l
		}
	}

	return shortest
}

func sequence_length(sequence []byte, iterations int) int {
	if iterations == 0 {
		return len(sequence)
	}

	entry := CacheEntry{sequence: string(sequence), iterations: iterations}
	if val, ok := cache[entry]; ok {
		return val
	}

	l := 0

	for _, part := range strings.SplitAfter(string(sequence), "A") {
		if part == "A" {
			l++
			continue
		}
		if len(part) == 0 {
			continue
		}

		var button byte = 'A'
		dir_sequence := []byte{}
		for _, ru := range part {
			b := byte(ru)
			if button != b {
				moves := direction_keypad(button, b)
				dir_sequence = append(dir_sequence, moves...)
			}
			dir_sequence = append(dir_sequence, 'A')
			button = b
		}

		l += sequence_length(dir_sequence, iterations-1)
	}

	cache[entry] = l
	return l
}

func find_direction_sequence(sequence []byte) []byte {
	var button byte = 'A'
	dir_sequence := []byte{}

	for _, ru := range sequence {
		b := byte(ru)
		if button != b {
			moves := direction_keypad(button, b)
			dir_sequence = append(dir_sequence, moves...)
		}
		dir_sequence = append(dir_sequence, 'A')
		button = b
	}

	return dir_sequence
}

func find_sequences(data []string, robots int) int {
	complexity := 0
	for _, line := range data {
		l := find_sequence(line, robots)
		n := int(line[0]-'0')*100 + int(line[1]-'0')*10 + int(line[2]-'0')
		complexity += l * n
	}
	return complexity
}

func Run() {
	loader.Day = 21
	data := loader.GetStrings()

	part1 := find_sequences(data, 2)
	part2 := find_sequences(data, 25)

	fmt.Printf("%d %d\n", part1, part2)
}
