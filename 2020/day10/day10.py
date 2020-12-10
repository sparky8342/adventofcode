#!/usr/bin/python3

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
dp = [0] * (max(numbers) + 1)
dp[0] = 1
for n in numbers:
	for i in (n-1, n-2, n-3):
		if i >= 0:
			dp[n] += dp[i]

print(dp[max(numbers)])
