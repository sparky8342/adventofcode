#!/usr/bin/python3

mod = 20201227

with open('input.txt') as f:
	card_public = int(f.readline().strip())
	room_public = int(f.readline().strip())

loop_size = 0
n = 1
while n != card_public:
	n *= 7
	n = n % mod
	loop_size += 1

print(pow(room_public, loop_size, mod))
