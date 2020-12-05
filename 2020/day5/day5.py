#!/usr/bin/python3

passes = open('input.txt').read().splitlines()

highest = 0
seats = {i for i in range(1024)}

for ps in passes:
	row_part = ps[:7]
	col_part = ps[7:]	

	mn = 0
	mx = 127
	for letter in row_part:
		if letter == 'F':
			mx = mn + int((mx - mn) / 2)
		elif letter == 'B':
			mn = mn + int((mx - mn) / 2) + 1
	row = mn

	mn = 0
	mx = 7
	for letter in col_part:
		if letter == 'L':
			mx = mn + int((mx - mn) / 2)
		elif letter == 'R':
			mn = mn + int((mx - mn) / 2) + 1

	col = mn
	id = row * 8 + col
	highest = max(highest, id)
	seats.remove(id)

print(highest)

for seat in seats:
	if seat + 1 not in seats and seat - 1 not in seats:
		print(seat)
		break
