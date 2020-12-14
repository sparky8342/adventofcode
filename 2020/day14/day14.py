#!/usr/bin/python3
import re

lines = [line for line in (open('input.txt').read().splitlines())]

# part1
mem = {}
mask1 = 0
mask2 = 0
for line in lines:
	match = re.search("mask = ([01X]+)", line)
	if match:
		mask1 = int(re.sub("X", "0", match.group(1)), 2)
		mask2 = int(re.sub("X", "1", match.group(1)), 2)
		continue

	match = re.search("mem\[(\d+)\] = (\d+)", line)
	if match:
		mem[int(match.group(1))] = (int(match.group(2)) | mask1) & mask2

print(sum(mem.values())) 

# part2
def get_addresses(num, mask):
	for i, char in enumerate(mask):
		if char == '1' or char == 'X':
			num[i] = char

	return get_combos(num, [])

def get_combos(num, c):
	for i, char in enumerate(num):
		if char == 'X':
			for digit in ('0', '1'):
				cp = num.copy()
				cp[i] = digit
				c = get_combos(cp, c)
			return c

	c.append(int("".join(num), 2))
	return c

mem = {}
mask = []
for line in lines:
	match = re.search("mask = ([01X]+)", line)
	if match:
		mask = list(match.group(1))
		continue

	match = re.search("mem\[(\d+)\] = (\d+)", line)
	if match:
		address = match.group(1)
		value = int(match.group(2))
		address = list(bin(int(address))[2:].zfill(36))

		addresses = get_addresses(address, mask)
		for add in addresses:
			mem[add] = value

print(sum(mem.values()))
