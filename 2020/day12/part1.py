#!/usr/bin/python3

actions = [line for line in (open('input.txt').read().splitlines())]

directions = {
	'E' : ( 1,  0),
	'W' : (-1,  0),
	'N' : ( 0,  1),
	'S' : ( 0, -1)
}
dir_sequence = ['N', 'E', 'S', 'W']

x = 0
y = 0
direction = 'E'

for action in actions:
	cmd = action[0]
	num = int(action[1:])

	if cmd == 'N':
		y += num
	elif cmd == 'S':
		y -= num
	elif cmd == 'E':
		x += num
	elif cmd == 'W':
		x -= num
	elif cmd == 'F':
		dx, dy = directions[direction]
		x += dx * num
		y += dy * num
	elif cmd == 'L' or cmd == 'R':
		steps = int(num / 90)
		if cmd == 'L':
			steps *= -1
		for i in range(len(dir_sequence)):
			if dir_sequence[i] == direction:
				direction = dir_sequence[(i + steps) % 4]
				break
print(abs(x) + abs(y))
