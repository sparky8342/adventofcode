#!/usr/bin/python3
import functools

part1 = 0
part2 = 0

data = open('input.txt').read().strip()
groups = data.split("\n\n")

for group in groups:
	people = group.split("\n")

	question_sets = []
	for person in people:
		question_set = {q for q in person}
		question_sets.append(question_set)

	part1 += len(functools.reduce(lambda a,b : a.union(b), question_sets))
	part2 += len(functools.reduce(lambda a,b : a.intersection(b), question_sets))

print(part1)
print(part2)
