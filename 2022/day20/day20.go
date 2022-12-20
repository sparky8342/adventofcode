package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const KEY = 811589153

type Node struct {
	val   int
	left  *Node
	right *Node
}

type List struct {
	start *Node
	order []*Node
	zero  *Node
}

func load_data(filename string) []int {
	data, _ := ioutil.ReadFile(filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	nums := []int{}
	for _, line := range strings.Split(string(data), "\n") {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return nums
}

func create_list(nums []int) List {
	list := List{}

	list.start = &Node{val: nums[0]}
	list.order = append(list.order, list.start)
	previous := list.start

	var node *Node
	for i := 1; i < len(nums); i++ {
		node = &Node{val: nums[i], left: previous}
		previous.right = node
		previous = node
		list.order = append(list.order, node)
		if nums[i] == 0 {
			list.zero = node
		}
	}

	node.right = list.start
	list.start.left = node

	return list
}

func (list *List) move_nums() {
	for _, node := range list.order {
		if node.val == 0 {
			continue
		}

		// remove node
		node.left.right = node.right
		node.right.left = node.left

		pos := node
		if node.val > 0 {
			for i := 0; i < node.val%(len(list.order)-1); i++ {
				pos = pos.right
			}
		} else {
			for i := 0; i < (node.val*-1)%(len(list.order)-1)+1; i++ {
				pos = pos.left
			}
		}

		// add in new place
		node.left = pos
		node.right = pos.right
		pos.right.left = node
		pos.right = node
	}
}

func (list *List) print_list() {
	node := list.start
	fmt.Println(node)
	node = node.right
	for node != list.start {
		fmt.Println(node)
		node = node.right
	}
	fmt.Println()
}

func (list *List) find_sum() int {
	sum := 0
	pos := list.zero
	for i := 1; i <= 3000; i++ {
		pos = pos.right
		if i%1000 == 0 {
			sum += pos.val
		}
	}
	return sum
}

func (list *List) apply_key() {
	start := list.start
	node := start

	node.val *= KEY
	node = node.right
	for node != start {
		node.val *= KEY
		node = node.right
	}
}

func main() {
	nums := load_data("input.txt")

	list := create_list(nums)
	list.move_nums()
	fmt.Println(list.find_sum())

	list = create_list(nums)
	list.apply_key()
	for i := 0; i < 10; i++ {
		list.move_nums()
	}
	fmt.Println(list.find_sum())
}
