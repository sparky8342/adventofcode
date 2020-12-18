#!/usr/bin/python3
import re

plus_precedence = False

def evaluate_match(match):
	return evaluate(match.group(1))

def add(match):
	return str(int(match.group(1)) + int(match.group(2)))

def evaluate(expr):
	while re.search("\(", expr):
		expr = re.sub("\(([^\(\)]+)\)", evaluate_match, expr)

	while plus_precedence and re.search("\+", expr):
		expr = re.sub("(\d+) \+ (\d+)", add, expr)

	parts = expr.split()
	value = int(parts[0])
	for i in range(1, len(parts), 2):
		value2 = int(parts[i + 1])
		if parts[i] == '+':
			value += value2
		elif parts[i] == '*':
			value *= value2

	return str(value)

lines = open('input.txt').read().splitlines()

total = 0
for line in lines:
	total += int(evaluate(line))
print(total)

plus_precedence = True
total = 0
for line in lines:
	total += int(evaluate(line))
print(total)
