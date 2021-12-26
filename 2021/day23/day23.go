package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strings"
)

type Pos [19]int

type State struct {
	pos    Pos
	parent *State
	cost   int
	index  int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

/*
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########

00 01 02 03 04 05 06 07 08 09 10
      11    13    15    17
      12    14    16    18

A = 1
B = 2
C = 3
D = 4
*/

var moves [][][]int
var nums map[byte]int
var spaces map[int][]int

func has_duplicates(nums []int) bool {
	m := map[int]struct{}{}
	for _, n := range nums {
		if _, seen := m[n]; seen {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}

func valid_move(nums []int) bool {
	last := nums[len(nums)-1]
	if last == 2 || last == 4 || last == 6 || last == 8 {
		return false
	}

	changes := 0
	corridor := false
	if nums[0] <= 10 {
		corridor = true
	}
	start := corridor
	for i := 1; i < len(nums); i++ {
		c := false
		if nums[i] <= 10 {
			c = true
		}
		if c != corridor {
			corridor = c
			changes++
		}
	}
	if changes == 1 || (changes == 2 && start == false) {
		return true
	} else {
		return false
	}
}

func init() {
	spaces = map[int][]int{
		0:  {1},
		1:  {0, 2},
		2:  {1, 3, 11},
		3:  {2, 4},
		4:  {3, 5, 13},
		5:  {4, 6},
		6:  {5, 7, 15},
		7:  {6, 8},
		8:  {7, 9, 17},
		9:  {8, 10},
		10: {9},
		11: {2, 12},
		12: {11},
		13: {4, 14},
		14: {13},
		15: {6, 16},
		16: {15},
		17: {8, 18},
		18: {17},
	}

	moves = [][][]int{}
	for i := 0; i < 19; i++ {
		moves = append(moves, [][]int{})
	}

	for i := 0; i < 19; i++ {
		queue := [][]int{{i}}

		for len(queue) > 0 {
			state := queue[0]
			queue = queue[1:]

			if valid_move(state) {
				mv := make([]int, len(state))
				copy(mv, state)
				move_start := mv[0]
				mv = mv[1:]
				moves[move_start] = append(moves[move_start], mv)
			}

			last := state[len(state)-1]
			for _, space := range spaces[last] {
				new_state := make([]int, len(state))
				copy(new_state, state)
				new_state = append(new_state, space)
				if !has_duplicates(new_state) {
					queue = append(queue, new_state)
				}
			}
		}
	}

	nums = map[byte]int{'A': 1, 'B': 2, 'C': 3, 'D': 4}
}

func get_data() Pos {
	data, _ := ioutil.ReadFile("input.txt")
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	pos := Pos{}
	pos[11] = nums[lines[2][3]]
	pos[12] = nums[lines[3][3]]
	pos[13] = nums[lines[2][5]]
	pos[14] = nums[lines[3][5]]
	pos[15] = nums[lines[2][7]]
	pos[16] = nums[lines[3][7]]
	pos[17] = nums[lines[2][9]]
	pos[18] = nums[lines[3][9]]
	return pos
}

func (p *Pos) copy_pos() Pos {
	new_pos := Pos{}
	for i := 0; i < 19; i++ {
		new_pos[i] = p[i]
	}
	return new_pos
}

func (pos *Pos) print_pos() {
	letters := []string{".", "A", "B", "C", "D"}
	fmt.Println("#############")
	fmt.Println("#" + letters[pos[0]] + letters[pos[1]] + letters[pos[2]] + letters[pos[3]] + letters[pos[4]] + letters[pos[5]] + letters[pos[6]] + letters[pos[7]] + letters[pos[8]] + letters[pos[9]] + letters[pos[10]] + "#")
	fmt.Println("###" + letters[pos[11]] + "#" + letters[pos[13]] + "#" + letters[pos[15]] + "#" + letters[pos[17]] + "###")
	fmt.Println("  #" + letters[pos[12]] + "#" + letters[pos[14]] + "#" + letters[pos[16]] + "#" + letters[pos[18]] + "#")
	fmt.Println("  #########")
	fmt.Println()
}

func search(start_pos Pos) int {
	start_state := State{pos: start_pos, cost: 0}
	valid_destinations := map[int]int{11: 1, 12: 1, 13: 2, 14: 2, 15: 3, 16: 3, 17: 4, 18: 4}
	end_pos := Pos{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4}
	costs := [5]int{0, 1, 10, 100, 1000}
	visited := map[Pos]struct{}{}

	queue := make(PriorityQueue, 1)
	queue[0] = &start_state
	heap.Init(&queue)

	for queue.Len() > 0 {
		state := heap.Pop(&queue).(*State)

		if _, seen := visited[state.pos]; seen {
			continue
		}
		visited[state.pos] = struct{}{}

		if state.pos == end_pos {
			return state.cost
		}

		for i := 0; i < 19; i++ {
			if state.pos[i] > 0 {
				// don't move if in correct position
				if i == 12 && state.pos[12] == 1 {
					continue
				}
				if i == 11 && state.pos[11] == 1 && state.pos[12] == 1 {
					continue
				}
				if i == 14 && state.pos[14] == 2 {
					continue
				}
				if i == 13 && state.pos[13] == 2 && state.pos[14] == 2 {
					continue
				}
				if i == 16 && state.pos[16] == 3 {
					continue
				}
				if i == 15 && state.pos[15] == 3 && state.pos[16] == 3 {
					continue
				}
				if i == 18 && state.pos[18] == 4 {
					continue
				}
				if i == 17 && state.pos[17] == 4 && state.pos[18] == 4 {
					continue
				}

				for _, move := range moves[i] {
					move_ok := true
					for _, step := range move {
						if state.pos[step] > 0 {
							move_ok = false
							break
						}
					}
					if move_ok {
						dest := move[len(move)-1]
						if dest >= 11 {
							val, _ := valid_destinations[dest]
							if val != state.pos[i] {
								continue
							}
						}
						new_pos := state.pos.copy_pos()
						new_pos[i] = 0
						new_pos[dest] = state.pos[i]
						heap.Push(&queue, &State{
							pos:    new_pos,
							cost:   state.cost + len(move)*costs[state.pos[i]],
							parent: state,
						})
					}
				}
			}
		}
	}
	return 0
}

func main() {
	start_pos := get_data()
	cost := search(start_pos)
	fmt.Println(cost)
}
