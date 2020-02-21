#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <algorithm>
using namespace std;

vector<string> split(string& str, const string& delimiter) {
	string s = str;
	vector<string> tokens;
	size_t pos = 0;
	string token;
	while ((pos = s.find(delimiter)) != string::npos) {
		token = s.substr(0, pos);
		if (token.length() > 0) {
			tokens.push_back(token);
		}
		s.erase(0, pos + delimiter.length());
	}
	if (s.length() > 0) {
		tokens.push_back(s);
	}
	return tokens;
}

int main() {
        vector<string> data;
        string line;
        ifstream in("input.txt");
        while (getline(in, line)) {
                data.push_back(line);
        }
        in.close();

	// part1
	int num = 0;
	for (string& coords : data) {
		vector<string> lengths_str;
		lengths_str = split(coords, "  ");
		vector<int> lengths;
		for (int i = 0; i < 3; i++) {
			lengths.push_back(stoi(lengths_str[i]));
		}
		sort(lengths.begin(), lengths.end());
		if (lengths[0] + lengths[1] > lengths[2]) {
			num++;
		}
	}
	cout << num << endl;

	// part2
	num = 0;
	vector<vector<int>> lengths;
	for (string& coords : data) {
		vector<string> lengths_str;
		lengths_str = split(coords, "  ");
		vector<int> row;
		for (int i = 0; i < 3; i++) {
			int n = stoi(lengths_str[i]);
			row.push_back(stoi(lengths_str[i]));
		}
		lengths.push_back(row);

		if (lengths.size() == 3) {
			for (int i = 0; i < 3; i++) {
				vector<int> col;
				for (int j = 0; j < 3; j++) {
					col.push_back(lengths[j][i]);
				}
				sort(col.begin(), col.end());
				if (col[0] + col[1] > col[2]) {
					num++;
				}
			}
			lengths.clear();
		}
	}
	cout << num << endl;
	return 0;
}
