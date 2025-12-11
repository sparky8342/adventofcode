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

type ThroughCacheEntry struct {
	name string
	dac  bool
	fft  bool
}

var cache map[string]int
var throughcache map[ThroughCacheEntry]int

func init() {
	cache = map[string]int{}
	throughcache = map[ThroughCacheEntry]int{}
}

func parse_data(data []string, return_device string) *Device {
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

	return devices[return_device]
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

func count_paths_through(node *Device, dac bool, fft bool) int {
	key := ThroughCacheEntry{
		name: node.name,
		dac:  dac,
		fft:  fft,
	}

	if val, ok := throughcache[key]; ok {
		return val
	}

	if node.name == "out" {
		if dac && fft {
			return 1
		} else {
			return 0
		}
	}

	if node.name == "dac" {
		dac = true
	} else if node.name == "fft" {
		fft = true
	}

	paths := 0
	for _, link := range node.links {
		paths += count_paths_through(link, dac, fft)
	}

	throughcache[key] = paths
	return paths
}

func Run() {
	loader.Day = 11
	data := loader.GetStrings()
	you := parse_data(data, "you")
	part1 := count_paths(you)

	svr := parse_data(data, "svr")
	part2 := count_paths_through(svr, false, false)

	fmt.Printf("%d %d\n", part1, part2)
}
