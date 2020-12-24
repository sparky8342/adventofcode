#!/usr/bin/python3
import re

lines = open('input.txt').read().splitlines()

tiles = set()

for line in lines:
	dirs = re.findall("(se|sw|ne|nw|e|w)", line)

	x = 0
	y = 0
	z = 0

	for d in dirs:
		if d == 'se':
			z += 1
			y -= 1
		elif d == 'sw':
			x -= 1
			z += 1
		elif d == 'ne':
			x += 1
			z -= 1
		elif d == 'nw':
			z -= 1
			y += 1
		elif d == 'e':
			x += 1
			y -= 1
		elif d == 'w':
			x -= 1
			y += 1

	place = (x, y, z)
	if place in tiles:
		tiles.remove(place)
	else:
		tiles.add(place)	

print(len(tiles))
