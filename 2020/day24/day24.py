#!/usr/bin/python3
import re

dirs = {
	'se' : (0, -1, 1),
	'sw' : (-1, 0, 1),
	'ne' : (1, 0 ,-1),
	'nw' : (0, 1, -1),
	'e'  : (1, -1, 0),
	'w'  : (-1, 1, 0)
}
changes = dirs.values()

lines = open('input.txt').read().splitlines()

def get_neighbours(tile):
	neighbours = 0
	spaces = set()
	for d in changes:
		space = (tile[0] + d[0], tile[1] + d[1], tile[2] + d[2])
		if space in tiles:
			neighbours += 1
		else:
			spaces.add(space)

	return neighbours, spaces

def generation(tiles):
	new_tiles = set()
	all_spaces = set()

	for tile in tiles:
		neighbours, spaces = get_neighbours(tile)
		if neighbours == 1 or neighbours == 2:
			new_tiles.add(tile)
		all_spaces = all_spaces.union(spaces)

	for space in all_spaces:
		neighbours, _ = get_neighbours(space)
		if neighbours == 2:
			new_tiles.add(space)

	return new_tiles

tiles = set()

# part 1
for line in lines:
	steps = re.findall("(se|sw|ne|nw|e|w)", line)

	x = 0
	y = 0
	z = 0

	for step in steps:
		change = dirs[step]
		x += change[0]
		y += change[1]
		z += change[2]

	place = (x, y, z)
	if place in tiles:
		tiles.remove(place)
	else:
		tiles.add(place)	

print(len(tiles))

# part 2
for _ in range(100):
	tiles = generation(tiles)

print(len(tiles))
