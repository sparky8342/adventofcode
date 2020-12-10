#!/usr/bin/python3

data = sorted([int(i) for i in open('input.txt').read().splitlines()])
data.insert(0, 0)
data.append(data[len(data) - 1] + 3)

diff_one = 0
diff_three = 0

for i in range(0, len(data) - 1):
	if data[i + 1] - data[i] == 1:
		diff_one += 1
	elif data[i + 1] - data[i] == 3:
		diff_three += 1

print(diff_one * diff_three)
