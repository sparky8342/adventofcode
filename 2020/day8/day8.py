#!/usr/bin/python3

def run(program):
	acm = 0
	pos = 0
	seen = set()

	while pos >= 0 and pos < len(program):
		if pos in seen:
			return True, acm
		seen.add(pos)
		ins, arg = program[pos]
		if ins == "acc":
			acm += arg
			pos += 1
		elif ins == "jmp":
			pos += arg
		elif ins == "nop":
			pos += 1

	return False, acm

program = [line.split() for line in open('input.txt').read().splitlines()]
for line in program:
	line[1] = int(line[1])

looped, acm = run(program)
print(acm)

for i in range(len(program)):
	ins = program[i][0]
	if ins == "nop":
		program[i][0] = "jmp"
	elif ins == "jmp":
		program[i][0] = "nop"
	else:
		continue

	looped, acm = run(program)
	if looped == False:
		print(acm)
		break

	program[i][0] = ins
