package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Ins struct {
	op  string
	reg string
	arg string
}

type Section []Ins

type Program struct {
	registers map[string]int
	sections  []Section
}

type State struct {
	w, y uint8
	x, z int32
	id   int
}

type DedupeState struct {
	w, y uint8
	x, z int32
}

func get_program() Program {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")

	sections := []Section{}
	instructions := Section{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		op := parts[0]
		reg := parts[1]
		arg := ""
		if op == "inp" && len(instructions) > 0 {
			sections = append(sections, instructions)
			instructions = []Ins{}
		}
		if op != "inp" {
			arg = parts[2]
		}
		instructions = append(instructions, Ins{op: op, reg: reg, arg: arg})
	}
	sections = append(sections, instructions)
	p := Program{registers: map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}, sections: sections}
	return p
}

func (state *State) get_arg(arg string) int {
	var num int
	switch arg {
	case "w":
		num = int(state.w)
	case "x":
		num = int(state.x)
	case "y":
		num = int(state.y)
	case "z":
		num = int(state.z)
	default:
		num, _ = strconv.Atoi(arg)
	}
	return num
}

func (program *Program) search() {
	for part := 1; part <= 2; part++ {
		states := []State{State{w: 0, x: 0, y: 0, z: 0, id: 0}}

		for _, section := range program.sections {
			for _, ins := range section {
				switch ins.op {
				case "inp":
					var i uint8
					new_states := []State{}
					for id, _ := range states {
						for i = 1; i <= 9; i++ {
							new_state := State{
								w:  i,
								x:  states[id].x,
								y:  states[id].y,
								z:  states[id].z,
								id: states[id].id*10 + int(i),
							}
							new_states = append(new_states, new_state)
						}
					}
					states = new_states
				case "add":
					for id, _ := range states {
						arg := states[id].get_arg(ins.arg)
						switch ins.reg {
						case "w":
							states[id].w += uint8(arg)
						case "x":
							states[id].x += int32(arg)
						case "y":
							states[id].y += uint8(arg)
						case "z":
							states[id].z += int32(arg)
						}
					}
				case "mul":
					for id, _ := range states {
						arg := states[id].get_arg(ins.arg)
						switch ins.reg {
						case "w":
							states[id].w *= uint8(arg)
						case "x":
							states[id].x *= int32(arg)
						case "y":
							states[id].y *= uint8(arg)
						case "z":
							states[id].z *= int32(arg)
						}
					}
				case "div":
					for id, _ := range states {
						arg := states[id].get_arg(ins.arg)
						switch ins.reg {
						case "w":
							states[id].w /= uint8(arg)
						case "x":
							states[id].x /= int32(arg)
						case "y":
							states[id].y /= uint8(arg)
						case "z":
							states[id].z /= int32(arg)
						}
					}
				case "mod":
					for id, _ := range states {
						arg := states[id].get_arg(ins.arg)
						switch ins.reg {
						case "w":
							states[id].w %= uint8(arg)
						case "x":
							states[id].x %= int32(arg)
						case "y":
							states[id].y %= uint8(arg)
						case "z":
							states[id].z %= int32(arg)
						}
					}
				case "eql":
					for id, _ := range states {
						arg := states[id].get_arg(ins.arg)
						switch ins.reg {
						case "w":
							if states[id].w == uint8(arg) {
								states[id].w = 1
							} else {
								states[id].w = 0
							}
						case "x":
							if states[id].x == int32(arg) {
								states[id].x = 1
							} else {
								states[id].x = 0
							}
						case "y":
							if states[id].y == uint8(arg) {
								states[id].y = 1
							} else {
								states[id].y = 0
							}
						case "z":
							if states[id].z == int32(arg) {
								states[id].z = 1
							} else {
								states[id].z = 0
							}
						}
					}
				}
			}
			mp := map[DedupeState]int{}
			for i, _ := range states {
				d := DedupeState{w: states[i].w, x: states[i].x, y: states[i].y, z: states[i].z}
				if val, ok := mp[d]; ok {
					if (part == 1 && states[i].id > val) || (part == 2 && states[i].id < val) {
						mp[d] = states[i].id
					}
				} else {
					mp[d] = states[i].id
				}
			}
			states = []State{}
			for k, v := range mp {
				state := State{w: k.w, x: k.x, y: k.y, z: k.z, id: v}
				states = append(states, state)
			}

		}

		var best int
		if part == 1 {
			best = 0
		} else {
			best = math.MaxInt64
		}
		for i, _ := range states {
			if states[i].z == 0 {
				if part == 1 {
					if states[i].id > best {
						best = states[i].id
					}
				} else if part == 2 {
					if states[i].id < best {
						best = states[i].id
					}
				}
			}
		}
		fmt.Println(best)
	}
}

func main() {
	program := get_program()
	program.search()
}
