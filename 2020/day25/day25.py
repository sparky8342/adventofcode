#!/usr/bin/python3

mod = 20201227

with open('input.txt') as f:
	card_public = int(f.readline().strip())
	room_public = int(f.readline().strip())

loop_size = 0
n = 1
while True:
	n *= 7
	n = n % mod
	loop_size += 1
	if n == card_public:
		break

n = 1
for _ in range(loop_size):
	n *= room_public
	n = n % mod

print(n)
