#include <iostream>
#include <string>
#include <fstream>
#include <map>
using namespace std;

int main() {
	int valid = 0;
	int ssl = 0;
	string ip_address;
	ifstream in("input.txt");
	while (getline(in, ip_address)) {
		bool in_brackets = false;
		bool found = false;
		bool abba_result = false;
		map <string, bool> aba;
		map <string, bool> bab;
		for (int i = 0; i < ip_address.size() - 2; i++) {
			if (ip_address[i] == '[') {
				in_brackets = true;
			}
			else if (ip_address[i] == ']') {
				in_brackets = false;
			}
			else {
				if (
					abba_result == false
					&& ip_address[i] != ip_address[i + 1]
					&& ip_address[i] == ip_address[i + 3]
					&& ip_address[i + 1] == ip_address[i + 2]
				) {
					if (in_brackets == true) {
						found = false;
						abba_result = true;
					}
					else {
						found = true;
					}
				}
				if (ip_address[i] != ip_address[i + 1] && ip_address[i] == ip_address[i + 2]) {
					string str = ip_address.substr(i, 3);
					if (in_brackets == true) {
						bab[str] = true;
					}
					else {
						aba[str] = true;
					}
				}
			}
		}
		if (found == true) {
			valid++;
		}
		for(map<string, bool>::iterator iter = aba.begin(); iter != aba.end(); ++iter) {
			string k = iter->first;
			string match;
			match = match + k[1] + k[0] + k[1];
			if (bab.count(match) == 1) {
				ssl++;
				break;
			}
		}
	}
	in.close();
	cout << valid << endl << ssl << endl;
	return 0;
}
