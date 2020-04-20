#include <iostream>
#include <fstream>
#include <regex>
using namespace std;

#define WIDTH 30
#define HEIGHT 35

struct Node {
	int x;
	int y;
	int used;
	int avail;
};

int main() {
	const regex node_regex("^/dev/grid/node-x([0-9]+)-y([0-9]+)[ ]+[0-9]+T[ ]+([0-9]+)T[ ]+([0-9]+)T[ ]+[0-9]+%$");
	vector<Node> nodes;
	char grid[HEIGHT][WIDTH];

	string line;
	smatch match;
	ifstream in("input.txt");
	getline(in,line);
	getline(in,line);
	while(getline(in, line)) {
		regex_match(line, match, node_regex);
		int x = stoi(match[1]);
		int y = stoi(match[2]);
		int used = stoi(match[3]);
		int avail = stoi(match[4]);
		Node node = Node {
			.x = x,
			.y = y,
			.used = used,
			.avail = avail,
		};
		// for part 1
		nodes.push_back(node);
		// for part 2
		if (x == WIDTH - 1 && y == 0) {
			grid[y][x] = 'G';
		}
		else if (used > 100) {
			grid[y][x] = '#';
		}
		else if (used == 0) {
			grid[y][x] = '_';
		}
		else {
			grid[y][x] = '.';
		}
	}
	in.close();

	// part 1
	int pairs = 0;
	for (int i = 0; i < nodes.size(); i++) {
		for (int j = 0; j < nodes.size(); j++) {
			if (i != j && nodes[i].used != 0 && nodes[i].used <= nodes[j].avail) {
			       pairs++;
		      	}	       
		}
	}
	cout << pairs << endl;

	// part 2
	/*
	for part 2, just print out the grid and solve by hand
	move the empty space _ next to the G,
	then rotate around to move the G to the top left

	45 moves to get next to the G
	then 28 * 5 to move to the top left
	answer is 185
	*/
	for (int y = 0; y < HEIGHT; y++) {
		for (int x = 0; x < WIDTH; x++) {
			cout << grid[y][x];
		}
		cout << endl;
	}
	cout << endl;

	return 0;
}
