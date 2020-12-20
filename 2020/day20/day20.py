#!/usr/bin/python3
from copy import deepcopy
import re
import math

SIZE = 10
END = SIZE - 1

def rotate_right(squares):
	new_squares = deepcopy(squares)

	for y in range(SIZE):
		for x in range(SIZE):
			new_squares[y][x] = squares[-(x+1)][y][:]

	return new_squares

def flip(squares):
	new_squares = deepcopy(squares)

	for y in range(SIZE):
		for x in range(SIZE):
			new_squares[END - y][x] = squares[y][x]

	return new_squares

def get_variations(squares):
	variations = [squares]
	for i in range(0, 7):
		squares = rotate_right(squares)
		if i == 3:
			squares = flip(squares)
		variations.append(squares)

	return variations

def side_edge_match(tiles, place1, place2):
	tile1, variation1 = place1
	tile2, variation2 = place2
	tile1 = tiles[tile1]
	tile2 = tiles[tile2]

	tile1_edge = ''
	tile2_edge = ''
	for i in range(SIZE):
		tile1_edge += tile1['variations'][variation1][i][END]
		tile2_edge += tile2['variations'][variation2][i][0]

	return tile1_edge == tile2_edge		

def top_bottom_edge_match(tiles, place1, place2):
	tile1, variation1 = place1
	tile2, variation2 = place2
	tile1 = tiles[tile1]
	tile2 = tiles[tile2]

	tile1_edge = "".join(tile1['variations'][variation1][END])
	tile2_edge = "".join(tile2['variations'][variation2][0])

	return tile1_edge == tile2_edge

def is_valid(tiles, square):
	size = len(square)
	tile_nums = set()
	for y in range(size):
		for x in range(size):
			tile_no = square[y][x][0]
			if tile_no == -1:
				continue
			if tile_no in tile_nums:
				return False
			tile_nums.add(tile_no)

	for y in range(size):
		for x in range(size - 1):
			if square[y][x][0] != -1 and square[y][x+1][0] != -1:
				if not side_edge_match(tiles, square[y][x], square[y][x + 1]):
					return False

	for x in range(size):
		for y in range(size - 1):
			if square[y][x][0] != -1 and square[y+1][x][0] != -1:
				if not top_bottom_edge_match(tiles, square[y][x], square[y+1][x]):
					return False

	return True

def calc_answer(tiles, square):
	end = len(square) - 1
	return tiles[square[0][0][0]]['id'] * tiles[square[0][end][0]]['id'] * tiles[square[end][0][0]]['id'] * tiles[square[end][end][0]]['id']


def search(tiles, square, combos, x, y):
	if not is_valid(tiles, square):
		return -1

	if x == len(square):
		x = 0
		y += 1

	if y == len(square):
		return calc_answer(tiles, square)

	for c in combos:
		square[y][x] = c
		ans = search(tiles, square, combos, x + 1, y)
		if ans != -1:
			return ans
		square[y][x] = (-1, -1)

	return -1

def backtrack(tiles):
	square_size = int(math.sqrt(len(tiles)))
	square = []
	for _ in range(square_size):
		square.append([(-1,-1) for _ in range(square_size)])

	combos = []
	for i in range(len(tiles)):
		for j in range(8):
			combos.append((i, j))

	return search(tiles, square, combos, 0, 0)

def get_input():
	tiles = []
	data = open('input.txt').read().strip()
	data = data.split("\n\n")
	for tile_data in data:
		lines = tile_data.split("\n")
		match = re.match("Tile (\d+):", lines[0])
		tile_id = int(match.group(1))
		squares = [list(l) for l in lines[1:]]
		tile = { 'id' : tile_id }
		tile['variations'] = get_variations(squares)
		tiles.append(tile)
		
	return tiles

tiles = get_input()
print(backtrack(tiles))
