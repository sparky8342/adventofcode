#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <algorithm>
using namespace std;

vector<string> split(string& s, const string& delimiter) {
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
	return 0;
}
