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
				modules[dest].inputs[name] = false
			}
		}
	}

	return modules
}

func press_button(modules map[string]Module, amount int) int {
	low_sent := 0
	high_sent := 0

	for i := 0; i < amount; i++ {
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
					modules[module.name] = module
					send := module.on

					for _, dest := range module.destinations {
						queue = append(queue, Pulse{source: module.name, destination: dest, value: send})
					}
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
				}
				if send == false {
					low_sent += len(module.destinations)
				} else {
					high_sent += len(module.destinations)
				}
			}
		}
	}

	return low_sent * high_sent
}

func main() {
	data := load_data("input.txt")
	modules := parse_data(data)
	fmt.Println(press_button(modules, 1000))
}
