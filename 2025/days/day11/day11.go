package day11

import (
	"fmt"
	"loader"
	"strings"
)

type Device struct {
	name  string
	links []*Device
}

var cache map[string]int

func init() {
	cache = map[string]int{}
}

func parse_data(data []string) *Device {
	devices := map[string]*Device{}

	for _, line := range data {
		line = strings.Replace(line, ":", "", 1)
		names := strings.Split(line, " ")
		for _, name := range names {
			if _, ok := devices[name]; !ok {
				devices[name] = &Device{name: name}
			}
		}
		for i := 1; i < len(names); i++ {
			devices[names[0]].links = append(devices[names[0]].links, devices[names[i]])
		}
	}

	return devices["you"]
}

func count_paths(node *Device) int {
	if val, ok := cache[node.name]; ok {
		return val
	}

	if node.name == "out" {
		return 1
	}

	paths := 0
	for _, link := range node.links {
		paths += count_paths(link)
	}

	cache[node.name] = paths
	return paths
}

func Run() {
	loader.Day = 11
	data := loader.GetStrings()
	you := parse_data(data)
	part1 := count_paths(you)

	fmt.Printf("%d %d\n", part1, 0)
}
