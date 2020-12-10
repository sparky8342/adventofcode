#!/usr/bin/python3
import functools

@functools.lru_cache(maxsize=None)
def search(num):
	total = 0
	end = True
	for i in (num + 1, num + 2, num + 3):
		if i in numbers:
			end = False
			total += search(i)

	if end:
		total += 1

	return total

numbers = set([int(i) for i in open('input.txt').read().splitlines()])
numbers.add(0)
numbers.add(max(numbers) + 3)

# part1
diff_one = 0
diff_three = 0
n = 0
while(1):
	if n + 1 in numbers:
		diff_one += 1
		n += 1
	elif n + 3 in numbers:
		diff_three += 1
		n += 3
	else:
		break

print(diff_one * diff_three)

# part 2
print(search(0))
