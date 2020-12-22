#!/usr/bin/python3

data = open('input.txt').read().strip()
player_data = data.split("\n\n")
player1 = [int(x) for x in player_data[0].split("\n")[1:]]
player2 = [int(x) for x in player_data[1].split("\n")[1:]]
	
while len(player1) > 0 and len(player2) > 0:
	top_cards = [player1.pop(0), player2.pop(0)]
	if top_cards[0] > top_cards[1]:
		player1.extend(top_cards)
	else:
		player2.append(top_cards[1])
		player2.append(top_cards[0])

winner = []
if len(player1) > 0:
	winner = player1
else:
	winner = player2

score = 0
for i, card in enumerate(winner):
	score += (len(winner) - i) * card

print(score)
