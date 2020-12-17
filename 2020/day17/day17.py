#!/usr/bin/python3

def get_neighbours(cubes, cube):
	neighbours = set()
	neighbours.add(())

	dimension = len(cube)

	for i in range(0, dimension):
		new_neighbours = set()
		for space in neighbours:
			new_neighbours.add(space + (cube[i],))
			new_neighbours.add(space + (cube[i] + 1,))
			new_neighbours.add(space + (cube[i] - 1,))
		neighbours = new_neighbours

	neighbours.remove(cube)

	neighbour_count = 0
	spaces = set()
	for neighbour in neighbours:
		if neighbour in cubes:
			neighbour_count += 1
		else:
			spaces.add(neighbour)

	return neighbour_count, spaces

def generation(cubes):
	new_cubes = set()
	all_spaces = set()

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
	cubes3d = generation(cubes3d)
	cubes4d = generation(cubes4d)

print(len(cubes3d))
print(len(cubes4d))
