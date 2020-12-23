#!/usr/bin/python3
from collections import deque

def play_game(cups, turns, part):
	cup_pos = 0
	cup = cups[cup_pos]

	length = len(cups)

	removed = [0, 0, 0]

	for turn in range(turns):
		remove_pos = cup_pos + 1
		for i in range(3):
			if remove_pos == length:
				remove_pos = 0
			removed[i] = cups[remove_pos]
			remove_pos += 1

		for n in removed:
			cups.remove(n)

		destination_cup = cup - 1
		while True:
			if destination_cup == 0:
				destination_cup = length
			if destination_cup not in removed:
				break
			destination_cup -= 1

		dest_pos = cups.index(destination_cup)
		dest_pos += 1
		for n in removed:
			cups.insert(dest_pos, n)
			dest_pos += 1

		cup_pos = (cups.index(cup) + 1) % length
		cup = cups[cup_pos]

	cups.rotate(cups.index(1) * -1)
	if part == 1:
		return "".join(str(x) for x in list(cups)[1:])
	elif part == 2:
		print(cups[0], cups[1], cups[2], cups[3])
		return cups[1] * cups[2]

# part 1
nums = [int(x) for x in open('input_test.txt').read().strip()]
cups = deque(nums)
print(play_game(cups, 100, 1))

# part 2
cups = deque(nums)
for i in range (10, 1000000 + 1):
	cups.append(i)

print(play_game(cups, 10000000, 2))

# 110 too low

