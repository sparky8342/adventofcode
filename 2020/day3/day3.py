#!/usr/bin/python3

def ski(grid, dx, dy):
	width = len(grid[0])
	x = 0
	y = 0
	trees = 0
	while y < len(grid):
		if grid[y][x % width] == '#':
			trees += 1
		y += dy
		x += dx
	return trees

grid = [list(line) for line in (open('input.txt').read().splitlines())]

print(ski(grid, 3, 1))

answer = 1
slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
for dx, dy in slopes:
	answer *= ski(grid, dx, dy)
print(answer)
