#!/usr/bin/python3
import functools
import sys

part1 = 0
part2 = 0

data = open('input.txt').read().strip()
groups = data.split("\n\n")

for group in groups:
	people = group.split("\n")

	questions = set()
	for person in people:
		for q in person:
			questions.add(q)

	part1 += len(questions)

	question_sets = []
	for person in people:
		question_set = set()
		for q in person:
			question_set.add(q)
		question_sets.append(question_set)

	reduced = functools.reduce(lambda a,b : a.intersection(b), question_sets)
	part2 += len(reduced)

print(part1)
print(part2)
