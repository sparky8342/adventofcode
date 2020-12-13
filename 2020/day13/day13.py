#!/usr/bin/python3

with open('input.txt') as f:
	leave_time = int(f.readline().strip())
	data = [x for x in f.readline().strip().split(",")]

# part 1
buses = [int(x) for x in data if x != 'x']
times = []
for bus in buses:
	times.append((bus - leave_time % bus, bus))

times.sort()
time = times[0]
print(time[0] * time[1])

# part 2
buses = []
for i, bus in enumerate(data):
	if bus == 'x':
		continue
	buses.append((int(bus), i))

position = 0
inc = 1
for bus in buses:
	time, offset = bus
	while (position + offset) % time != 0:
		position += inc
	inc *= time

print(position)
