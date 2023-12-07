package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

type Hand struct {
	cards string
	typ   int
	bid   int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func get_type(hand string) int {
	distinct_cards := map[rune]int{}
	for _, ru := range hand {
		distinct_cards[ru]++
	}

	switch len(distinct_cards) {
	case 1:
		{
			return FIVE_OF_A_KIND
		}
	case 2:
		{
			for _, amount := range distinct_cards {
				if amount == 1 || amount == 4 {
					return FOUR_OF_A_KIND
				} else {
					return FULL_HOUSE
				}
			}
		}
	case 3:
		{
			for _, amount := range distinct_cards {
				if amount == 3 {
					return THREE_OF_A_KIND
				} else if amount == 2 {
					return TWO_PAIR
				}
			}
		}
	case 4:
		{
			return ONE_PAIR
		}
	case 5:
		{
			return HIGH_CARD
		}
	}

	return -1
}

func card_to_num(card byte) int {
	switch card {
	case 'A':
		{
			return 14
		}
	case 'K':
		{
			return 13
		}
	case 'Q':
		{
			return 12
		}
	case 'J':
		{
			return 11
		}
	case 'T':
		{
			return 10
		}
	default:
		{
			return int(card - '0')
		}
	}
}

func parse_data(data []string) []Hand {
	hands := []Hand{}
	for _, line := range data {
		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{cards: parts[0], typ: get_type(parts[0]), bid: bid})
	}
	return hands
}

func winnings(hands []Hand) int {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].typ == hands[j].typ {
			for k := 0; k < 5; k++ {
				a := card_to_num(hands[i].cards[k])
				b := card_to_num(hands[j].cards[k])
				if a != b {
					return a > b
				}
			}
			return true
		} else {
			return hands[i].typ > hands[j].typ
		}
	})

	score := 0
	for i := 0; i < len(hands); i++ {
		score += (len(hands) - i) * hands[i].bid
	}

	return score
}

func main() {
	data := load_data("input.txt")
	hands := parse_data(data)
	score := winnings(hands)
	fmt.Println(score)
}
