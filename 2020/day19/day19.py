#!/usr/bin/python3
import re

def eval_rules(rule_no, cache):
	if rule_no in cache:
		return cache[rule_no]

	rule_text = ''

	rule = rules[rule_no]
	tokens = rule.split()
	for token in tokens:
		if token.isnumeric():
			rule_text += eval_rules(token, cache)
		else:
			rule_text += token

	if len(rule_text) > 1:
		rule_text = "(?:" + rule_text + ")"

	cache[rule_no] = rule_text
	return rule_text

def get_regex(rules):
	regex = eval_rules('0', {})
	regex = '^' + regex + '$'
	return regex

with open('input.txt') as f:
	line = f.readline().strip()

	rules = {}
	messages = []

        # rules
	while line != "":
		match = re.match("(\d+): (.*)", line)
		rule_no = match.group(1)
		rule = re.sub('"', '', match.group(2))
		rules[rule_no] = rule
		line = f.readline().strip()

	# messages
	line = f.readline().strip()
	while line:
		messages.append(line)
		line = f.readline().strip()

# part 1
regex = get_regex(rules)
valid = 0
for message in messages:
	if re.match(regex, message):
		valid += 1

print(valid)

# part 2
# these rules from the puzzle:
# rules['8'] = '42 | 42 8'
# rules['11'] = '42 31 | 42 11 31'

# converted by hand:
rules['8'] = '42 {1,}'
rules['11'] = '42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31'

regex = get_regex(rules)
valid = 0
for message in messages:
	if re.match(regex, message):
		valid += 1

print(valid)
