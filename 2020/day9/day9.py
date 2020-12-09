#!/usr/bin/python3

def valid(numbers, target):
	nset = set()
	for number in numbers:
		nset.add(number)
		if target - number in numbers:
			return True
	return False

def sum_range(numbers, start, end):
	total = 0
	for i in range(start, end + 1):
		total += numbers[i]
	return total

data = [int(i) for i in open('input.txt').read().splitlines()]
preamble_size = 25

# part 1
numbers = data[0:preamble_size]
testdata = data[preamble_size:]

while valid(numbers, testdata[0]):
	numbers = numbers[1:]
	numbers.append(testdata[0])
	testdata = testdata[1:]

part1 = testdata[0]
print(part1)

# part 2
start = 0
end = 0
total = data[0]
while total != part1:
	if total < part1:
		end += 1
	elif total > part1:
		start += 1
	total = sum_range(data, start, end)

numbers = [data[i] for i in range(start, end + 1)]
print(min(numbers) + max(numbers))
