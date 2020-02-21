#include <iostream>
#include <string>
#include <fstream>
#include <vector>
using namespace std;

struct Pos {
	int x;
	int y;
};

int main() {
	vector<string> instructions;
	string line;
	ifstream in("input.txt");
	while (getline(in, line)) {
		instructions.push_back(line);
	}
	in.close();

	vector<vector<string>> grids;
	vector<Pos> starts;

	vector<string> grid;
	grid.push_back("123");
	grid.push_back("456");
	grid.push_back("789");
	grids.push_back(grid);
	Pos pos = Pos{ .x = 1, .y = 1 };
	starts.push_back(pos);

	vector<string> grid2;
	grid2.push_back("00100");
	grid2.push_back("02340");
	grid2.push_back("56789");
	grid2.push_back("0ABC0");
	grid2.push_back("00D00");
	grids.push_back(grid2);
	Pos pos2 = Pos{ .x = 0, .y = 2 };
	starts.push_back(pos2);

	for (int i = 0; i < grids.size(); i++) {
		grid = grids[i];
		Pos startpos = starts[i];
		int x = startpos.x;
		int y = startpos.y;

		for (string& line : instructions) {
			for (char& direction : line) {
				if (direction == 'U' && y > 0 && grid[y-1][x] != '0') {
					y--;
				}
				else if (direction == 'D' && y < grid.size() - 1 && grid[y+1][x] != '0') {
					y++;
				}
				else if (direction == 'L' && x > 0 && grid[y][x-1] != '0') {
					x--;
				}
				else if (direction == 'R' && x < grid.size() - 1 && grid[y][x+1] != '0') {
					x++;
				}
			}
			cout << grid[y][x];
		}
		cout << endl;
	}
	return 0;
}
