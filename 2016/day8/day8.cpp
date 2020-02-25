#include <iostream>
#include <string>
#include <fstream>
#include <regex>
using namespace std;

#define WIDTH 50
#define HEIGHT 6

int main() {
	bool grid[HEIGHT][WIDTH];
	for (int i = 0; i < HEIGHT; i++) {
		for (int j = 0; j < WIDTH; j++) {
			grid[i][j] = false;
		}
	}

	const regex rect("^rect ([0-9]+)x([0-9]+)$");
	const regex ro_col("^rotate column x=([0-9]+) by ([0-9]+)$");
	const regex ro_row("^rotate row y=([0-9]+) by ([0-9]+)$");

	string line;
	ifstream in("input.txt");
	while (getline(in, line)) {
		smatch base_match;
                if (regex_match(line, base_match, rect)) {
			int x = stoi(base_match[1]);
			int y = stoi(base_match[2]);
			for (int dy = 0; dy < y; dy++) {
				for (int dx = 0; dx < x; dx++) {
					grid[dy][dx] = true;
				}
			}
		}
		else if (regex_match(line, base_match, ro_col)) {
			int col = stoi(base_match[1]);
			int no = stoi(base_match[2]);

			bool col_array[HEIGHT];
			for (int i = 0; i < HEIGHT; i++) {
				col_array[i] = grid[i][col];
			}

			bool new_col[HEIGHT];
			for (int i = 0; i < HEIGHT; i++) {
				int p = i - no;
				if (p < 0) {
					p += HEIGHT;
				} 
				new_col[i] = col_array[p];
			}

			for (int i = 0; i < HEIGHT; i++) {
				grid[i][col] = new_col[i];
			}
		}
		else if (regex_match(line, base_match, ro_row)) {
			int row = stoi(base_match[1]);
			int no = stoi(base_match[2]);

			bool row_array[WIDTH];
			for (int i = 0; i < WIDTH; i++) {
				row_array[i] = grid[row][i];
			}

			bool new_row[WIDTH];
			for (int i = 0; i < WIDTH; i++) {
				int p = i - no;
				if (p < 0) {
					p += WIDTH;
				} 
				new_row[i] = row_array[p];
			}

			for (int i = 0; i < WIDTH; i++) {
				grid[row][i] = new_row[i];
			}
		}
	}
	in.close();

	int on = 0;
	for (int i = 0; i < HEIGHT; i++) {
		for (int j = 0; j < WIDTH; j++) {
			if (grid[i][j] == true) {
				on++;
				cout << '#';
			}
			else {
				cout << '.';
			}
		}
		cout << endl;
	}

	cout << on << endl;
	return 0;
}
