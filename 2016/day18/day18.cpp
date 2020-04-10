#include <iostream>
#include <fstream>
using namespace std;

int findsafe(string line, int rows) {
	int safe = 0;
	for (int i = 0; i < line.size(); i++) {
		if (line[i] == '.') {
			safe++;
		}
	}

	for (int r = 0; r < rows - 1; r++) {
		line = "." + line + ".";

		string newline;
		for (int i = 1; i < line.size() - 1; i++) {
			/*		
			Space is a trap if:
			Its left and center tiles are traps, but its right tile is not.
			Its center and right tiles are traps, but its left tile is not.
			Only its left tile is a trap.
			Only its right tile is a trap.
			*/
			if (
				(line[i-1] == '^' && line[i] == '^' && line[i+1] == '.') ||
				(line[i-1] == '.' && line[i] == '^' && line[i+1] == '^') ||
				(line[i-1] == '^' && line[i] == '.' && line[i+1] == '.') ||
				(line[i-1] == '.' && line[i] == '.' && line[i+1] == '^')
			) {
				newline = newline + "^";
			}
			else {
				newline = newline + ".";
				safe++;
			}
		}

		line = newline;
	}

	return safe;
}

int main() {
	string line;
	ifstream in("input.txt");
	getline(in, line);
	in.close();

	int safe = findsafe(line, 40);
	cout << safe << endl;

	safe = findsafe(line, 400000);
	cout << safe << endl;

	return 0;
}
