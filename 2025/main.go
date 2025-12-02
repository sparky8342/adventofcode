package main

import (
	"days/day1"
	"days/day2"
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
		}
	} else {
		fmt.Println("Day 1")
		day1.Run()
		fmt.Println("Day 2")
		day2.Run()
	}
}
