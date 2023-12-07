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
	cards  string
	cards2 string
	typ    int
	typ2   int
	bid    int
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
	case 'j':
		{
			return 1
		}
	default:
		{
			return int(card - '0')
		}
	}
}

func convert_wildcards(cards string) (string, string) {
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
		return cards, cards
	}

	bytes := []byte(cards)
	bytes2 := []byte(cards)
	b := byte(highest_card)
	for i := 0; i < 5; i++ {
		if bytes[i] == 'J' {
			bytes[i] = b
		}
		if bytes2[i] == 'J' {
			bytes2[i] = 'j'
		}
	}

	return string(bytes), string(bytes2)
}

func parse_data(data []string) []Hand {
	hands := []Hand{}
	for _, line := range data {
		parts := strings.Fields(line)
		cards := parts[0]
		cards_converted, cards_with_joker := convert_wildcards(cards)
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{
			cards:  cards,
			cards2: cards_with_joker,
			typ:    get_type(cards),
			typ2:   get_type(cards_converted),
			bid:    bid,
		})
	}
	return hands
}

func comp_cards(a string, b string) bool {
	for i := 0; i < 5; i++ {
		card_a := card_to_num(a[i])
		card_b := card_to_num(b[i])
		if card_a != card_b {
			return card_a > card_b
		}
	}
	return true
}

func winnings(hands []Hand) (int, int) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].typ == hands[j].typ {
			return comp_cards(hands[i].cards, hands[j].cards)
		} else {
			return hands[i].typ > hands[j].typ
		}
	})
	score := 0
	for i := 0; i < len(hands); i++ {
		score += (len(hands) - i) * hands[i].bid
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].typ2 == hands[j].typ2 {
			return comp_cards(hands[i].cards2, hands[j].cards2)
		} else {
			return hands[i].typ2 > hands[j].typ2
		}
	})
	score_wild := 0
	for i := 0; i < len(hands); i++ {
		score_wild += (len(hands) - i) * hands[i].bid
	}

	return score, score_wild
}

func main() {
	data := load_data("input.txt")
	hands := parse_data(data)
	score, wildcard_score := winnings(hands)
	fmt.Println(score)
	fmt.Println(wildcard_score)
}
