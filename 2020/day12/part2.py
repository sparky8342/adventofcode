#!/usr/bin/python3

actions = [line for line in (open('input.txt').read().splitlines())]

ship_x = 0
ship_y = 0
waypoint_x = 10
waypoint_y = 1

for action in actions:
	cmd = action[0]
	num = int(action[1:])

	if cmd == 'N':
		waypoint_y += num
	elif cmd == 'S':
		waypoint_y -= num
	elif cmd == 'E':
		waypoint_x += num
	elif cmd == 'W':
		waypoint_x -= num
	elif cmd == 'F':
		ship_x += waypoint_x * num
		ship_y += waypoint_y * num
	elif cmd == 'R':
		steps = int(num / 90)
		for i in range(steps):
			waypoint_x, waypoint_y = waypoint_y, waypoint_x * -1
	elif cmd == 'L':
		steps = int(num / 90)
		for i in range(steps):
			waypoint_x, waypoint_y = waypoint_y * -1, waypoint_x

print(abs(ship_x) + abs(ship_y))
