package day24

import (
	"fmt"
	"loader"
	"strings"
)

type Gate struct {
	typ    string
	inputs [2]string
	output string
}

func parse_data(data [][]string) (map[string]int, []Gate) {
	wires := map[string]int{}

	// initial wires
	for _, line := range data[0] {
		parts := strings.Split(line, ": ")
		wires[parts[0]] = int(parts[1][0] - '0')
	}

	// gates
	gates := make([]Gate, len(data[1]))
	for i, line := range data[1] {
		parts := strings.Split(line, " ")
		for _, wire := range []string{parts[0], parts[2], parts[4]} {
			if _, ok := wires[wire]; !ok {
				wires[wire] = -1
			}
		}
		gates[i] = Gate{
			typ:    parts[1],
			inputs: [2]string{parts[0], parts[2]},
			output: parts[4],
		}
	}

	return wires, gates
}

func process(wires map[string]int, gates []Gate) int {
	queue := make([]Gate, len(gates))
	copy(queue, gates)

	for len(queue) > 0 {
		gate := queue[0]
		queue = queue[1:]

		if wires[gate.inputs[0]] == -1 || wires[gate.inputs[1]] == -1 {
			queue = append(queue, gate)
			continue
		}

		switch gate.typ {
		case "AND":
			wires[gate.output] = wires[gate.inputs[0]] & wires[gate.inputs[1]]
		case "OR":
			wires[gate.output] = wires[gate.inputs[0]] | wires[gate.inputs[1]]
		case "XOR":
			wires[gate.output] = wires[gate.inputs[0]] ^ wires[gate.inputs[1]]
		}
	}

	n := 0
	wire_num := 0
	for {
		wire := fmt.Sprintf("z%02d", wire_num)
		if val, ok := wires[wire]; ok {
			n += val << wire_num
		} else {
			break
		}
		wire_num++
	}
	return n
}

func Run() {
	loader.Day = 24
	data := loader.GetStringGroups()
	wires, gates := parse_data(data)

	part1 := process(wires, gates)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
