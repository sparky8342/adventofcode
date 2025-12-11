package day10

import (
	"fmt"
	"github.com/aclements/go-z3/z3"
	"loader"
	"strconv"
	"strings"
)

type Machine struct {
	goal    int
	buttons []int
	joltage []int
}

type Entry struct {
	state   int
	presses int
}

func parse_data(data []string) []Machine {
	machines := []Machine{}

	for _, line := range data {
		parts := strings.Split(line, " ")
		l := len(parts)

		str := parts[0][1 : len(parts[0])-1]
		goal := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '#' {
				goal = goal | (1 << i)
			}
		}

		buttons := []int{}
		for i := 1; i < l-1; i++ {
			str := parts[i][1 : len(parts[i])-1]
			num_strs := strings.Split(str, ",")
			button := 0
			for _, n_str := range num_strs {
				button = button | (1 << int(n_str[0]-'0'))
			}
			buttons = append(buttons, button)
		}

		joltage_str := parts[l-1][1 : len(parts[l-1])-1]
		num_strs := strings.Split(joltage_str, ",")
		joltage := make([]int, len(num_strs))
		for i, n_str := range num_strs {
			n, err := strconv.Atoi(n_str)
			if err != nil {
				panic(err)
			}
			joltage[i] = n
		}

		machines = append(machines, Machine{
			goal:    goal,
			buttons: buttons,
			joltage: joltage,
		})

	}

	return machines
}

func bfs(machine Machine) int {
	queue := []Entry{Entry{}}
	visited := map[int]struct{}{}
	visited[0] = struct{}{}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		if entry.state == machine.goal {
			return entry.presses
		}

		for _, button := range machine.buttons {
			new_state := entry.state ^ button
			if _, ok := visited[new_state]; !ok {
				queue = append(queue, Entry{
					state:   new_state,
					presses: entry.presses + 1,
				})
				visited[new_state] = struct{}{}
			}
		}
	}

	return -1
}

func joltage_calc(machine Machine) int {
	ctx := z3.NewContext(nil)
	s := z3.NewSolver(ctx)
	zero := ctx.FromInt(int64(0), ctx.IntSort()).(z3.Int)

	counts := []z3.Int(nil)
	for i := range machine.buttons {
		button := ctx.IntConst(fmt.Sprintf("button%d", i))
		s.Assert(button.GE(zero))
		counts = append(counts, button)
	}

	for i, joltage := range machine.joltage {
		presses := []z3.Int(nil)
		for j, button := range machine.buttons {
			if button&(1<<i) != 0 {
				presses = append(presses, counts[j])
			}
		}
		s.Assert(zero.Add(presses...).Eq(ctx.FromInt(int64(joltage), ctx.IntSort()).(z3.Int)))
	}

	total := ctx.IntConst("total")
	s.Assert(total.Eq(zero.Add(counts...)))

	var min_presses int64 = 0
	for {
		if sat, err := s.Check(); !sat || err != nil {
			return int(min_presses)
		}

		min_presses, _, _ = s.Model().Eval(total, true).(z3.Int).AsInt64()

		s.Assert(total.LT(ctx.FromInt(min_presses, ctx.IntSort()).(z3.Int)))
	}
}

func presses_needed(machines []Machine) int {
	presses := 0
	for _, machine := range machines {
		presses += bfs(machine)
	}
	return presses
}

func joltage_presses_needed(machines []Machine) int {
	presses := 0
	for _, machine := range machines {
		presses += joltage_calc(machine)
	}
	return presses
}

func Run() {
	loader.Day = 10
	data := loader.GetStrings()
	machines := parse_data(data)
	part1 := presses_needed(machines)
	part2 := joltage_presses_needed(machines)

	fmt.Printf("%d %d\n", part1, part2)
}
