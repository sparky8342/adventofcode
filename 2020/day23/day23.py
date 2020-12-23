#!/usr/bin/python3

class CircularList:
	# circular list - each element is [value, prev_pos, next_pos]
	# lookup hash used to avoid scanning through for values

	def __init__(self):
		self.lst = []
		self.lookup = {}
		self.deleted = []

	def __str__(self):
		st = ''
		for item in self.lst:
			st += ",".join(str(x) for x in item) + ' '
		return st

	def append_to_list(self, value):
		if len(self.lst) == 0:
			self.lst.append([value, 0, 0])
			self.lookup[value] = 0
		else:
			new = [value, len(self.lst) - 1, 0]
			self.lst[-1][2] = len(self.lst)
			self.lst.append(new)
			self.lst[0][1] = len(self.lst) - 1
			self.lookup[value] = len(self.lst) - 1

	def get(self, position):
		return self.lst[position][0]

	def del_three_after(self, position):
		removed = []
		entry = self.lst[position]
		nxt = entry[2]
		for _ in range(3):
			removed.append(self.lst[nxt][0])
			self.lst[nxt][0] = -1
			self.deleted.append(nxt)
			nxt = self.lst[nxt][2]

		self.lst[position][2] = nxt
		self.lst[nxt][1] = position
		return removed

	def add_after_pos(self, pos, value):
		# assuming we always have deleted positions,
		# which is true for this puzzle
		free_pos = self.deleted.pop()
		self.lst[free_pos] = [value, pos, self.lst[pos][2]]
		self.lst[pos][2] = free_pos
		self.lookup[value] = free_pos
		return free_pos

	def add_after_value(self, value, to_add):
		pos = self.lookup[value]
		for val in to_add:
			pos = self.add_after_pos(pos, val)

	def next(self, pos):
		nxt = self.lst[pos][2]
		return self.lst[nxt][0], nxt

	def elements_after_value(self, value, amount):
		pos = self.lookup[value]
		out = []
		for _ in range(amount):
			pos = self.lst[pos][2]
			out.append(self.lst[pos][0])
		return out

def play_game(cups, length, turns, part):
	cup_pos = 0
	cup = cups.get(cup_pos)

	for _ in range(turns):
		removed = cups.del_three_after(cup_pos)

		destination_cup = cup - 1
		while True:
			if destination_cup == 0:
				destination_cup = length
			if destination_cup not in removed:
				break
			destination_cup -= 1

		cups.add_after_value(destination_cup, removed)
		cup, cup_pos = cups.next(cup_pos)

	if part == 1:
		ans = "".join(str(x) for x in cups.elements_after_value(1, length - 1))
		return(ans)
	elif part == 2:
		values = cups.elements_after_value(1, 2)
		return values[0] * values[1]


nums = [int(x) for x in open('input.txt').read().strip()]

# part 1
cups = CircularList()
for n in nums:
	cups.append_to_list(n)

print(play_game(cups, 9, 100, 1))

# part 2
cups = CircularList()
for n in nums:
	cups.append_to_list(n)
for i in range (10, 1000001):
	cups.append_to_list(i)

print(play_game(cups, 1000000, 10000000, 2))
