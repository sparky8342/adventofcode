#!/usr/bin/python3
import copy

class Grid:
	def __init__(self, squares):
		self.squares = squares
		self.height = len(squares)
		self.width = len(squares[0])
		self.stable = False
		self.update_occupied()

	def count_neighbours(self, x, y):
		neighbours = 0
		for dy in range(-1,2):
			for dx in range(-1,2):
				if dx == 0 and dy == 0:
					continue
				newx = x + dx
				newy = y + dy
				if newx < 0 or newx == self.width or newy < 0 or newy == self.height:
					continue
				if self.squares[newy][newx] == '#':
					neighbours += 1
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
				elif self.squares[y][x] == '#' and neighbours >= 4:
					new_squares[y][x] = 'L'
					self.stable = False

		self.squares = new_squares
		self.update_occupied()

	def count_seen_neighbours(self, x, y):
		neighbours = 0
		for dy in range(-1,2):
			for dx in range(-1,2):
				if dx == 0 and dy == 0:
					continue
				newx = x + dx
				newy = y + dy
				while 1:
					if newx < 0 or newx == self.width or newy < 0 or newy == self.height:
						break
					if self.squares[newy][newx] == '#':
						neighbours += 1
						break
					elif self.squares[newy][newx] == 'L':
						break
					newx += dx
					newy += dy
		return neighbours

	def generation2(self):
		new_squares = copy.deepcopy(self.squares)
		self.stable = True
		for y in range(self.height):
			for x in range(self.width):
				if self.squares[y][x] == '.':
					continue
				neighbours = self.count_seen_neighbours(x, y)
				if self.squares[y][x] == 'L' and neighbours == 0:
					new_squares[y][x] = '#'
					self.stable = False
				elif self.squares[y][x] == '#' and neighbours >= 5:
					new_squares[y][x] = 'L'
					self.stable = False

		self.squares = new_squares
		self.update_occupied()

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
g = Grid(copy.deepcopy(squares))

while g.stable == False:
	g.generation()
print(g.occupied)

g = Grid(squares)
while g.stable == False:
	g.generation2()
print(g.occupied)
