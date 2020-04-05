#include <iostream>
#include <string>
#include <vector>
#include <queue>
#include <set>
using namespace std;

#define INPUT_NO 1358
#define TARGET_X 31
#define TARGET_Y 39

struct Space {
	int x;
	int y;
	int depth;
};

struct Move {
	int dx;
	int dy;
};

bool is_open(int x, int y) {
	if (x < 0 || y < 0) {
		return false;
	}
	int n = (x*x + 3*x + 2*x*y + y + y*y) + INPUT_NO;
	int count = 0; 
	while (n) { 
		n &= (n - 1); 
		count++; 
	}
        return count % 2 == 0;
}

vector<Space> get_neighbours(Space space) {
	const vector<Move> moves = {
		Move{ .dx = -1, .dy =  0 },
		Move{ .dx =  1, .dy =  0 },
		Move{ .dx =  0, .dy = -1 },
		Move{ .dx =  0, .dy =  1 }
	};

	vector<Space> neighbours;
	for (auto move : moves) {
		Space neighbour = Space{ .x = space.x + move.dx, .y = space.y + move.dy, .depth = space.depth + 1 };
		if (is_open(neighbour.x, neighbour.y)) {
			neighbours.push_back(neighbour);
		}
	}
	
	return neighbours;
}

int main() {
	queue <Space> spaces;	
	spaces.push(Space{ .x = 1, .y = 1, .depth = 0 });

	set <string> visited;

	int within_fifty = 0;
	
	while (spaces.size() > 0) {
		Space space = spaces.front();
		spaces.pop();

		if (space.x == TARGET_X && space.y == TARGET_Y) {
			cout << space.depth << endl << within_fifty << endl;
			break;
		}

		string serial = to_string(space.x) + "_" + to_string(space.y);
		if (visited.count(serial) == 1) {
			continue;
		}
		visited.insert(serial);

		if (space.depth <= 50) {
			within_fifty++;
		}

		vector<Space> neighbours = get_neighbours(space);
		for (auto neighbour : neighbours) {
			spaces.push(neighbour);
		}
	}

	return 0;
}
