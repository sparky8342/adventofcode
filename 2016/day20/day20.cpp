#include <iostream>
#include <fstream>
#include <regex>
#include <vector>
using namespace std;

#define MAX 4294967295

int main() {
	const regex range_regex("^([0-9]+)-([0-9]+)$");

	vector<pair<long,long>> ranges;

	string line;
	ifstream in("input.txt");
	while (getline(in, line)) {
		smatch match;
		if (regex_match(line, match, range_regex)) {
			//cout << match[1] << " " << match[2] << endl;
			ranges.push_back(make_pair(stol(match[1]), stol(match[2])));
		}
	}

	sort(ranges.begin(), ranges.end()); 

	bool found_min = false;
	long min = 0;
	long ip = 0;
	long free_ips = 0;
	for (int i = 0; i < ranges.size(); i++) {
		pair<long,long> range = ranges[i];

		if (range.first <= ip) {
			if (range.second + 1 > ip) {
				ip = range.second + 1;
			}
		}
		else {
			if (found_min == false) {
				min = ip;
				found_min = true;
			}
			free_ips += range.first - ip;
			ip = range.second + 1;
		}

	}
	if (ip < MAX) {
		free_ips += MAX - ip + 1;
	}

	cout << min << endl << free_ips << endl;
	return 0;
}
