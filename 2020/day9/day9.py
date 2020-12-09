#!/usr/bin/python3

def valid(numbers, target):
	nset = set()
	for number in numbers:
		nset.add(number)
		if target - number in numbers:
			return True
	return False

data = [int(i) for i in open('input.txt').read().splitlines()]

numbers = data[0:25]
data = data[25:]

while(1):
	if not valid(numbers, data[0]):
		print(data[0])
		break

	numbers = numbers[1:]
	numbers.append(data[0])
	data = data[1:]
