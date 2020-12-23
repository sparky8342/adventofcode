#!/usr/bin/python3
from collections import deque

cups = deque(int(x) for x in open('input.txt').read().strip())

cup_pos = 0
cup = cups[cup_pos]

for _ in range(100):
	cups.rotate((cup_pos + 1) * -1)
	removed = []
	for _ in range(3):
		removed.append(cups.popleft())

	destination_cup = cup - 1
	while True:
		if destination_cup == 0:
			destination_cup = 9
		if destination_cup not in removed:
			break
		destination_cup -= 1

	dest_pos = cups.index(destination_cup)
	dest_pos += 1
	for n in removed:
		cups.insert(dest_pos, n)
		dest_pos += 1

	cup_pos = (cups.index(cup) + 1) % 9
	cup = cups[cup_pos]

cups.rotate(cups.index(1) * -1)
print("".join(str(x) for x in list(cups)[1:]))
