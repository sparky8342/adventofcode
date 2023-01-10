package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const CUTOFF1 = 30
const CUTOFF2 = 26

type Valve struct {
	name         string
	rate         int
	destinations []*Valve
	open         bool
}

type Empty struct {
}

type Path struct {
	nodes    map[string]Empty
	pressure int
}

func get_hash(valve *Valve, open_valves map[string]Empty, pressure int, minute int) string {
	hash := valve.name

	keys := make([]string, len(open_valves))
	i := 0
	for k := range open_valves {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, name := range keys {
		hash += name
	}

	hash += strconv.Itoa(pressure) + ":" + strconv.Itoa(minute)

	return hash
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) *Valve {
	r := regexp.MustCompile("Valve (\\w+) has flow rate=(\\d+); tunnels? leads? to valves? (.*)")

	var start *Valve = nil
	valves_map := map[string]*Valve{}
	destinations_map := map[string][]string{}

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
	}

	// set destinations
	for name, valve := range valves_map {
		for _, destination_name := range destinations_map[name] {
			valve.destinations = append(valve.destinations, valves_map[destination_name])
		}
	}

	return start
}

func paths_cross(a Path, b Path) bool {
	for key := range a.nodes {
		if _, exists := b.nodes[key]; exists {
			return true
		}
	}
	return false
}

func find_best_pair(paths map[string]Path) int {
	best_pressure := 0

	path_slice := []Path{}
	for _, path := range paths {
		path_slice = append(path_slice, path)
	}

	sort.Slice(path_slice, func(i, j int) bool {
		return path_slice[i].pressure > path_slice[j].pressure
	})

	for i := 0; i < len(path_slice); i++ {
		for j := i + 1; j < len(path_slice); j++ {
			pressure := path_slice[i].pressure + path_slice[j].pressure
			if pressure < best_pressure {
				break
			}
			if paths_cross(path_slice[i], path_slice[j]) {
				continue
			}
			if pressure > best_pressure {
				best_pressure = pressure
			}
		}
	}

	return best_pressure
}

func find_best_pressure(start *Valve) (int, int) {
	best_pressure := 0
	open_valves := map[string]Empty{}
	visited := map[string]Empty{}
	paths := map[string]Path{}
	dfs(start, 0, 0, 0, &best_pressure, open_valves, visited, []string{}, paths)
	return best_pressure, find_best_pair(paths)
}

func dfs(valve *Valve, minute int, pressure int, pressure2 int, best_pressure *int, open_valves map[string]Empty, visited map[string]Empty, path []string, paths map[string]Path) {
	if minute == CUTOFF2 {
		p := strings.Join(path, " ")
		val, exists := paths[p]
		if !exists || val.pressure < pressure2 {
			new_p := Path{nodes: map[string]Empty{}, pressure: pressure2}
			for _, name := range path {
				new_p.nodes[name] = Empty{}
			}
			paths[p] = new_p
		}
	}
	if minute >= CUTOFF1 {
		if pressure > *best_pressure {
			*best_pressure = pressure
		}
		return
	}

	// use visited hash, but only before minute 25 to save memory
	if minute < 25 {
		hash := get_hash(valve, open_valves, pressure, minute)
		if _, exists := visited[hash]; exists {
			return
		}
		visited[hash] = Empty{}
	}

	// turn on
	if valve.rate > 0 && !valve.open {
		valve.open = true
		open_valves[valve.name] = Empty{}
		dfs(valve, minute+1, pressure+(CUTOFF1-minute-1)*valve.rate, pressure2+(CUTOFF2-minute-1)*valve.rate, best_pressure, open_valves, visited, append(path, valve.name), paths)
		valve.open = false
		delete(open_valves, valve.name)
	}

	// move
	for _, destination := range valve.destinations {
		dfs(destination, minute+1, pressure, pressure2, best_pressure, open_valves, visited, path, paths)
	}
}

func main() {
	data := load_data("input.txt")
	start := parse_data(data)
	best_pressure, best_pressure2 := find_best_pressure(start)
	fmt.Println(best_pressure)
	fmt.Println(best_pressure2)
}
