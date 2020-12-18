#!/usr/bin/python3
import re

def evaluate_match(match):
	return evaluate(match.group(1))

def evaluate(expr):
	while re.search("\(", expr):
		expr = re.sub("\(([^\(\)]+)\)", evaluate_match, expr)

	parts = expr.split()
	value = int(parts[0])
	for i in range(1, len(parts), 2):
		value2 = int(parts[i + 1])
		if parts[i] == '+':
			value += value2
		elif parts[i] == '-':
			value -= value2
		elif parts[i] == '*':
			value *= value2
		elif parts[i] == '/':
			value /= value2

	return str(value)

total = 0
with open('input.txt') as f:
	line = f.readline().strip()
	while line:
		total += int(evaluate(line))
		line = f.readline().strip()

print(total)
