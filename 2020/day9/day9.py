#!/usr/bin/python3

def valid(numbers, target):
	nset = set()
	for number in numbers:
		nset.add(number)
		if target - number in numbers:
			return True
	return False

data = [int(i) for i in open('input.txt').read().splitlines()]
PRE_SIZE = 25

# part 1
part1 = 0
for i in range (0, len(data)):
	num = i + PRE_SIZE
	numbers = data[i:num]
	if not valid(numbers, data[num]):
		part1 = data[num]
		break
print(part1)

# part 2
start = 0
end = 0
total = data[0]
while total != part1:
	if total < part1:
		end += 1
		total += data[end]
	elif total > part1:
		total -= data[start]
		start += 1

numbers = [data[i] for i in range(start, end + 1)]
print(min(numbers) + max(numbers))
