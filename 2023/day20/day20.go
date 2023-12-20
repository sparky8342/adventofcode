package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	BROADCASTER = 1
	FLIP_FLOP   = 2
	CONJUNCTION = 3
)

type Module struct {
	name         string
	typ          int
	on           bool
	destinations []string
	inputs       map[string]bool
	presses      int
}

type Pulse struct {
	source      string
	destination string
	value       bool
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) map[string]Module {
	modules := map[string]Module{}

	for _, line := range data {
		parts := strings.Split(line, " -> ")
		module := Module{inputs: map[string]bool{}}

		if parts[0] == "broadcaster" {
			module.name = parts[0]
			module.typ = BROADCASTER
		} else {
			module.name = parts[0][1:]
			if parts[0][0] == '%' {
				module.typ = FLIP_FLOP
			} else if parts[0][0] == '&' {
				module.typ = CONJUNCTION
			}
		}

		module.destinations = strings.Split(parts[1], ", ")

		modules[module.name] = module
	}

	// setup inputs for conjunction modules
	for name, module := range modules {
		for _, dest := range module.destinations {
			dest_module := modules[dest]
			if dest_module.typ == CONJUNCTION {
				dest_module.inputs[name] = false
			}
			modules[dest] = dest_module
		}
	}

	return modules
}

func press_button(modules map[string]Module, amount int, find_rx bool) int {
	low_sent := 0
	high_sent := 0

	// specific solution for my input
	// when these 4 modules output a true
	// xn outputs a false to rx
	key_modules := map[string]int{}
	if find_rx {
		key_modules = map[string]int{"hn": 0, "mp": 0, "xf": 0, "fz": 0}
	}

	presses := 0
	for {
		presses++
		low_sent += 1 // button
		queue := []Pulse{Pulse{destination: "broadcaster", value: false}}

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			module := modules[pulse.destination]

			switch module.typ {
			case BROADCASTER:
				for _, dest := range module.destinations {
					queue = append(queue, Pulse{source: module.name, destination: dest, value: false})
				}
				low_sent += len(module.destinations)
			case FLIP_FLOP:
				if pulse.value == false {
					module.on = !module.on
					send := module.on

					for _, dest := range module.destinations {
						queue = append(queue, Pulse{source: module.name, destination: dest, value: send})
						if val, exists := key_modules[dest]; exists && val == 0 && send == false {
							key_modules[dest] = presses
						}
					}

					modules[module.name] = module

					if send == false {
						low_sent += len(module.destinations)
					} else {
						high_sent += len(module.destinations)
					}
				}
			case CONJUNCTION:
				var send bool
				module.inputs[pulse.source] = pulse.value
				if pulse.value == false {
					send = true
				} else {
					send = false
					for _, val := range module.inputs {
						if val == false {
							send = true
							break
						}
					}
				}

				for _, dest := range module.destinations {
					queue = append(queue, Pulse{source: module.name, destination: dest, value: send})
					if val, exists := key_modules[dest]; exists && val == 0 && send == false {
						key_modules[dest] = presses
					}

				}

				modules[module.name] = module

				if send == false {
					low_sent += len(module.destinations)
				} else {
					high_sent += len(module.destinations)
				}
			}
		}
		if !find_rx && presses == amount {
			break
		}

		if key_modules["hn"] > 0 && key_modules["mp"] > 0 && key_modules["xf"] > 0 && key_modules["fz"] > 0 {
			return lcm(key_modules["hn"], key_modules["mp"], key_modules["xf"], key_modules["fz"])
		}
	}

	return low_sent * high_sent
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func main() {
	data := load_data("input.txt")
	modules := parse_data(data)
	fmt.Println(press_button(modules, 1000, false))
	modules = parse_data(data)
	fmt.Println(press_button(modules, 1, true))
}
