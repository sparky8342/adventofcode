package day8

import (
	"fmt"
	"loader"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	x int
	y int
	z int
}

type intset map[int]struct{}

func (s intset) add(value int) {
	s[value] = struct{}{}
}

func (s intset) contains(value int) bool {
	_, ok := s[value]
	return ok
}

func union(a intset, b intset) intset {
	combined := intset{}
	for id := range a {
		combined.add(id)
	}
	for id := range b {
		combined.add(id)
	}
	return combined
}

func square(n int) int {
	return n * n
}

func parse_data(data []string) []Box {
	boxes := make([]Box, len(data))
	for i, line := range data {
		parts := strings.Split(line, ",")
		nums := make([]int, 3)
		for j := 0; j < 3; j++ {
			n, err := strconv.Atoi(parts[j])
			if err != nil {
				panic(err)
			}
			nums[j] = n
		}
		boxes[i] = Box{
			x: nums[0],
			y: nums[1],
			z: nums[2],
		}
	}

	return boxes
}

func connect_boxes(boxes []Box, connections int) (int, int) {
	pair_dists := [][3]float64{}
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			pair_dists = append(pair_dists,
				[3]float64{
					float64(i),
					float64(j),
					math.Sqrt(float64(square(boxes[i].x-boxes[j].x) + square(boxes[i].y-boxes[j].y) + square(boxes[i].z-boxes[j].z))),
				},
			)
		}
	}

	sort.Slice(pair_dists, func(i, j int) bool {
		return pair_dists[i][2] < pair_dists[j][2]
	})

	sets := make([]intset, len(boxes))
	for i := range boxes {
		s := intset{}
		s.add(i)
		sets[i] = s
	}

	var i int
	for i = 0; i < connections; i++ {
		box1, box2 := int(pair_dists[i][0]), int(pair_dists[i][1])
		combine := []int{}
		for j, set := range sets {
			if set.contains(box1) || set.contains(box2) {
				combine = append(combine, j)
				if len(combine) == 2 {
					break
				}
			}
		}
		if len(combine) == 2 {
			first, second := combine[0], combine[1]
			combined := union(sets[first], sets[second])
			sets = append(sets[0:second], sets[second+1:]...)
			sets = append(sets[0:first], sets[first+1:]...)
			sets = append(sets, combined)
		}
	}

	sort.Slice(sets, func(i, j int) bool {
		return len(sets[i]) > len(sets[j])
	})

	part1 := len(sets[0]) * len(sets[1]) * len(sets[2])
	var part2 int

	for ; ; i++ {
		box1, box2 := int(pair_dists[i][0]), int(pair_dists[i][1])
		combine := []int{}
		for j, set := range sets {
			if set.contains(box1) || set.contains(box2) {
				combine = append(combine, j)
				if len(combine) == 2 {
					break
				}
			}
		}
		if len(combine) == 2 {
			first, second := combine[0], combine[1]
			combined := union(sets[first], sets[second])
			sets = append(sets[0:second], sets[second+1:]...)
			sets = append(sets[0:first], sets[first+1:]...)
			sets = append(sets, combined)

			if len(sets) == 1 {
				part2 = boxes[box1].x * boxes[box2].x
				break
			}
		}
	}

	return part1, part2
}

func Run() {
	loader.Day = 8
	data := loader.GetStrings()
	boxes := parse_data(data)
	part1, part2 := connect_boxes(boxes, 1000)

	fmt.Printf("%d %d\n", part1, part2)
}
