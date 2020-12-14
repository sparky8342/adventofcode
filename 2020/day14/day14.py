#!/usr/bin/python3
import re

lines = [line for line in (open('input.txt').read().splitlines())]

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
