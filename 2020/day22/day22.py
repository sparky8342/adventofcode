#!/usr/bin/python3
from collections import deque

data = open('input.txt').read().strip()
player_data = data.split("\n\n")
player1 = [int(x) for x in player_data[0].split("\n")[1:]]
player2 = [int(x) for x in player_data[1].split("\n")[1:]]

# part 1
p1 = deque(player1)
p2 = deque(player2)

while len(p1) > 0 and len(p2) > 0:
	top_cards = [p1.popleft(), p2.popleft()]
	if top_cards[0] > top_cards[1]:
		p1.append(top_cards[0])
		p1.append(top_cards[1])
	else:
		p2.append(top_cards[1])
		p2.append(top_cards[0])

winner = []
if len(p1) > 0:
	winner = p1
else:
	winner = p2

score = 0
for i, card in enumerate(winner):
	score += (len(winner) - i) * card

print(score)

# part 2
p1 = deque(player1)
p2 = deque(player2)

def get_key(q1, q2):
	return " ".join(str(x) for x in (list(q1) + [':'] + list(q2)))

cache = {}
def play(p1, p2):
	cache_key = get_key(p1, p2)
	if cache_key in cache:
		return cache[cache_key]

	seen = set()
	while len(p1) > 0 and len(p2) > 0:
		key = get_key(p1, p2)
		if key in seen:
			cache[cache_key] = [1, p1]
			return 1, p1
		seen.add(key)

		winner = 0
		top_cards = [p1.popleft(), p2.popleft()]
		if top_cards[0] <= len(p1) and top_cards[1] <= len(p2):
			cp1 = deque(list(p1)[0:top_cards[0]])
			cp2 = deque(list(p2)[0:top_cards[1]])
			winner, _ = play(cp1, cp2)
		elif top_cards[0] > top_cards[1]:
			winner = 1
		else:
			winner = 2

		if winner == 1:
			p1.append(top_cards[0])
			p1.append(top_cards[1])
		elif winner == 2:
			p2.append(top_cards[1])
			p2.append(top_cards[0])

	if len(p1) > 0:
		cache[cache_key] = [1, p1]
		return 1, p1
	else:
		cache[cache_key] = [2, p2]
		return 2, p2

_, cards = play(p1, p2)
score = 0
for i, card in enumerate(cards):
	score += (len(cards) - i) * card

print(score)
