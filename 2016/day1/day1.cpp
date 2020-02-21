#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <map>
using namespace std;

struct Move {
	int dx;
	int dy;
};

vector<string> split(string& s, const string& delimiter) {
	std::vector<std::string> tokens;
	size_t pos = 0;
	std::string token;
	while ((pos = s.find(delimiter)) != std::string::npos) {
		token = s.substr(0, pos);
		tokens.push_back(token);
		s.erase(0, pos + delimiter.length());
	}
	if (s.length() > 0) {
		tokens.push_back(s);
	}
	return tokens;
}

int main() {
        string input;
        ifstream in("input.txt");
        getline(in, input);
        in.close();

        std::vector<std::string> directions;
        directions = split(input, ", ");

	map<char, map<char, char>> turns;
	turns['^']['L'] = '<';
	turns['^']['R'] = '>';
	turns['>']['L'] = '^';
	turns['>']['R'] = 'v';
	turns['v']['L'] = '>';
	turns['v']['R'] = '<';
	turns['<']['L'] = 'v';
	turns['<']['R'] = '^';

	map<char, Move> movement;
	movement['^'] = Move{ .dx =  0, .dy = -1 };
	movement['>'] = Move{ .dx =  1, .dy =  0 };
	movement['v'] = Move{ .dx =  0, .dy =  1 };
	movement['<'] = Move{ .dx = -1, .dy =  0 };

        int x = 0;
        int y = 0;
	char direction = '^';

        for (string& element : directions) {
		direction = turns[direction][element[0]];
		for (int i = 0; i < stoi(element.substr(1, string::npos)); i++) {
			x = x + movement[direction].dx;
			y = y + movement[direction].dy;
		}
        }

	cout << abs(x) + abs(y) << endl;
        return 0;
}
