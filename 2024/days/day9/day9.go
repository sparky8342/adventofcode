package day9

import (
	"fmt"
	"loader"
)

func defrag(data []byte) int {
	disk := []int{}
	file := true
	for i, b := range data {
		if file {
			for j := 0; j < int(b-'0'); j++ {
				disk = append(disk, i/2)
			}
		} else {
			for j := 0; j < int(b-'0'); j++ {
				disk = append(disk, -1)
			}
		}
		file = !file
	}

	p1 := 0
	for disk[p1] != -1 {
		p1++
	}
	p2 := len(disk) - 1

	for p1 < p2 {
		disk[p1] = disk[p2]
		disk[p2] = -1
		for disk[p1] != -1 {
			p1++
		}
		for disk[p2] == -1 {
			p2--
		}
	}

	checksum := 0
	for i, n := range disk {
		if n == -1 {
			break
		}
		checksum += i * n
	}

	return checksum
}

func Run() {
	loader.Day = 9
	data := loader.GetOneLine()

	part1 := defrag(data)

	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
