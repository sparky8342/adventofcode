#!/usr/bin/python3
import re

with open('input.txt') as f:
        lines = f.read().splitlines()

part1 = 0
part2 = 0
for line in lines:
	match = re.search("^(\d+)-(\d+) (\w): (\w+)$", line)
	mn = int(match.group(1))
	mx = int(match.group(2))
	letter = match.group(3)
	password = match.group(4)

	count = 0
	for char in password:
		if char == letter:
			count += 1

	if mn <= count and count <= mx:
		part1 += 1

	count = 0
	for pos in (mn - 1, mx - 1):
		if password[pos] == letter:
			count += 1

	if count == 1:
		part2 += 1

print(part1)
print(part2)
