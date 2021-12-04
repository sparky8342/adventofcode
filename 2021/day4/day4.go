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
		win := true
		for j := 0; j < 5; j++ {
			if card[i][j] != -1 {
				win = false
				break
			}
		}
		if win {
			return true
		}
		win = true
		for j := 0; j < 5; j++ {
			if card[j][i] != -1 {
				win = false
				break
			}
		}
		if win {
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

func draw(numbers []int, cards []Card) {
	amount := len(cards)
	for _, num := range numbers {
		next_cards := []Card{}
		for _, card := range cards {
			update_card(card, num)
			if check_win(card) {
				if len(cards) == amount || len(cards) == 1 {
					sum := unmarked_sum(card)
					fmt.Println(sum * num)
				}
			} else {
				next_cards = append(next_cards, card)
			}
		}
		cards = next_cards
	}
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
	draw(numbers, cards)
}
