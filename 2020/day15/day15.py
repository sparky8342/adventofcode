#!/usr/bin/python3

def get(sequence, target):
	seen = {}
	for i, n in enumerate(sequence[:-1]):
		seen[n] = i
	last = sequence[-1]
	pos = len(sequence) - 1

	while pos < target - 1:
		nxt = 0
		if last in seen:
			nxt = pos - seen[last]

		seen[last] = pos
		last = nxt
		pos += 1

	return last

sequence = [int(x) for x in open('input.txt').read().strip().split(",")]

print(get(sequence, 2020))
print(get(sequence, 30000000))
