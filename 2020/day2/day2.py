#!/usr/bin/python3

with open('input.txt') as f:
        lines = f.read().splitlines()

part1 = 0
part2 = 0
for line in lines:
	parts = line.split(" ")
	mn, mx = parts[0].split("-")
	mn = int(mn)
	mx = int(mx)
	letter = parts[1][0]
	count = 0
	for char in parts[2]:
		if char == letter:
			count += 1

	if mn <= count and count <= mx:
		part1 += 1

	count = 0
	for pos in (mn - 1, mx - 1):
		if parts[2][pos] == letter:
			count += 1

	if count == 1:
		part2 += 1
	
print(part1)
print(part2)
