#!/usr/bin/python3
import copy

class Grid:
	def __init__(self, squares, rule_type):
		self.squares = squares
		self.rule_type = rule_type
		if rule_type == 1:
			self.max_neighbours = 4
			self.count_seen = False
		elif rule_type == 2:
			self.max_neighbours = 5
			self.count_seen = True
		self.height = len(squares)
		self.width = len(squares[0])
		self.stable = False
		self.occupied = 0

	def in_bounds(self, x, y):
		if x < 0 or x >= self.width or y < 0 or y >= self.height:
			return False
		return True

	def count_neighbours(self, x, y):
		neighbours = 0
		for dy in range(-1,2):
			for dx in range(-1,2):
				if dx == 0 and dy == 0:
					continue
				newx = x + dx
				newy = y + dy
				while self.in_bounds(newx, newy):
					if self.squares[newy][newx] == '#':
						neighbours += 1
						break
					elif self.squares[newy][newx] == 'L':
						break
					if self.count_seen:
						newx += dx
						newy += dy
					else:
						break
		return neighbours

	def generation(self):
		new_squares = copy.deepcopy(self.squares)
		self.stable = True
		for y in range(self.height):
			for x in range(self.width):
				if self.squares[y][x] == '.':
					continue
				neighbours = self.count_neighbours(x, y)
				if self.squares[y][x] == 'L' and neighbours == 0:
					new_squares[y][x] = '#'
					self.stable = False
				elif self.squares[y][x] == '#' and neighbours >= self.max_neighbours:
					new_squares[y][x] = 'L'
					self.stable = False

		self.squares = new_squares
		if self.stable:
			self.update_occupied()

	def run(self):
		while self.stable == False:
			self.generation()

	def update_occupied(self):
		self.occupied = 0
		for y in range(self.height):
			for x in range(self.width):
				if self.squares[y][x] == '#':
					self.occupied += 1

	def print_grid(self):
		for y in range(self.height):
			for x in range(self.width):
				print(self.squares[y][x], end="")
			print()
		print()

squares = [list(line) for line in (open('input.txt').read().splitlines())]
g = Grid(squares, 1)
g.run()
print(g.occupied)

g = Grid(squares, 2)
g.run()
print(g.occupied)
