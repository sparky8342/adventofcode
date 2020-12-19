#!/usr/bin/python3
import re

cache = {}
def eval_rules(rule_no):
	if rule_no in cache:
		return cache[rule_no]

	rule_text = ''

	rule = rules[rule_no]
	tokens = rule.split()
	for token in tokens:
		if token == '|':
			rule_text = rule_text + '|'
		elif token.isnumeric():
			rule_text = rule_text + eval_rules(token) 
		else:
			rule_text = rule_text + token

	if len(rule_text) > 1:
		rule_text = "(" + rule_text + ")"

	cache[rule_no] = rule_text
	return rule_text

with open('input.txt') as f:
	line = f.readline().strip()

	rules = {}
	messages = []

        # rules
	while line != "":
		match = re.match("(\d+): \"(\w)\"", line)
		if match:
			rules[match.group(1)] = match.group(2)
		else:
			match = re.match("(\d+): (.*)", line)
			rules[match.group(1)] = match.group(2)
		line = f.readline().strip()

	line = f.readline().strip()
	while line:
		messages.append(line)
		line = f.readline().strip()

regex = eval_rules('0')
regex = '^' + regex + '$'

valid = 0
for message in messages:
	if re.match(regex, message):
		valid += 1

print(valid)
