#!/usr/bin/python3
import re

# walk up tree to find all parents
def find_parents(node):
	parents = set()
	for parent in nodes[node]['parents']:
		parents.add(parent)
		parents = parents.union(find_parents(parent))
	return parents

# walk down to find all children and calculate amount
def count_bags(node):
	total = 1
	for child in nodes[node]['children']:
		total += child['amount'] * count_bags(child['name'])
	return total	

nodes = {}
data = open('input.txt').read().splitlines()

for line in data:
	match = re.search("(.+) bags contain (.+)\.", line)
	name = match.group(1)
	children = match.group(2)
	children = re.sub(" bag(s)?", "", children)
	children = re.sub(", ", ",", children)

	node = { 'parents' : [], 'children' : [] }	

	if children != 'no other':
		children = children.split(",")
		for child in children:
			amount = int(child[0])
			child_name = child[2:]
			node['children'].append({ 'name' : child_name, 'amount' : amount })
	
	nodes[name] = node

# set parents
for parent, node in nodes.items():
	for child in node['children']:
		name = child['name']
		nodes[name]['parents'].append(parent)

print(len(find_parents('shiny gold')))
print(count_bags('shiny gold') - 1)
