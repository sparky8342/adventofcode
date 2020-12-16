#!/usr/bin/python3
import re

rules = {}
tickets = []
valid_tickets = []
rule_set = set()
my_ticket = []
no_fields = 0

def in_any_rule(rules, value):
	for r in rules.values():
		for ra in r:
			if ra[0] <= value <= ra[1]:
				return True
	return False

def valid_for_rule(rules, rule, value):
	for r in rules[rule]:
		if r[0] <= value <= r[1]:
			return True
	return False

with open('input.txt') as f:
	line = f.readline().strip()

	# rules
	while line != "":
		match = re.match("^(.*?): (\d+)-(\d+) or (\d+)-(\d+)$", line)
		name = match.group(1)
		range1 = (int(match.group(2)), int(match.group(3)))
		range2 = (int(match.group(4)), int(match.group(5)))
		rules[name] = [ range1, range2 ]
		rule_set.add(name)
		line = f.readline().strip()

	f.readline() # 'your ticket:'
	my_ticket = [int(x) for x in f.readline().strip().split(",")]
	valid_tickets.append(my_ticket)
	f.readline() # blank line
	f.readline() # 'nearby tickets:'
        
	line = f.readline()
	while line:
		tickets.append([int(x) for x in line.strip().split(",")])
		line = f.readline()

no_fields = len(rule_set)

# part 1
error = 0
for ticket in tickets:
	ticket_ok = True
	for value in ticket:
		if not in_any_rule(rules, value):
			error += value
			ticket_ok = False

	if ticket_ok == True:
		valid_tickets.append(ticket)

print(error)

# part 2
# set up sets of possible fields for each place
possible_fields = []
for _ in range(no_fields):
	possible_fields.append(rule_set.copy())

# remove possiblities based on tickets
for i, fields in enumerate(possible_fields):
	bad_fields = set()
	for field in fields:
		for ticket in valid_tickets:
			if not valid_for_rule(rules, field, ticket[i]):
				bad_fields.add(field)
				break

	possible_fields[i] -= bad_fields

# remove possibilities where one field is left
done = False
while not done:
	done = True
	for i, fields in enumerate(possible_fields):
		if len(fields) == 1:
			(field,) = fields
			for j in range(no_fields):
				if i != j and field in possible_fields[j]:
					possible_fields[j].remove(field)
					done = False

answer = 1
for i, fields in enumerate(possible_fields):
	(field,) = fields
	if re.match("^departure", field):
		answer *= my_ticket[i]

print(answer)
