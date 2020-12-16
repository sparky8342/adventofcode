#!/usr/bin/python3
import re

rules = {}
tickets = []
all_ranges = []
valid_tickets = []
rule_set = set()
my_ticket = []

with open('input.txt') as f:
	line = f.readline().strip()

	# rules
	while line != "":
		match = re.match("^(.*?): (\d+)-(\d+) or (\d+)-(\d+)$", line)
		name = match.group(1)
		range1 = (int(match.group(2)), int(match.group(3)))
		range2 = (int(match.group(4)), int(match.group(5)))
		rules[name] = [ range1, range2 ]
		all_ranges.append(range1)
		all_ranges.append(range2)
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

# part 1
error = 0
for ticket in tickets:
	ticket_ok = True
	for value in ticket:
		ok = False
		for rnge in all_ranges:
			if rnge[0] <= value <= rnge[1]:
				ok = True
				break
		if ok == False:
			ticket_ok = False
			error += value

	if ticket_ok == True:
		valid_tickets.append(ticket)

print(error)

# set up sets of possible fields for each place
possible_fields = []
for _ in range(len(rules)):
	possible_fields.append(rule_set.copy())

# remove possiblities based on tickets
for i, fields in enumerate(possible_fields):
	bad_fields = set()
	for field in fields:
		field_done = False
		for ticket in valid_tickets:
			if field_done == True:
				break
			range_ok = False
			for rnge in rules[field]:
				if rnge[0] <= ticket[i] <= rnge[1]:
					range_ok = True
			if range_ok == False:	
				bad_fields.add(field)
				field_done = True

	possible_fields[i] -= bad_fields

# remove possibilities where one field is left
while True:
	# if there's 1 left in each set, we are done
	left = 0
	for fields in possible_fields:
		left += len(fields)
	if left == len(possible_fields):
		break

	for i, fields in enumerate(possible_fields):
		if len(fields) == 1:
			(field,) = fields
			for j in range(len(possible_fields)):
				if i != j and field in possible_fields[j]:
					possible_fields[j].remove(field)

answer = 1
for i, fields in enumerate(possible_fields):
	(field, ) = fields
	if re.match("^departure", field):
		answer *= my_ticket[i]

print(answer)
