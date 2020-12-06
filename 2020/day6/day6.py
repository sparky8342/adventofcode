#!/usr/bin/python3
import functools

part1 = 0
part2 = 0
with open('input.txt') as f:
	questions = set()
	question_sets = []
	for line in f:
		line = line.strip()

		if line:
			question_set = set()
			for char in line:
				questions.add(char)
				question_set.add(char)
			question_sets.append(question_set)
		else:
			part1 += len(questions)
			questions = set()
			reduced = functools.reduce(lambda a,b : a.intersection(b), question_sets)
			part2 += len(reduced)
			question_sets = []

	part1 += len(questions)
	reduced = functools.reduce(lambda a,b : a.intersection(b), question_sets)
	part2 += len(reduced)	

print(part1)
print(part2)
