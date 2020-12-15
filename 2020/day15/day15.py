#!/usr/bin/python3

def get(sequence, target):
	seen1 = {}
	seen2 = {}

	for i, n in enumerate(sequence):
		seen2[n] = i

	last = sequence[-1]
	pos = len(sequence)

	while pos < target:
		nxt = 0
		if last in seen1:
			nxt = seen2[last] - seen1[last]

		if nxt in seen2:
			seen1[nxt] = seen2[nxt]
		seen2[nxt] = pos

		pos += 1
		last = nxt

	return last

sequence = [int(x) for x in open('input.txt').read().strip().split(",")]

print(get(sequence, 2020))
print(get(sequence, 30000000))
