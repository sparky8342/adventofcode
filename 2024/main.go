package main

import (
	"days/day1"
	"days/day10"
	"days/day11"
	"days/day12"
	"days/day13"
	"days/day14"
	"days/day15"
	"days/day16"
	"days/day17"
	"days/day18"
	"days/day19"
	"days/day2"
	"days/day20"
	"days/day21"
	"days/day22"
	"days/day23"
	"days/day3"
	"days/day4"
	"days/day5"
	"days/day6"
	"days/day7"
	"days/day8"
	"days/day9"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		part, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %v\n", err)
			os.Exit(1)
		}
		switch part {
		case 1:
			day1.Run()
		case 2:
			day2.Run()
		case 3:
			day3.Run()
		case 4:
			day4.Run()
		case 5:
			day5.Run()
		case 6:
			day6.Run()
		case 7:
			day7.Run()
		case 8:
			day8.Run()
		case 9:
			day9.Run()
		case 10:
			day10.Run()
		case 11:
			day11.Run()
		case 12:
			day12.Run()
		case 13:
			day13.Run()
		case 14:
			day14.Run()
		case 15:
			day15.Run()
		case 16:
			day16.Run()
		case 17:
			day17.Run()
		case 18:
			day18.Run()
		case 19:
			day19.Run()
		case 20:
			day20.Run()
		case 21:
			day21.Run()
		case 22:
			day22.Run()
		case 23:
			day23.Run()
		}
	} else {
		fmt.Println("Day 1")
		day1.Run()
		fmt.Println("Day 2")
		day2.Run()
		fmt.Println("Day 3")
		day3.Run()
		fmt.Println("Day 4")
		day4.Run()
		fmt.Println("Day 5")
		day5.Run()
		fmt.Println("Day 6")
		day6.Run()
		fmt.Println("Day 7")
		day7.Run()
		fmt.Println("Day 8")
		day8.Run()
		fmt.Println("Day 9")
		day9.Run()
		fmt.Println("Day 10")
		day10.Run()
		fmt.Println("Day 11")
		day11.Run()
		fmt.Println("Day 12")
		day12.Run()
		fmt.Println("Day 13")
		day13.Run()
		fmt.Println("Day 14")
		day14.Run()
		fmt.Println("Day 15")
		day15.Run()
		fmt.Println("Day 16")
		day16.Run()
		fmt.Println("Day 17")
		day17.Run()
		fmt.Println("Day 18")
		day18.Run()
		fmt.Println("Day 19")
		day19.Run()
		fmt.Println("Day 20")
		day20.Run()
		fmt.Println("Day 21")
		day21.Run()
		fmt.Println("Day 22")
		day22.Run()
		fmt.Println("Day 23")
		day23.Run()
	}
}
