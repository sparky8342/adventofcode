#!/usr/bin/python3
from functools import reduce
from collections import defaultdict
import re

lines = open('input.txt').read().splitlines()

all_ingredients = defaultdict(int)
groups = {}

for line in lines:
	match = re.match("(.*?)\(contains (.*?)\)", line)
	ingredients = set(match.group(1).split())
	allergens = match.group(2).split(", ")
	for ingredient in ingredients:
		all_ingredients[ingredient] += 1
	for allergen in allergens:
		if allergen in groups:
			groups[allergen] = groups[allergen].intersection(ingredients)
		else:
			groups[allergen] = ingredients


# find singles and remove from other sets
def remove_allergen(groups, ingredient, keep):
	removed = False
	for allergen in groups:
		if allergen == keep:
			continue
		if ingredient in groups[allergen]:
			groups[allergen].remove(ingredient)
			removed = True
	return removed

done = False
while done == False:
	done = True
	for allergen in groups:
		a_set = groups[allergen]
		if len(a_set) == 1:
			(ingredient,) = a_set
			if remove_allergen(groups, ingredient, allergen):
				done = False

# remove ingredients with allergens
for allergen in groups:
	st = groups[allergen]
	(ingredient,) = st
	del all_ingredients[ingredient]

# part 1
print(sum(all_ingredients.values()))

# part 2
allergens = sorted(groups.keys())
ingredients = []
for allergen in allergens:
	(ingredient, ) = groups[allergen]
	ingredients.append(ingredient)

print(",".join(ingredients))
