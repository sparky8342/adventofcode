#!/usr/bin/python3
import re
import functools

class Node:
	def __init__(self, name):
		self.name = name
		self.parents = []
		self.children = []

class ChildLink:
	def __init__(self, child, amount):
		self.child = child
		self.amount = amount

# walk up tree to find all parents
@functools.lru_cache(maxsize=None) # memoize this function
def find_parents(node):
	parents = set()
	for parent in node.parents:
		parents.add(parent.name)
		parents = parents.union(find_parents(parent))
	return parents

# walk down to find all children and calculate amount
@functools.lru_cache(maxsize=None)
def count_bags(node):
	total = 1
	for child_link in node.children:
		total += child_link.amount * count_bags(child_link.child)
	return total	

def get_node(name):
	if name not in nodes:
		nodes[name] = Node(name)
	return nodes[name]

nodes = {}
data = open('input.txt').read().splitlines()

for line in data:
	match = re.search("(.+) bags contain (.+)\.", line)
	name = match.group(1)
	children = match.group(2)
	children = re.sub(" bag(s)?", "", children)
	children = re.sub(", ", ",", children)

	node = get_node(name)

	if children != 'no other':
		children = children.split(",")
		for child in children:
			amount = int(child[0])
			child_name = child[2:]
			child_node = get_node(child_name)
			node.children.append(ChildLink(child_node, amount))
			child_node.parents.append(node)
	
	nodes[name] = node

shiny_gold = nodes['shiny gold']
print(len(find_parents(shiny_gold)))
print(count_bags(shiny_gold) - 1)
