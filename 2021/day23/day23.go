package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strings"
)

const POS_SIZE = 27

type Pos [POS_SIZE]int

type State struct {
	pos  Pos
	cost int
}

type Move struct {
	from int
	to   int
	cost int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

var nums map[byte]int
var destinations [][]int
var moves_out map[int]int
var costs [5]int

func init() {
	nums = map[byte]int{'A': 1, 'B': 2, 'C': 3, 'D': 4}
	destinations = [][]int{{0, 0}, {2, 14}, {4, 18}, {6, 22}, {8, 26}}
	moves_out = map[int]int{11: 2, 15: 4, 19: 6, 23: 8}
	costs = [5]int{0, 1, 10, 100, 1000}
}

func get_data() Pos {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	pos := Pos{}
	pos[11] = nums[lines[2][3]]
	pos[12] = nums[lines[3][3]]
	pos[15] = nums[lines[2][5]]
	pos[16] = nums[lines[3][5]]
	pos[19] = nums[lines[2][7]]
	pos[20] = nums[lines[3][7]]
	pos[23] = nums[lines[2][9]]
	pos[24] = nums[lines[3][9]]
	return pos
}

func (p *Pos) copy_pos() Pos {
	new_pos := Pos{}
	for i := 0; i < POS_SIZE; i++ {
		new_pos[i] = p[i]
	}
	return new_pos
}

func (pos *Pos) print_pos() {
	letters := []string{".", "A", "B", "C", "D"}
	fmt.Println("#############")
	fmt.Println("#" + letters[pos[0]] + letters[pos[1]] + letters[pos[2]] + letters[pos[3]] + letters[pos[4]] + letters[pos[5]] + letters[pos[6]] + letters[pos[7]] + letters[pos[8]] + letters[pos[9]] + letters[pos[10]] + "#")
	fmt.Println("###" + letters[pos[11]] + "#" + letters[pos[15]] + "#" + letters[pos[19]] + "#" + letters[pos[23]] + "###")
	fmt.Println("  #" + letters[pos[12]] + "#" + letters[pos[16]] + "#" + letters[pos[20]] + "#" + letters[pos[24]] + "#")
	fmt.Println("  #" + letters[pos[13]] + "#" + letters[pos[17]] + "#" + letters[pos[21]] + "#" + letters[pos[25]] + "#")
	fmt.Println("  #" + letters[pos[14]] + "#" + letters[pos[18]] + "#" + letters[pos[22]] + "#" + letters[pos[26]] + "#")
	fmt.Println("  #########")
	fmt.Println()
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func (pos *Pos) get_moves(part int) []Move {
	/*
		00 01 02 03 04 05 06 07 08 09 10
		      11    15    19    23
		      12    16    20    24
		      13    17    21    25
		      14    18    22    26
	*/

	moves := []Move{}

	// moves from hallway to rooms
outer:
	for i := 0; i <= 10; i++ {
		if pos[i] == 0 {
			continue
		}
		piece := pos[i]
		top, room := destinations[piece][0], destinations[piece][1]
		dir := 1
		if i > top {
			dir = -1
		}
		for j := i + dir; j != top; j += dir {
			if pos[j] > 0 {
				continue outer
			}
		}
		depth := 4
		if part == 1 {
			room -= 2
			depth = 2
		}
		for j := room; j >= room-3; j-- {
			if pos[j] == 0 {
				moves = append(moves,
					Move{
						from: i,
						to:   j,
						cost: (abs(i-top) + depth - (room - j)) * costs[piece],
					},
				)
				break
			}
			if pos[j] != piece {
				break
			}
		}
	}

	// moves from rooms to hallway
	for i := 11; i <= 26; i += 4 {
		for j := 0; j < 4; j++ {
			if pos[i+j] == 0 {
				continue
			}
			for k := moves_out[i] - 1; k >= 0; k-- {
				if k == 2 || k == 4 || k == 6 || k == 8 {
					continue
				}
				if pos[k] == 0 {
					moves = append(moves,
						Move{
							from: i + j,
							to:   k,
							cost: (j + 1 + moves_out[i] - k) * costs[pos[i+j]],
						},
					)
				} else {
					break
				}
			}
			for k := moves_out[i] + 1; k <= 10; k++ {
				if k == 2 || k == 4 || k == 6 || k == 8 {
					continue
				}
				if pos[k] == 0 {
					moves = append(moves,
						Move{
							from: i + j,
							to:   k,
							cost: (j + 1 + k - moves_out[i]) * costs[pos[i+j]],
						},
					)
				} else {
					break
				}
			}
			break
		}
	}

	return moves
}

func search(start_pos Pos, part int) int {
	start_state := State{pos: start_pos, cost: 0}
	visited := map[Pos]int{}

	var end_pos Pos
	if part == 1 {
		end_pos = Pos{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 2, 2, 0, 0, 3, 3, 0, 0, 4, 4, 0, 0}
	} else if part == 2 {
		end_pos = Pos{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4}
	}

	queue := make(PriorityQueue, 1)
	queue[0] = &start_state
	heap.Init(&queue)

	for queue.Len() > 0 {
		state := heap.Pop(&queue).(*State)

		if val, seen := visited[state.pos]; seen {
			if val <= state.cost {
				continue
			}
		}
		visited[state.pos] = state.cost

		if state.pos == end_pos {
			return state.cost
		}

		moves := state.pos.get_moves(part)

		for _, move := range moves {
			new_pos := state.pos.copy_pos()
			new_pos[move.from] = 0
			new_pos[move.to] = state.pos[move.from]
			heap.Push(&queue, &State{
				pos:  new_pos,
				cost: state.cost + move.cost,
			})
		}
	}
	return 0
}

func main() {
	start_pos := get_data()
	cost := search(start_pos, 1)
	fmt.Println(cost)

	// move row down
	start_pos[14] = start_pos[12]
	start_pos[18] = start_pos[16]
	start_pos[22] = start_pos[20]
	start_pos[26] = start_pos[24]

	// insert 2 rows
	start_pos[12] = 4
	start_pos[13] = 4
	start_pos[16] = 3
	start_pos[17] = 2
	start_pos[20] = 2
	start_pos[21] = 1
	start_pos[24] = 1
	start_pos[25] = 3

	cost = search(start_pos, 2)
	fmt.Println(cost)
}
