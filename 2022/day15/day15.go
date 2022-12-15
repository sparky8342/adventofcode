package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const ROW = 2000000
const BEACON_MIN = 0
const BEACON_MAX = 4000000
const MULTIPLIER = 4000000

const True = 1
const False = 2
const Beacon = 3

type Sensor struct {
	x                  int
	y                  int
	beacon_x           int
	beacon_y           int
	manhattan_distance int
}

type Area struct {
	sensors []Sensor
	x_min   int
	x_max   int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}

func get_manhattan_distance(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func parse_data(data []string) Area {
	r := regexp.MustCompile(".*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+).*?([\\-0-9]+)")
	sensors := []Sensor{}
	x_min := math.MaxInt32
	x_max := math.MinInt32

	for _, line := range data {
		match := r.FindStringSubmatch(line)
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		beacon_x, _ := strconv.Atoi(match[3])
		beacon_y, _ := strconv.Atoi(match[4])
		manhattan_distance := get_manhattan_distance(x, y, beacon_x, beacon_y)
		sensors = append(sensors, Sensor{x: x, y: y, beacon_x: beacon_x, beacon_y: beacon_y, manhattan_distance: manhattan_distance})
		x_min = min(min(x_min, x), beacon_x)
		x_max = max(max(x_max, x), beacon_x)
	}

	return Area{sensors: sensors, x_min: x_min, x_max: x_max}
}

func location_invalid(area Area, x int, y int) int {
	for _, sensor := range area.sensors {
		if x == sensor.beacon_x && y == sensor.beacon_y {
			return Beacon
		}
		dist := get_manhattan_distance(x, y, sensor.x, sensor.y)
		if dist <= sensor.manhattan_distance {
			return True
		}
	}
	return False
}

func invalid_line(area Area, y int) int {
	invalid := 0
	for x := area.x_min; x < area.x_max; x++ {
		if location_invalid(area, x, y) == True {
			invalid++
		}
	}
	for x := area.x_min - 1; location_invalid(area, x, y) == True; x-- {
		invalid++
	}
	for x := area.x_max; location_invalid(area, x, y) == True; x++ {
		invalid++
	}
	return invalid
}

func find_beacon(area Area, min int, max int) int {
	// check spaces around edge of sensor range
	for _, sensor := range area.sensors {
		for x, y := sensor.x, sensor.y-sensor.manhattan_distance-1; x < sensor.x+sensor.manhattan_distance+1 && y <= sensor.y; x, y = x+1, y+1 {
			if x < min || x > max || y < min || y > max {
				continue
			}
			if location_invalid(area, x, y) == False {
				return x*MULTIPLIER + y
			}
		}

		for x, y := sensor.x+sensor.manhattan_distance+1, sensor.y; x >= sensor.x && y < sensor.y+sensor.manhattan_distance+1; x, y = x-1, y+1 {
			if x < min || x > max || y < min || y > max {
				continue
			}
			if location_invalid(area, x, y) == False {
				return x*MULTIPLIER + y
			}
		}

		for x, y := sensor.x, sensor.y+sensor.manhattan_distance+1; x > sensor.x-sensor.manhattan_distance-1 && y >= sensor.y; x, y = x-1, y-1 {
			if x < min || x > max || y < min || y > max {
				continue
			}
			if location_invalid(area, x, y) == False {
				return x*MULTIPLIER + y
			}
		}

		for x, y := sensor.x-sensor.manhattan_distance-1, sensor.y; x <= sensor.x && y > sensor.y-sensor.manhattan_distance-1; x, y = x+1, y-1 {
			if x < min || x > max || y < min || y > max {
				continue
			}
			if location_invalid(area, x, y) == False {
				return x*MULTIPLIER + y
			}
		}
	}

	return 0
}

func main() {
	data := load_data("input.txt")
	area := parse_data(data)
	fmt.Println(invalid_line(area, ROW))

	fmt.Println(find_beacon(area, BEACON_MIN, BEACON_MAX))
}
