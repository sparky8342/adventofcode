package main

import (
	"fmt"
)

const NUM_ELVES = 3005290

type Elf struct {
	id   int
	next *Elf
}

func init_elves() *Elf {
	head := &Elf{id: 1}

	elf := head

	for i := 2; i <= NUM_ELVES; i++ {
		next := &Elf{id: i}
		elf.next = next
		elf = next
	}

	elf.next = head
	return head
}

func part1() int {
	elf := init_elves()

	for elf != elf.next {
		elf.next = elf.next.next
		elf = elf.next
	}

	return elf.id
}

func part2() int {
	elf := init_elves()

	opposite := elf
	for i := 1; i < NUM_ELVES/2; i++ {
		opposite = opposite.next
	}

	for elf != elf.next {
		opposite.next = opposite.next.next
		elf = elf.next

		opposite.next = opposite.next.next
		opposite = opposite.next
		elf = elf.next
	}

	return elf.id
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
