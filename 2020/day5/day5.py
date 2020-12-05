#!/usr/bin/python3
import re

passes = open('input.txt').read().splitlines()

highest = 0
seats = {i for i in range(1024)}

for ps in passes:
	ps = re.sub('[FL]', '0', ps)	
	ps = re.sub('[BR]', '1', ps)
	id = int(ps, 2)
	highest = max(highest, id)
	seats.remove(id)

print(highest)

for seat in seats:
	if seat + 1 not in seats and seat - 1 not in seats:
		print(seat)
		break
