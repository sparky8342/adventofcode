package day14

import (
	"fmt"
	"loader"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func mod(x, d int) int {
	rem := x % d
	if rem < 0 {
		rem += d
	}
	return rem
}

func parse_data(data []string) []Robot {
	r := regexp.MustCompile(".*?(\\d+).*?(\\d+).*?([\\-]{0,1}\\d+).*?([\\-]{0,1}\\d+)")

	robots := make([]Robot, len(data))
	for i, line := range data {
		matches := r.FindStringSubmatch(line)
		nums := make([]int, 4)
		for j := 1; j <= 4; j++ {
			n, err := strconv.Atoi(matches[j])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
				os.Exit(1)
			}
			nums[j-1] = n
		}
		robot := Robot{
			x:  nums[0],
			y:  nums[1],
			vx: nums[2],
			vy: nums[3],
		}
		robots[i] = robot
	}

	return robots
}

func move(robots []Robot, width int, height int, steps int) int {
	quads := make([]int, 4)
	middle_x := width / 2
	middle_y := height / 2
	for _, robot := range robots {
		robot.x = mod((robot.x + robot.vx*steps), width)
		robot.y = mod((robot.y + robot.vy*steps), height)
		if robot.x < middle_x && robot.y < middle_y {
			quads[0]++
		} else if robot.x > middle_x && robot.y < middle_y {
			quads[1]++
		} else if robot.x < middle_x && robot.y > middle_y {
			quads[2]++
		} else if robot.x > middle_x && robot.y > middle_y {
			quads[3]++
		}
	}
	return quads[0] * quads[1] * quads[2] * quads[3]
}

func Run() {
	loader.Day = 14
	data := loader.GetStrings()
	robots := parse_data(data)

	part1 := move(robots, 101, 103, 100)
	part2 := -1

	fmt.Printf("%d %d\n", part1, part2)
}
