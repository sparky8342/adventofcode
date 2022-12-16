package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Valve struct {
	name         string
	rate         int
	destinations []*Valve
}

type Empty struct {
}

type State struct {
	location       *Valve
	open_valves    map[string]Empty
	minute         int
	total_pressure int
}

func (state *State) Clone(minute int) *State {
	new_state := State{
		location:       state.location,
		total_pressure: state.total_pressure,
		minute:         minute,
	}

	open := map[string]Empty{}
	for name, _ := range state.open_valves {
		open[name] = Empty{}
	}
	new_state.open_valves = open

	return &new_state
}

func (state *State) Hash() string {
	hash := state.location.name

	keys := make([]string, len(state.open_valves))
	i := 0
	for k := range state.open_valves {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, name := range keys {
		hash += name
	}

	hash += strconv.Itoa(state.total_pressure)

	return hash
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) (*Valve, int) {
	r := regexp.MustCompile("Valve (\\w+) has flow rate=(\\d+); tunnels? leads? to valves? (.*)")

	var start *Valve = nil
	valves_map := map[string]*Valve{}
	destinations_map := map[string][]string{}
	valves_to_switch_on := 0

	// make all valves first
	for _, line := range data {
		match := r.FindStringSubmatch(line)
		name := match[1]
		rate, _ := strconv.Atoi(match[2])

		valve := &Valve{name: name, rate: rate}
		valves_map[name] = valve

		destinations := match[3]

		for _, destination := range strings.Split(destinations, ", ") {
			destinations_map[name] = append(destinations_map[name], destination)
		}

		if name == "AA" {
			start = valve
		}

		if rate > 0 {
			valves_to_switch_on++
		}
	}

	// set destinations
	for name, valve := range valves_map {
		for _, destination_name := range destinations_map[name] {
			valve.destinations = append(valve.destinations, valves_map[destination_name])
		}
	}

	return start, valves_to_switch_on
}

func bfs(start *Valve, valves_to_switch_on int) int {
	start_state := &State{location: start}

	queue := []*State{start_state}
	visited := map[string]Empty{}

	best := 0

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if len(state.open_valves) == valves_to_switch_on || state.minute == 30 {
			if state.total_pressure > best {
				best = state.total_pressure
			}
			continue
		}

		// possible actions - turn on or move

		// turn on
		if _, on := state.open_valves[state.location.name]; !on {
			if state.location.rate > 0 {
				new_state := state.Clone(state.minute + 1)
				new_state.total_pressure = state.total_pressure + state.location.rate*(30-state.minute-1)
				new_state.open_valves[state.location.name] = Empty{}

				hash := new_state.Hash()
				if _, seen := visited[hash]; !seen {
					visited[hash] = Empty{}
					queue = append(queue, new_state)
				}
			}
		}

		// moves
		for _, destination := range state.location.destinations {
			new_state := state.Clone(state.minute + 1)
			new_state.location = destination

			hash := new_state.Hash()
			if _, seen := visited[hash]; !seen {
				visited[hash] = Empty{}
				queue = append(queue, new_state)
			}
		}
	}

	return best
}

func main() {
	data := load_data("input.txt")
	start, valves_to_switch_on := parse_data(data)
	fmt.Println(bfs(start, valves_to_switch_on))
}
