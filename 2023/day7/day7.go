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
	typ2  int
	bid   int
}

func load_data(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func get_type(cards string) int {
	distinct_cards := map[rune]int{}
	for _, ru := range cards {
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

func card_to_num(card byte, wild bool) int {
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
			if wild {
				return 1
			} else {
				return 11
			}
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

func convert_wildcards(cards string) string {
	card_counts := map[rune]int{}
	var highest_card rune
	highest_amount := 0

	for _, ru := range cards {
		card_counts[ru]++
		if ru != 'J' && card_counts[ru] > highest_amount {
			highest_amount = card_counts[ru]
			highest_card = ru
		}
	}

	if _, exists := card_counts['J']; !exists {
		return cards
	}

	bytes := []byte(cards)
	b := byte(highest_card)
	for i := 0; i < 5; i++ {
		if bytes[i] == 'J' {
			bytes[i] = b
		}
	}

	return string(bytes)
}

func parse_data(data []string) []Hand {
	hands := []Hand{}
	for _, line := range data {
		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{
			cards: parts[0],
			typ:   get_type(parts[0]),
			typ2:  get_type(convert_wildcards(parts[0])),
			bid:   bid,
		})
	}
	return hands
}

func winnings(hands []Hand) (int, int) {
	scores := make([]int, 2)

	for typ := 0; typ <= 1; typ++ {
		wild := false
		if typ == 1 {
			wild = true
		}
		sort.Slice(hands, func(i, j int) bool {
			if (typ == 0 && hands[i].typ == hands[j].typ) || (typ == 1 && hands[i].typ2 == hands[j].typ2) {
				for k := 0; k < 5; k++ {
					a := card_to_num(hands[i].cards[k], wild)
					b := card_to_num(hands[j].cards[k], wild)
					if a != b {
						return a > b
					}
				}
				return true
			} else {
				return (typ == 0 && hands[i].typ > hands[j].typ) || (typ == 1 && hands[i].typ2 > hands[j].typ2)
			}
		})

		score := 0
		for i := 0; i < len(hands); i++ {
			score += (len(hands) - i) * hands[i].bid
		}
		scores[typ] = score
	}

	return scores[0], scores[1]
}

func main() {
	data := load_data("input.txt")
	hands := parse_data(data)
	score, wildcard_score := winnings(hands)
	fmt.Println(score)
	fmt.Println(wildcard_score)
}
