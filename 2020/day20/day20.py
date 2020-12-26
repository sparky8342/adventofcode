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
	variations = []
	for i in range(0, 8):
		squares = rotate_right(squares)
		if i == 4:
			squares = flip(squares)

		# checksums of edges
		top = int("".join(squares[0]).replace('.','0').replace('#','1'), 2)
		bottom = int("".join(squares[END]).replace('.','0').replace('#','1'), 2)
		left = ''
		right = ''
		for i in range(SIZE):
			left += squares[i][0]
			right += squares[i][END]
		left = int(left.replace('.','0').replace('#','1'), 2)
		right = int(right.replace('.','0').replace('#','1'), 2)

		variations.append({'squares' : squares, 'top' : top, 'bottom' : bottom, 'left' : left, 'right' : right})

	return variations

def side_edge_match(tiles, place1, place2):
	tile1, variation1 = place1
	tile2, variation2 = place2
	tile1 = tiles[tile1]
	tile2 = tiles[tile2]
	return tile1['variations'][variation1]['right'] == tile2['variations'][variation2]['left']

def top_bottom_edge_match(tiles, place1, place2):
	tile1, variation1 = place1
	tile2, variation2 = place2
	tile1 = tiles[tile1]
	tile2 = tiles[tile2]
	return tile1['variations'][variation1]['bottom'] == tile2['variations'][variation2]['top']

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

def search(tiles, square, combos, x, y, answer_squares):
	if not is_valid(tiles, square):
		return

	if x == len(square):
		x = 0
		y += 1

	if y == len(square):
		answer_squares.append(deepcopy(square))
		return

	for c in combos:
		square[y][x] = c
		search(tiles, square, combos, x + 1, y, answer_squares)
		square[y][x] = (-1, -1)

def backtrack(tiles):
	square_size = int(math.sqrt(len(tiles)))
	square = []
	for _ in range(square_size):
		square.append([(-1,-1) for _ in range(square_size)])

	combos = []
	for i in range(len(tiles)):
		for j in range(8):
			combos.append((i, j))

	answer_squares = []
	search(tiles, square, combos, 0, 0, answer_squares)
	return answer_squares

def get_sea_monster():
	return ['                  # ','#    ##    ##    ###',' #  #  #  #  #  #   ']

def remove_sea_monster_at(grid, x, y):
	grid_size = len(grid)
	sea_monster = get_sea_monster()
	sea_monster_h = len(sea_monster)
	sea_monster_w = len(sea_monster[0])

	for sy in range(sea_monster_h):
		if y + sy == grid_size:
			return
		for sx in range(sea_monster_w):
			if x + sx == grid_size:
				return
			if sea_monster[sy][sx] == '#':
				if grid[y + sy][x + sx] != '#':
					return

	for sy in range(sea_monster_h):
		for sx in range(sea_monster_w):
			if sea_monster[sy][sx] == '#':
				grid[y + sy][x + sx] = '.'

def count_rough(grid):
	size = len(grid)
	for y in range(size):
		for x in range(size):
			remove_sea_monster_at(grid, x, y)

	rough = 0
	for y in range(size):
		for x in range(size):
			if grid[y][x] == '#':
				rough += 1

	return rough

def make_image(tiles, square):
	size = len(square)
	grid_size = size * (SIZE - 2)
	grid = []
	for _ in range(grid_size):
		grid.append(['.' for _ in range(grid_size)])

	for y in range(size):
		for x in range(size):
			tile_no, variation = square[y][x]
			squares = tiles[tile_no]['variations'][variation]['squares']
			for sq_y in range(1, SIZE-1):
				for sq_x in range(1, SIZE-1):
					grid[y * (SIZE - 2) + sq_y - 1][x * (SIZE - 2) + sq_x - 1] = squares[sq_y][sq_x]

	return grid

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

# part 1
tiles = get_input()
squares = backtrack(tiles)
ans = calc_answer(tiles, squares[0])
print(ans)

# part 2
rough_amounts = []
for square in squares:
	grid = make_image(tiles, square)
	rough = count_rough(grid)
	rough_amounts.append(rough)

print(min(rough_amounts))
