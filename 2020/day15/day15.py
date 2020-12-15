#!/usr/bin/python3

def get(sequence, target):
	seen = {}
	for i, n in enumerate(sequence[:-1]):
		seen[n] = i
	last = sequence[-1]

	for pos in range(len(sequence) - 1, target - 1):
		nxt = 0
		if last in seen:
			nxt = pos - seen[last]

		seen[last] = pos
		last = nxt

	return last

sequence = [int(x) for x in open('input.txt').read().strip().split(",")]

print(get(sequence, 2020))
print(get(sequence, 30000000))
