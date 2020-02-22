#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <map>
#include <regex>
using namespace std;

bool customsort(const pair<char,int> &a, const pair<char,int> &b) { 
	if (a.second > b.second) {
		return true;
	}
	else if (a.second < b.second) {
		return false;
	}
	else if (a.first < b.first) {
		return true;
	}
	else {
		return false;
	}
} 

int main() {
	int total = 0;
	const regex name_regex("([a-z\\-]+)([0-9]+)\\[(.*)\\]");
	string line;
	int northpole;
	ifstream in("input.txt");
	while (getline(in, line)) {
		smatch base_match;
		regex_match(line, base_match, name_regex);

		string code = base_match[1]; 
		int sector_id = stoi(base_match[2]);
		string checksum = base_match[3];

		map<char, int> letters;

		for (int i = 0; i < code.size(); i++) {
			char ch = code[i];
			if (ch != '-') {
				if (letters.count(ch) > 0) {
					letters[ch]++;
				}
				else {
					letters[ch] = 1;
				}
			}
		}

		vector<pair<char, int>> vec;
		copy(letters.begin(), letters.end(), back_inserter<vector<pair<char, int>>>(vec));
		sort(vec.begin(), vec.end(), customsort);

		bool ok = true;
		for (int i = 0; i < checksum.size(); i++) {
			if (vec[i].first != checksum[i]) {
				ok = false;
				break;
			}
		}
		if (ok == true) {
			total += sector_id;
			int shift_id = sector_id % 26;
			string room_name;
			for (int i = 0; i < code.size(); i++) {
				char ch = code[i];
				if (ch != '-') {
					int n = int(ch);
					n += shift_id;
					if (n > 122) {
						n -= 26;
					}
					room_name += char(n);
				}
				else {
					room_name += '-';
				}
			}
			if (room_name == "northpole-object-storage-") {
				northpole = sector_id;
			}
		}	

        }
        in.close();
	cout << total << endl << northpole << endl;
	return 0;
}
