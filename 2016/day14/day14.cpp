#include <iostream>
#include <fstream>
#include <openssl/md5.h>
#include <sstream>
#include <iomanip>
#include <vector>
#include <map>
#include <algorithm>
using namespace std;

#define INPUT_SALT "jlmsuwbz"

struct Hash {
	string hash_str;
	int seen;
	bool three;
	bool five;
	char repeat3;
	char repeat5;
};

string get_hash(string str) {
	unsigned char result[MD5_DIGEST_LENGTH];
	MD5((unsigned char*)str.c_str(), str.size(), result);

	ostringstream sout;
	sout<<hex<<setfill('0');
	for(long long c: result)
	{
		    sout<<setw(2)<<(long long)c;
	}

	string out = sout.str();
	return out;
}

bool in_a_row(string str, char &ch, int amount) {
	int count = 1;
	for (int i = 1; i < str.size(); i++) {
		if (str[i] == str[i-1]) {
			count++;
			if (count == amount) {
				ch = str[i];
				return true;
			}
		}
		else {
			count = 1;
		}
	}

	return false;
}

bool three_in_a_row(string str, char &ch) {
	return in_a_row(str, ch, 3);
}

bool five_in_a_row(string str, char &ch) {
	return in_a_row(str, ch, 5);
}

Hash generate_hash_struct(int n) {
	string str = INPUT_SALT + to_string(n);
	string hash_str = get_hash(str);

	Hash hash = Hash{ .hash_str = hash_str, .seen = n };
	char ch;
	if (three_in_a_row(hash_str, ch)) {
		hash.three = true;
		hash.repeat3 = ch;
	}
	if (five_in_a_row(hash_str, ch)) {
		hash.five = true;
		hash.repeat5 = ch;
	}
	return hash;
}


int main() {
	vector <Hash> valid_hashes;
	vector <Hash> hash_window;

	// set up window of 1001 hashes
	for (int i = 0; i <= 1000; i++) {
		hash_window.push_back(generate_hash_struct(i));
	}

	int n = 1001;
	while (1) {
		// find first three in a row, keep 1001 in the window
		while (hash_window[0].three = false) {
			hash_window.erase(hash_window.begin());
			hash_window.push_back(generate_hash_struct(n));
			n++;
		}
			
		Hash candidate = hash_window[0];
		bool ok = false;
		for (int i = 1; i <= 1001; i++) {
			Hash hash = hash_window[i];
			if (hash.five == true && hash.repeat5 == candidate.repeat3) {
				ok = true;
				break;
			}
		}
		if (ok) {
			valid_hashes.push_back(candidate);
			cout << candidate.hash_str << " " << candidate.seen << endl;
			if (valid_hashes.size() == 64) {
				cout << valid_hashes[63].seen << endl;
				break;
			}
		}
	
		hash_window.erase(hash_window.begin());
		hash_window.push_back(generate_hash_struct(n));
		n++;
	}

	return 0;
}
