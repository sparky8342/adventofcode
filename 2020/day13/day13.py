#!/usr/bin/python3

with open('input.txt') as f:
	leave_time = int(f.readline().strip())
	buses = [int(x) for x in f.readline().strip().split(",") if x != 'x']

times = []
for bus in buses:
	times.append((bus - leave_time % bus, bus))

times = sorted(times)
time = times[0]
print(time[0] * time[1])
