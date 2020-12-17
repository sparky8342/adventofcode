#!/usr/bin/python3

def get_neighbours_3d(cubes, cube):
	neighbours = 0
	spaces = set()
	for dx in (-1, 0, 1):
		for dy in (-1, 0, 1):
			for dz in (-1, 0, 1):
				if dx == 0 and dy == 0 and dz == 0:
					continue
				position = (cube[0] + dx, cube[1] + dy, cube[2] + dz)
				if position in cubes:
					neighbours += 1
				else:
					spaces.add(position)

	return neighbours, spaces

def get_neighbours_4d(cubes, cube):
	neighbours = 0
	spaces = set()
	for dx in (-1, 0, 1):
		for dy in (-1, 0, 1):
			for dz in (-1, 0, 1):
				for dw in (-1, 0, 1):
					if dx == 0 and dy == 0 and dz == 0 and dw == 0:
						continue
					position = (cube[0] + dx, cube[1] + dy, cube[2] + dz, cube[3] + dw)
					if position in cubes:
						neighbours += 1
					else:
						spaces.add(position)

	return neighbours, spaces

def generation(cubes, dimension):
	new_cubes = set()
	all_spaces = set()
	get_neighbours = None
	if dimension == 3:
		get_neighbours = get_neighbours_3d
	elif dimension == 4:
		get_neighbours = get_neighbours_4d
	
	for cube in cubes:
		neighbours, spaces = get_neighbours(cubes, cube)
		if neighbours == 2 or neighbours == 3:
			new_cubes.add(cube)
		all_spaces = all_spaces.union(spaces)

	for space in all_spaces:
		neighbours, _ = get_neighbours(cubes, space)
		if neighbours == 3:
			new_cubes.add(space)

	return new_cubes

grid = [list(line) for line in (open('input.txt').read().splitlines())]

cubes3d = set()
cubes4d = set()

for y in range(len(grid)):
	for x in range(len(grid[0])):
		if grid[y][x] == '#':
			cubes3d.add((x, y, 0))
			cubes4d.add((x, y, 0, 0))


for _ in range(6):
	cubes3d = generation(cubes3d, 3)
	cubes4d = generation(cubes4d, 4)

print(len(cubes3d))
print(len(cubes4d))
