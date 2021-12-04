package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Card [][]int

func unmarked_sum(card Card) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if card[i][j] != -1 {
				sum += card[i][j]
			}
		}
	}
	return sum
}

func check_win(card Card) bool {
	for i := 0; i < 5; i++ {
		win1 := true
		win2 := true
		for j := 0; j < 5; j++ {
			if card[i][j] != -1 {
				win1 = false
			}
			if card[j][i] != -1 {
				win2 = false
			}
			if !win1 && !win2 {
				break
			}
		}
		if win1 || win2 {
			return true
		}
	}
	return false
}

func update_card(card Card, number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if card[i][j] == number {
				card[i][j] = -1
				return
			}
		}
	}
}

func draw(numbers []int, cards []Card) (int, int) {
	winners := []int{}
	initial_amount := len(cards)
	for _, num := range numbers {
		next_cards := []Card{}
		for _, card := range cards {
			update_card(card, num)
			if check_win(card) {
				if len(cards) == initial_amount || len(cards) == 1 {
					score := unmarked_sum(card) * num
					winners = append(winners, score)
				}
			} else {
				next_cards = append(next_cards, card)
			}
		}
		if len(next_cards) == 0 {
			break
		}
		cards = next_cards
	}
	return winners[0], winners[1]
}

func get_data() ([]int, []Card) {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")

	draw := strings.Split(lines[0], ",")
	draw_numbers := []int{}
	for _, num_str := range draw {
		num, _ := strconv.Atoi(num_str)
		draw_numbers = append(draw_numbers, num)
	}

	cards := []Card{}
	for i := 2; i < len(lines); i += 6 {
		card := [][]int{}
		for j := i; j < i+5; j++ {
			row_str := strings.Fields(lines[j])
			row := []int{}
			for _, num_str := range row_str {
				num, _ := strconv.Atoi(num_str)
				row = append(row, num)
			}
			card = append(card, row)
		}
		cards = append(cards, card)
	}

	return draw_numbers, cards
}

func main() {
	numbers, cards := get_data()
	first_winner, last_winner := draw(numbers, cards)
	fmt.Println(first_winner)
	fmt.Println(last_winner)
}
