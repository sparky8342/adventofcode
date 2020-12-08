#!/usr/bin/python3

program = [line.split() for line in open('input.txt').read().splitlines()]
for line in program:
	line[1] = int(line[1])

acm = 0
pos = 0
seen = set()

while 1:
	if pos in seen:
		print(acm)
		break
	seen.add(pos)
	ins, arg = program[pos]
	if ins == "acc":
		acm += arg
		pos += 1
	elif ins == "jmp":
		pos += arg
	elif ins == "nop":
		pos += 1
