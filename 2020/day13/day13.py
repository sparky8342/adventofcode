#!/usr/bin/python3

with open('input.txt') as f:
	leave_time = int(f.readline().strip())
	buses = [x for x in f.readline().strip().split(",")]

# part 1
times = []
for bus in buses:
	if bus == 'x':
		continue
	bus = int(bus)
	times.append((bus - leave_time % bus, bus))

times.sort()
time = times[0]
print(time[0] * time[1])

# part 2
targets = []
for i, bus in enumerate(buses):
	if bus == 'x':
		continue
	targets.append((int(bus), i))

position = 0
inc = 1
for target in targets:
	time, offset = target
	while (position + offset) % time != 0:
		position += inc
	inc *= time

print(position)
