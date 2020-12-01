#!/usr/bin/python3

def sum2(numbers, target):
	i = 0
	j = len(numbers) - 1
	k = numbers[i] + numbers[j]

	while k != target:
		if k > target:
			j -= 1
		else:
			i += 1
		k = numbers[i] + numbers[j]

	return numbers[i], numbers[j]

with open('input.txt') as f:
	numbers = [int(i) for i in f.read().splitlines()]
numbers.sort()
target = 2020

num1, num2 = sum2(numbers, target)
print(num1 * num2)

for i in numbers:
	num1, num2 = sum2(numbers, target - i)
	if num1 != i and num2 != i:
		print(i * num1 * num2)
		break
