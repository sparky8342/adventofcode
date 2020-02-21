#include <iostream>
#include <string>
#include <fstream>
#include <vector>
using namespace std;

int main() {
	int grid[3][3];	
	int n = 1;
	for (int y = 0; y < 3; y++) {
		for (int x = 0; x < 3; x++) {
			grid[x][y] = n;
			n++;
		}
	}

	vector<string> instructions;
	string line;
	ifstream in("input.txt");
	while (getline(in, line)) {
		instructions.push_back(line);
	}
	in.close();

	int x = 1;
	int y = 1;

	for (string& line : instructions) {	
		for (char& direction : line) {
			if (direction == 'U' && y > 0) {
				y--;
			}
			else if (direction == 'D' && y < 2) {
				y++;
			}
			else if (direction == 'L' && x > 0) {
				x--;
			}
			else if (direction == 'R' && x < 2) {
				x++;
			}
		}
		cout << grid[x][y];
	}

	cout << endl;
	return 0;
}
