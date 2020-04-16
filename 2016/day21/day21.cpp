#include <iostream>
#include <fstream>
#include <regex>
using namespace std;

vector<char> scramble(vector<char>code, vector<string> &operations) {
	const regex swap_pos("^swap position ([0-9]+) with position ([0-9]+)$");
	const regex swap_letter("^swap letter ([a-z]) with letter ([a-z])$");
	const regex rotate_left("^rotate left ([0-9]+) step(s)?$");
	const regex rotate_right("^rotate right ([0-9]+) step(s)?$");
	const regex rotate_pos("^rotate based on position of letter ([a-z])$");
	const regex reverse_pos("^reverse positions ([0-9]+) through ([0-9]+)$");
	const regex move("^move position ([0-9]+) to position ([0-9]+)$");

	vector<string>::iterator it;
	for (it=operations.begin(); it!=operations.end(); ++it) {
		smatch match;
		if (regex_match(*it, match, swap_pos)) {
			int x = stoi(match[1]);
			int y = stoi(match[2]);
			char tmp = code[x];
			code[x] = code[y];
			code[y] = tmp;
		}
		else if (regex_match(*it, match, swap_letter)) {
			int x;
			int y;
			for (int i = 0; i < code.size(); i++) {
				if (code[i] == match[1]) {
					x = i;
				}
				else if (code[i] == match[2]) {
					y = i;
				}
			}
			char tmp = code[x];
			code[x] = code[y];
			code[y] = tmp;
		}
		else if (regex_match(*it, match, rotate_left)) {
			int x = stoi(match[1]);
			rotate(code.begin(), code.begin() + x, code.end()); 
		}
		else if (regex_match(*it, match, rotate_right)) {
			int x = stoi(match[1]);
			rotate(code.begin(), code.begin() + code.size() - x, code.end()); 
		}
		else if (regex_match(*it, match, rotate_pos)) {
			int x;
			for (int i = 0; i < code.size(); i++) {
				if (code[i] == match[1]) {
					x = i;
					break;
				}
			}
			if (x >= 4) {
				x++;
			}
			x++;
			x = x % code.size();
			rotate(code.begin(), code.begin() + code.size() - x, code.end());
		}
		else if (regex_match(*it, match, reverse_pos)) {
			int x = stoi(match[1]);
			int y = stoi(match[2]);
			vector<char> new_code = code;

			int newpos = x;	
			for (int i = y; i >= x; i--) {
				new_code[newpos] = code[i];
				newpos++;
			}
			code = new_code;
		}
		else if (regex_match(*it, match, move)) {
			int x = stoi(match[1]);
			int y = stoi(match[2]);
			int val = code[x];
			code.erase(code.begin() + x);
			code.insert(code.begin() + y, val);
		}
	}

	return code;
}

int main() {
	vector<string> operations;
	string line;
	ifstream in("input.txt");
	while (getline(in, line)) {
		operations.push_back(line);
	}
	in.close();

	vector<char> code = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'};
	code = scramble(code, operations);

	for (int i = 0; i < code.size(); i++) {
		cout << code[i];
	}
	cout << endl;

	return 0;
}
