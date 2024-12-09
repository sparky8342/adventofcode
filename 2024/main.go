package main

import (
	"days/day1"
	"days/day2"
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
	}
}
