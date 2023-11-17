package main

import (
	"fmt"
)

const NO_RECIPES = 909441

func mak(n int) int {
	r := []int{3, 7}
	e1 := 0
	e2 := 1

	for i := 1; i <= n+9; i++ {
		next := r[e1] + r[e2]
		if next > 9 {
			r = append(r, 1)
			r = append(r, next%10)
		} else {
			r = append(r, next)
		}
		e1 = (e1 + r[e1] + 1) % len(r)
		e2 = (e2 + r[e2] + 1) % len(r)
	}

	out := 0
	for i := n; i <= n+9; i++ {
		out = out*10 + r[i]
	}
	return out
}

func find(n int) int {
	r := []int{3, 7}
	e1 := 0
	e2 := 1

	l := 0
	tmp := n
	for tmp > 0 {
		l++
		tmp /= 10
	}

	comp_end := 1
	if n&1 == 1 {
		comp_end = 2
	}

loop:
	for {
		next := r[e1] + r[e2]
		if next > 9 {
			r = append(r, 1)
			r = append(r, next%10)
		} else {
			r = append(r, next)
		}
		e1 = (e1 + r[e1] + 1) % len(r)
		e2 = (e2 + r[e2] + 1) % len(r)

		comp := n
		i := len(r) - comp_end
		for comp > 0 {
			if comp%10 != r[i] {
				continue loop
			}
			comp /= 10
			i--
		}
		return len(r) - l - 1
	}

	return -1
}

func main() {
	fmt.Println(mak(NO_RECIPES))
	fmt.Println(find(NO_RECIPES))
}
