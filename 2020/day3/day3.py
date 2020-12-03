#!/usr/bin/python3

def ski(dx, dy):
	x = 0
	y = 0
	trees = 0
	while y < len(grid):
		if grid[y][x] == '#':
			trees += 1
		y += dy
		x += dx
		if x >= width:
			x = x - width
	return trees

with open('input.txt') as f:
        lines = f.read().splitlines()

grid = []
for line in lines:
	grid.append(list(line))

width = len(lines[0])
print(ski(3,1))

answer = 1
slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
for dx, dy in slopes:
	answer *= ski(dx, dy)

print(answer)
