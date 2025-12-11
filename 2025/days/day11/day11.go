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

type CacheEntry struct {
	name string
	dac  bool
	fft  bool
}

var cache map[CacheEntry]int

func init() {
	cache = map[CacheEntry]int{}
}

func parse_data(data []string) (*Device, *Device) {
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

	return devices["you"], devices["svr"]
}

func dfs(device *Device, dac bool, fft bool, mode int) int {
	key := CacheEntry{
		name: device.name,
		dac:  dac,
		fft:  fft,
	}
	if val, ok := cache[key]; ok {
		return val
	}

	if device.name == "out" {
		if mode == 1 || dac && fft {
			return 1
		} else {
			return 0
		}
	}

	if device.name == "dac" {
		dac = true
	} else if device.name == "fft" {
		fft = true
	}

	paths := 0
	for _, link := range device.links {
		paths += dfs(link, dac, fft, mode)
	}

	cache[key] = paths
	return paths
}

func count_paths(device *Device) int {
	cache = map[CacheEntry]int{}
	return dfs(device, false, false, 1)
}

func count_paths_through(device *Device) int {
	cache = map[CacheEntry]int{}
	return dfs(device, false, false, 2)
}

func Run() {
	loader.Day = 11
	data := loader.GetStrings()
	you, svr := parse_data(data)
	part1 := count_paths(you)
	part2 := count_paths_through(svr)

	fmt.Printf("%d %d\n", part1, part2)
}
