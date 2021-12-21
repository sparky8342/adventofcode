package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func get_data() (int, int) {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	player1 := int(lines[0][len(lines[0])-1] - '0')
	player2 := int(lines[1][len(lines[1])-1] - '0')
	return player1, player2
}

func main() {
	player1, player2 := get_data()

	places := []int{player1, player2}
	scores := []int{0, 0}

	die := 0
	rolls := 0

outer:
	for {
		for player := 0; player <= 1; player++ {
			for i := 0; i < 3; i++ {
				die = die%100 + 1
				places[player] = (places[player]+die-1)%10 + 1
			}
			rolls += 3
			scores[player] += places[player]
			if scores[player] >= 1000 {
				fmt.Println(scores[(player+1)%2] * rolls)
				break outer
			}
		}
	}
}
