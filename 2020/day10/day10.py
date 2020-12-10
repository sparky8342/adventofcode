#!/usr/bin/python3
import functools

@functools.lru_cache(maxsize=None)
def search(num):
	total = 0
	end = True
	for i in (num + 1, num + 2, num + 3):
		if i in data_set:
			end = False
			total += search(i)

	if end:
		total += 1

	return total

data = sorted([int(i) for i in open('input.txt').read().splitlines()])
data.insert(0, 0)
data.append(data[len(data) - 1] + 3)

# part1
diff_one = 0
diff_three = 0

for i in range(0, len(data) - 1):
	if data[i + 1] - data[i] == 1:
		diff_one += 1
	elif data[i + 1] - data[i] == 3:
		diff_three += 1

print(diff_one * diff_three)

# part 2
data_set = set(i for i in data)
print(search(0))
