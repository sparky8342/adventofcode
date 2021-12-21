package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type CacheKey struct {
	places [2]int
	scores [2]int
	player int
	rolls  int
}

var cache map[CacheKey][2]int

func init() {
	cache = make(map[CacheKey][2]int)
}

func get_data() (int, int) {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	player1 := int(lines[0][len(lines[0])-1] - '0')
	player2 := int(lines[1][len(lines[1])-1] - '0')
	return player1, player2
}

func play(player1 int, player2 int) (int, int) {
	p1, p2 := _play([2]int{player1, player2}, [2]int{0, 0}, 0, 0)
	return p1, p2
}

func _play(places [2]int, scores [2]int, player int, rolls int) (int, int) {
	key := CacheKey{places: places, scores: scores, player: player, rolls: rolls}
	if val, found := cache[key]; found {
		return val[0], val[1]
	}

	if scores[0] >= 21 {
		return 1, 0
	}
	if scores[1] >= 21 {
		return 0, 1
	}

	player1_total := 0
	player2_total := 0
	for i := 1; i <= 3; i++ {
		pl := [2]int{places[0], places[1]}
		sc := [2]int{scores[0], scores[1]}

		place := places[player] + i
		if place > 10 {
			place -= 10
		}
		pl[player] = place

		var p1, p2 int
		if rolls == 2 {
			sc[player] += pl[player]
			p1, p2 = _play(pl, sc, (player+1)%2, 0)
		} else {
			p1, p2 = _play(pl, sc, player, rolls+1)
		}
		player1_total += p1
		player2_total += p2
	}

	cache[key] = [2]int{player1_total, player2_total}
	return player1_total, player2_total
}

func main() {
	player1, player2 := get_data()

	// part 1
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

	// part 2
	p1, p2 := play(player1, player2)
	if p1 > p2 {
		fmt.Println(p1)
	} else {
		fmt.Println(p2)
	}
}
