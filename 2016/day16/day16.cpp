#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
using namespace std;

#define INPUT "10001001100000001"

string find_checksum(vector<int> bits, int target_length) {
	// dragon curve
	int size = bits.size();
	while (size < target_length) {
		bits.push_back(0);
		for (int i = size - 1; i >= 0; i--) {
			bits.push_back(bits[i] ^ 1);
		}
		size = bits.size();
	}

	// work out checksum
	vector<int>::const_iterator start = bits.begin();
	vector<int>::const_iterator end = bits.begin() + target_length;
	vector<int> checksum(start, end);

	while (checksum.size() % 2 == 0) {
		vector<int> new_checksum;
		for (int i = 0; i < checksum.size(); i += 2) {
			if (checksum[i] == checksum[i + 1]) {
				new_checksum.push_back(1);
			}
			else {
				new_checksum.push_back(0);
			}
		}
		checksum = new_checksum;
	}

	string checksum_str;
	for (int i = 0; i < checksum.size(); i++) {
		checksum_str = checksum_str + to_string(checksum[i]);
	}
	return checksum_str;
}

int main() {
	string input = INPUT;
	vector<int> bits;
	for (int i = 0; i < input.size(); i++) {
		bits.push_back(input[i] - '0');
	}

	// part 1
	string checksum = find_checksum(bits, 272);
	cout << checksum << endl;

	// part 2
	checksum = find_checksum(bits, 35651584);
	cout << checksum << endl;

	return 0;
}
