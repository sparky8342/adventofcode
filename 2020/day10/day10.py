#!/usr/bin/python3

numbers = sorted([int(i) for i in open('input.txt').read().splitlines()])
numbers.insert(0, 0)
numbers.append(numbers[len(numbers) - 1] + 3)

# part1
diff_one = 0
diff_three = 0
for i in range(0, len(numbers) - 1):
	if numbers[i + 1] - numbers[i] == 1:
		diff_one += 1
	elif numbers[i + 1] - numbers[i] == 3:
		diff_three += 1

print(diff_one * diff_three)

# part 2
dp = [0] * (max(numbers) + 1)
dp[0] = 1
for n in numbers:
	for i in (n-1, n-2, n-3):
		if i >= 0:
			dp[n] += dp[i]

print(dp[max(numbers)])
