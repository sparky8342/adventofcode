package main

import (
	"days/day1"
	"days/day2"
	"days/day3"
	"days/day4"
	"days/day5"
	"days/day6"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		part, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
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
	}
}
