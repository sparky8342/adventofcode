#include <iostream>
#include <fstream>
#include <vector>
#include <map>
using namespace std;

bool customsort(const pair<char,int> &a, const pair<char,int> &b) {
	return b.second < a.second;
}

int main() {
	vector<map<char, int>> lettercounts;
	string line;
	ifstream in("input.txt");
	getline(in, line);

	while (getline(in, line)) {
		for (int i = 0; i < line.size(); i++) {
			if (lettercounts.size() < i + 1) {
				map<char, int> charcount;
				charcount[line[i]] = 1;
				lettercounts.push_back(charcount);
			}
			else if (lettercounts[i].count(line[i]) == 1) {
				lettercounts[i][line[i]]++;
			}
			else {
				lettercounts[i][line[i]] = 1;
			}
		}
	}
	in.close();

	string most_common;
	string least_common;

	for (map<char, int>& m : lettercounts) {
		vector<pair<char, int>> vec;
                copy(m.begin(), m.end(), back_inserter<vector<pair<char, int>>>(vec));
                sort(vec.begin(), vec.end(), customsort);
		most_common += vec[0].first;
		least_common += vec[vec.size() - 1].first;
	}
	cout << most_common << endl << least_common << endl;

	return 0;
}
