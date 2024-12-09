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

func defrag_files(data []byte) int {
	disk := [][]int{}
	file := true
	for i, b := range data {
		amount := int(b - '0')
		if file {
			disk = append(disk, []int{amount, i / 2})
		} else if amount > 0 {
			disk = append(disk, []int{amount, -1})
		}
		file = !file
	}

	p2 := len(disk) - 1

	for p2 > 0 {
		for disk[p2][1] == -1 {
			p2--
		}

		p1 := 0
		for p1 < p2 && (disk[p1][1] != -1 || disk[p1][0] < disk[p2][0]) {
			p1++
		}
		if p1 == p2 {
			p2--
			continue
		}

		if disk[p1][0] == disk[p2][0] {
			disk[p1][1] = disk[p2][1]
			disk[p2][1] = -1
		} else if disk[p1][0] > disk[p2][0] {
			diff := disk[p1][0] - disk[p2][0]
			disk[p1][0] = disk[p2][0]
			disk[p1][1] = disk[p2][1]
			disk[p2][1] = -1
			disk = append(disk[:p1+2], disk[p1+1:]...)
			disk[p1+1] = []int{diff, -1}
		}
	}

	checksum := 0
	pos := 0
	for _, entry := range disk {
		if entry[1] == -1 {
			pos += entry[0]
		} else {
			for i := 0; i < entry[0]; i++ {
				checksum += pos * entry[1]
				pos++
			}
		}
	}

	return checksum
}

func Run() {
	loader.Day = 9
	data := loader.GetOneLine()

	part1 := defrag(data)
	part2 := defrag_files(data)

	fmt.Printf("%d %d\n", part1, part2)
}
