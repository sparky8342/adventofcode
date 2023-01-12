package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const TIME_PART1 = 24
const TIME_PART2 = 32

const VISITED_CACHE_TIME = 30

type BluePrint struct {
	id             int
	ore            int
	clay           int
	obsidian_ore   int
	obsidian_clay  int
	geode_ore      int
	geode_obsidian int
	max_ore_robots int
}

type State struct {
	time                                                   int
	ore, clay, obsidian, geode                             int
	ore_robots, clay_robots, obsidian_robots, geode_robots int
}

type VisitedState struct {
	ore, clay, obsidian, geode                             int
	ore_robots, clay_robots, obsidian_robots, geode_robots int
}

func (state *State) collect_ore() {
	state.ore += state.ore_robots
	state.clay += state.clay_robots
	state.obsidian += state.obsidian_robots
	state.geode += state.geode_robots
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func parse_data(data []string) []BluePrint {
	r := regexp.MustCompile(".*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+)")
	blueprints := []BluePrint{}

	for _, line := range data {
		match := r.FindStringSubmatch(line)
		id, _ := strconv.Atoi(match[1])
		ore, _ := strconv.Atoi(match[2])
		clay, _ := strconv.Atoi(match[3])
		obsidian_ore, _ := strconv.Atoi(match[4])
		obsidian_clay, _ := strconv.Atoi(match[5])
		geode_ore, _ := strconv.Atoi(match[6])
		geode_obsidian, _ := strconv.Atoi(match[7])

		blueprint := BluePrint{
			id:             id,
			ore:            ore,
			clay:           clay,
			obsidian_ore:   obsidian_ore,
			obsidian_clay:  obsidian_clay,
			geode_ore:      geode_ore,
			geode_obsidian: geode_obsidian,
			max_ore_robots: max(clay, max(obsidian_ore, geode_ore)),
		}

		blueprints = append(blueprints, blueprint)
	}

	return blueprints
}

func visited_state(state State) VisitedState {
	return VisitedState{
		ore:             state.ore,
		clay:            state.clay,
		obsidian:        state.obsidian,
		geode:           state.geode,
		ore_robots:      state.ore_robots,
		clay_robots:     state.clay_robots,
		obsidian_robots: state.obsidian_robots,
		geode_robots:    state.geode_robots,
	}
}

func search(blueprint BluePrint, max_time int) int {
	start := State{ore_robots: 1}

	stack := []State{start}

	visited := map[VisitedState]int{}
	visited[visited_state(start)] = 0

	max := 0

	for len(stack) > 0 {
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if state.time == max_time {
			if state.geode > max {
				max = state.geode
			}
			continue
		}

		if state.ore >= blueprint.geode_ore && state.obsidian >= blueprint.geode_obsidian {
			new_state := state
			new_state.ore -= blueprint.geode_ore
			new_state.obsidian -= blueprint.geode_obsidian
			new_state.collect_ore()
			new_state.geode_robots++
			new_state.time++
			vstate := visited_state(new_state)
			if val, exists := visited[vstate]; !exists || val > new_state.time {
				stack = append(stack, new_state)
				if new_state.time < VISITED_CACHE_TIME {
					visited[vstate] = new_state.time
				}
			}
			continue
		}

		if state.ore >= blueprint.obsidian_ore && state.clay >= blueprint.obsidian_clay {
			new_state := state
			new_state.ore -= blueprint.obsidian_ore
			new_state.clay -= blueprint.obsidian_clay
			new_state.collect_ore()
			new_state.obsidian_robots++
			new_state.time++
			vstate := visited_state(new_state)
			if val, exists := visited[vstate]; !exists || val > new_state.time {
				stack = append(stack, new_state)
				if new_state.time < VISITED_CACHE_TIME {
					visited[vstate] = new_state.time
				}
			}
		}

		if state.ore >= blueprint.clay {
			new_state := state
			new_state.ore -= blueprint.clay
			new_state.collect_ore()
			new_state.clay_robots++
			new_state.time++
			vstate := visited_state(new_state)
			if val, exists := visited[vstate]; !exists || val > new_state.time {
				stack = append(stack, new_state)
				if new_state.time < VISITED_CACHE_TIME {
					visited[vstate] = new_state.time
				}
			}
		}

		if state.ore_robots < blueprint.max_ore_robots && state.ore >= blueprint.ore {
			new_state := state
			new_state.ore -= blueprint.ore
			new_state.collect_ore()
			new_state.ore_robots++
			new_state.time++
			vstate := visited_state(new_state)
			if val, exists := visited[vstate]; !exists || val > new_state.time {
				stack = append(stack, new_state)
				if new_state.time < VISITED_CACHE_TIME {
					visited[vstate] = new_state.time
				}
			}
		}

		new_state := state
		new_state.collect_ore()
		new_state.time++
		vstate := visited_state(new_state)
		if val, exists := visited[vstate]; !exists || val > new_state.time {
			stack = append(stack, new_state)
			if new_state.time < VISITED_CACHE_TIME {
				visited[vstate] = new_state.time
			}
		}

	}

	return max
}

func find_quality(blueprints []BluePrint) int {
	quality := 0
	for _, blueprint := range blueprints {
		quality += blueprint.id * search(blueprint, TIME_PART1)
	}
	return quality
}

func main() {
	data := load_data("input.txt")
	blueprints := parse_data(data)
	fmt.Println(find_quality(blueprints))
	fmt.Println(search(blueprints[0], TIME_PART2) * search(blueprints[1], TIME_PART2) * search(blueprints[2], TIME_PART2))
}
