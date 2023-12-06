package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func parse_data(data []string) ([]Race, Race) {
	time_strs := strings.Fields(data[0])
	distance_strs := strings.Fields(data[1])

	// individual
	races := []Race{}
	for i := 1; i < len(time_strs); i++ {
		time, _ := strconv.Atoi(time_strs[i])
		distance, _ := strconv.Atoi(distance_strs[i])
		races = append(races, Race{time: time, distance: distance})
	}

	// combined
	time := 0
	for _, ru := range data[0] {
		if ru >= '0' && ru <= '9' {
			time = time*10 + int(ru-'0')
		}
	}
	distance := 0
	for _, ru := range data[1] {
		if ru >= '0' && ru <= '9' {
			distance = distance*10 + int(ru-'0')
		}
	}
	combined_race := Race{time: time, distance: distance}

	return races, combined_race
}

func race_wins(race Race) int {
	middle := race.time / 2

	// find min
	left := 0
	right := middle

	for left < right {
		mid := left + (right-left)/2
		dist := (race.time - mid) * mid
		if dist <= race.distance {
			left = mid + 1
		} else if dist > race.distance {
			right = mid
		}
	}
	min := left

	// find max
	left = middle
	right = race.time

	for left < right {
		mid := left + (right-left)/2
		dist := (race.time - mid) * mid
		if dist <= race.distance {
			right = mid
		} else if dist > race.distance {
			left = mid + 1
		}
	}

	max := left
	return max - min
}

func race_wins_product(races []Race) int {
	product := 1
	for _, race := range races {
		product *= race_wins(race)
	}
	return product
}

func main() {
	data := load_data("input.txt")
	races, combined_race := parse_data(data)
	fmt.Println(race_wins_product(races))
	fmt.Println(race_wins(combined_race))

}
