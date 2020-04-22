#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <map>
#include <queue>
#include <set>
#include <algorithm>
using namespace std;

struct Number {
	int x;
	int y;
	int number;
	map<int, int> distances;
	bool operator<(const Number& a) const {
		return number < a.number;
	}
};

struct Space {
	int x;
	int y;
	int depth;
	string serialise() {
		return to_string(x) + "_" + to_string(y);
	};
};

struct Move {
	int dx;
	int dy;
};

vector<Space> get_neighbours(vector<string> &maze, Space &space) {
	const vector<Move> moves = {
		Move{ .dx = -1, .dy =  0 },
		Move{ .dx =  1, .dy =  0 },
		Move{ .dx =  0, .dy = -1 },
		Move{ .dx =  0, .dy =  1 }
	};

	vector<Space> neighbours;
	for (auto move : moves) {
		Space neighbour = Space{ .x = space.x + move.dx, .y = space.y + move.dy };
		if (maze[neighbour.y][neighbour.x] != '#') {
			neighbours.push_back(neighbour);
		}
	}
	
	return neighbours;
}

vector<Number> find_numbers(vector<string> maze) {
	vector<Number> numbers;
	for (int y = 0; y < maze.size(); y++) {
		for (int x = 0; x < maze[y].size(); x++) {
			if (maze[y][x] != '#' && maze[y][x] != '.') {
				numbers.push_back(Number{ .x = x, .y = y, .number = maze[y][x] - '0' });
			}
		}
	}
	return numbers;
}


int bfs(vector<string> &maze, Number &n, Number &n2) {
	queue<Space> spaces;
	Space start = Space{ .x = n.x, .y = n.y, .depth = 0 };
	spaces.push(start);

	set<string> visited;

	while (spaces.size() > 0) {
		Space sp = spaces.front();
		spaces.pop();

		if (sp.x == n2.x && sp.y == n2.y) {
			return sp.depth;
		}

		string serial = sp.serialise();
		if (visited.count(serial) == 1) {
			continue;
		}
		visited.insert(serial);

		vector<Space> neighbours = get_neighbours(maze, sp);
		for (Space neighbour : neighbours) {
			neighbour.depth = sp.depth + 1;
			spaces.push(neighbour);
		}

	}
}

int main() {
	vector<string> maze;	

	string line;
        ifstream in("input.txt");
        while (getline(in, line)) {
		maze.push_back(line);
	}
	in.close();

	vector<Number> numbers = find_numbers(maze);
	sort(numbers.begin(), numbers.end());

	for (int i = 0; i < numbers.size(); i++) {
		for (int j = 0; j < numbers.size(); j++) {
			Number n = numbers[i];
			Number n2 = numbers[j];

			if (n.number == n2.number) {
				continue;
			}
			if (n.distances.count(n2.number) == 1) {
				continue;
			}
			int distance = bfs(maze, n, n2);
			numbers[i].distances.insert({n2.number, distance});
			numbers[j].distances.insert({n.number, distance});
		}
	}

	vector<int> sequence = { 0, 1, 2, 3, 4, 5, 6, 7 };
	int best_distance = 99999;

        do {
		if (sequence[0] != 0) {
			continue;
		}
		int distance = 0;
		for (int i = 0; i < sequence.size() - 1; i++) {
			distance += numbers[sequence[i]].distances[numbers[sequence[i+1]].number];
		}
		if (distance < best_distance) {
			best_distance = distance;
		}
        } while (std::next_permutation(sequence.begin(), sequence.end()));

	cout << best_distance << endl;

	return 0;
}
